package cache

import (
	"context"
	"fmt"
	"go-weather/internal/utilities/app"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICacheStorage interface {
	Get(ctx context.Context, key string) (*CacheResult, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}

type CacheResult struct {
	Key   string
	Value interface{}
}

type RedisStorage struct {
	Redis *redis.Client
}

func New(appConfig *app.AppConfig) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     appConfig.Config.GetString("cache.host"),
		Password: appConfig.Config.GetString("cache.passowrd"),
		DB:       0, // use default DB
	})

	c := &RedisStorage{
		Redis: rdb,
	}

	return c
}

func (c *RedisStorage) Get(ctx context.Context, key string) (*CacheResult, error) {
	value, err := c.Redis.Get(ctx, key).Result()
	if err != nil {
		return &CacheResult{}, fmt.Errorf("error when get key %v and error %v", key, err)
	}

	return &CacheResult{
		Key:   key,
		Value: value,
	}, nil
}

func (c *RedisStorage) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err := c.Redis.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("error when get key %v and error %v", key, err)
	}

	return nil
}
