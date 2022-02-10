// @file: main.go
// @date: 2022/02/10

package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("connect to redis failed: ", err)
	}
	keys := rdb.Keys(ctx, "*").Val()
	for _, key := range keys {
		log.Println(key)
		kt := rdb.Type(ctx, key).Val()
		switch kt {
		case "string":
			log.Println(rdb.Get(ctx, key).Val())
		case "list":
			log.Println(rdb.LRange(ctx, key, 0, -1).Val())
		case "set":
			log.Println(rdb.SMembers(ctx, key).Val())
		case "zset":
			log.Println(rdb.ZRange(ctx, key, 0, -1).Val())
		case "hash":
			log.Println(rdb.HGetAll(ctx, key).Val())
		default:
			log.Println("unknown type")
		}
	}
}
