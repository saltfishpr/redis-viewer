// Package rv redis operation.
package rv

import (
	"context"
	"sync/atomic"

	"github.com/saltfishpr/redis-viewer/internal/constant"

	"github.com/go-redis/redis/v8"
)

// CountKeys .
func CountKeys(rdb redis.UniversalClient, match string) (int, error) {
	ctx := context.TODO()

	switch rdb := rdb.(type) {
	case *redis.ClusterClient:
		var count int64

		err := rdb.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
			iter := client.Scan(ctx, 0, match, 0).Iterator()
			for iter.Next(ctx) {
				atomic.AddInt64(&count, 1)
				if count > constant.MaxScanCount {
					break
				}
			}
			if err := iter.Err(); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return 0, err
		}

		return int(count), nil
	default:
		var count int

		iter := rdb.Scan(ctx, 0, match, 0).Iterator()
		for iter.Next(ctx) {
			count++
			if count > constant.MaxScanCount {
				break
			}
		}
		if err := iter.Err(); err != nil {
			return 0, err
		}

		return count, nil
	}
}

//nolint:govet
type keyMessage struct {
	Key string
	Err error
}

// GetKeys .
func GetKeys(
	rdb redis.UniversalClient,
	cursor uint64,
	match string,
	count int64,
) <-chan keyMessage {
	res := make(chan keyMessage, 1)

	go func() {
		ctx := context.TODO()

		switch rdb := rdb.(type) {
		case *redis.ClusterClient:
			var i int64

			err := rdb.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
				iter := client.Scan(ctx, cursor, match, 0).Iterator()
				for iter.Next(ctx) {
					atomic.AddInt64(&i, 1)
					if i > count {
						break
					}
					res <- keyMessage{iter.Val(), nil}
				}
				if err := iter.Err(); err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				res <- keyMessage{"", err}
			}
		default:
			keys, _, err := rdb.Scan(ctx, cursor, match, count).Result()
			if err != nil {
				res <- keyMessage{"", err}
			} else {
				for _, key := range keys {
					res <- keyMessage{key, nil}
				}
			}
		}

		close(res)
	}()

	return res
}
