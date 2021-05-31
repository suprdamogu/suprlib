package suprlib

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func RedisNew(ip string, port string, password string) *redis.Client {
	if ip == "" {
		ip = "localhost"
	}
	if port == "" {
		port = "6379"
	}
	addr := ip + ":" + port

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}

func RedisSet(client *redis.Client, key string, value string) error {
	ctx := context.Background()
	return client.Set(ctx, key, value, 0).Err()
}

func RedisGet(client *redis.Client, key string) (string, error) {
	ctx := context.Background()
	return client.Get(ctx, key).Result()
}
