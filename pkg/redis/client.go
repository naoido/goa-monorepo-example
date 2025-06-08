package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

func NewClient(ctx context.Context, addr, password string, db int) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
	client := redis.NewClient(opts)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redisへの接続に失敗しました: %w", err)
	}
	return client, nil
}
