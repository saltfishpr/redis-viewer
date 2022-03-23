// @file: main.go
// @date: 2022/02/10

package main

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/saltfishpr/redis-viewer/internal/config"

	"github.com/go-redis/redis/v8"
)

func main() {
	config.LoadConfig()
	cfg := config.GetConfig()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addrs,
		Password: cfg.Password,
	})
	fmt.Println(countKeys(rdb))
	keys := getKeys(rdb)
	for key := range keys {
		fmt.Println(key)
	}
}

func countKeys(rdb *redis.ClusterClient) int {
	ctx := context.TODO()
	var count int64
	rdb.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
		iter := client.Scan(ctx, 0, "", 0).Iterator()
		for iter.Next(ctx) {
			atomic.AddInt64(&count, 1)
		}
		if err := iter.Err(); err != nil {
			return err
		}
		return nil
	})
	return int(count)
}

func getKeys(rdb *redis.ClusterClient) <-chan string {
	ctx := context.TODO()
	keys := make(chan string, 1)
	go func() {
		rdb.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
			iter := client.Scan(ctx, 0, "", 0).Iterator()
			for iter.Next(ctx) {
				keys <- iter.Val()
			}
			if err := iter.Err(); err != nil {
				return err
			}
			return nil
		})
		close(keys)
	}()
	return keys
}
