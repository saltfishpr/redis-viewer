// @description:
// @file: rv.go
// @date: 2022/03/07

package rv

import (
	"context"
	"sync/atomic"

	"github.com/SaltFishPr/redis-viewer/internal/constant"

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

// GetKeys .
func GetKeys(rdb redis.UniversalClient, cursor uint64, match string, count int64) <-chan string {
	keys := make(chan string, 1)

	go func() {
		ctx := context.TODO()
		switch rdb := rdb.(type) {
		case *redis.ClusterClient:
			err := rdb.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
				iter := client.Scan(ctx, cursor, match, count).Iterator()
				for iter.Next(ctx) {
					keys <- iter.Val()
				}
				if err := iter.Err(); err != nil {
					return err
				}
				return nil
			})
			if err != nil {

			}
		default:
			iter := rdb.Scan(ctx, cursor, match, count).Iterator()
			for iter.Next(ctx) {
				keys <- iter.Val()
			}
			if err := iter.Err(); err != nil {

			}
		}
		close(keys)
	}()

	return keys
}
