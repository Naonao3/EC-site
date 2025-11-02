package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/Naonao3/EC-site/backend/config"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("Successfully connected to Redis")
	return client, nil
}

func CloseRedis(client *redis.Client) error {
	return client.Close()
}