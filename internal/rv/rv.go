// Package rv redis operation.
package rv

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/go-redis/redis/v8"
)

// errDone is a sentinel used to signal early exit from ForEachMaster without
// treating it as a real error.
var errDone = errors.New("done")

type KeyMessage struct {
	Key string
	Err error
}

// ScanKeys scans matched keys and returns one page of results (offset/limit).
// It is compatible with both standalone Redis and Redis Cluster.
func ScanKeys(rdb redis.UniversalClient, match string, offset, limit int64) <-chan KeyMessage {
	res := make(chan KeyMessage)

	go func() {
		defer close(res)

		ctx := context.Background()
		if limit <= 0 {
			return
		}

		pattern := normalizeMatch(match)
		var skipped int64
		var sent int64

		scanNode := func(ctx context.Context, client redis.UniversalClient) error {
			var cursor uint64
			for {
				keys, next, err := client.Scan(ctx, cursor, pattern, 100).Result()
				if err != nil {
					return err
				}

				for _, key := range keys {
					if skipped < offset {
						skipped++
						continue
					}
					if sent >= limit {
						return errDone
					}

					res <- KeyMessage{Key: key}
					sent++
				}

				cursor = next
				if cursor == 0 {
					return nil
				}
			}
		}

		var err error
		if cc, ok := rdb.(*redis.ClusterClient); ok {
			err = cc.ForEachMaster(ctx, func(ctx context.Context, c *redis.Client) error {
				return scanNode(ctx, c)
			})
		} else {
			err = scanNode(ctx, rdb)
		}

		if errors.Is(err, errDone) {
			err = nil
		}
		if err != nil {
			res <- KeyMessage{Err: err}
		}
	}()

	return res
}

// CountKeys returns the total number of keys matching match.
// It is compatible with both standalone Redis and Redis Cluster.
func CountKeys(rdb redis.UniversalClient, match string) (int64, error) {
	ctx := context.Background()
	pattern := normalizeMatch(match)
	var total atomic.Int64

	countNode := func(ctx context.Context, client redis.UniversalClient) error {
		var cursor uint64
		for {
			keys, next, err := client.Scan(ctx, cursor, pattern, 100).Result()
			if err != nil {
				return err
			}
			total.Add(int64(len(keys)))
			cursor = next
			if cursor == 0 {
				break
			}
		}
		return nil
	}

	if cc, ok := rdb.(*redis.ClusterClient); ok {
		if err := cc.ForEachMaster(ctx, func(ctx context.Context, c *redis.Client) error {
			return countNode(ctx, c)
		}); err != nil {
			return 0, err
		}
	} else {
		if err := countNode(ctx, rdb); err != nil {
			return 0, err
		}
	}

	return total.Load(), nil
}

func normalizeMatch(match string) string {
	if match == "" {
		return "*"
	}
	return match
}
