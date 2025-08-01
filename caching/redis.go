package caching

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"orderservice/config"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

var redisClient *RedisClient

func InitializeRedisClient() *RedisClient {
	cfg := config.Load()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	redisClient = &RedisClient{client: client}
	return redisClient
}

func GetRedisClient() *RedisClient {
	if redisClient == nil {
		log.Fatalf("Redis client is not initialized. Call InitializeRedis first.")
	}
	return redisClient
}

func (r *RedisClient) Set(ctx context.Context, key string, data string) error {
	err := r.client.Set(ctx, key, data, 0).Err()
	if err != nil {
		return fmt.Errorf("could not set payee data: %v", err)
	}
	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("payee data not found for key: %s", key)
	}
	if err != nil {
		return "", fmt.Errorf("could not get payee data: %v", err)
	}
	return data, nil
}

func (r *RedisClient) HSetData(ctx context.Context, key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("could not marshal data to JSON: %v", err)
	}

	err = r.client.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("could not set data in Redis: %v", err)
	}

	return nil
}

func (r *RedisClient) GetData(ctx context.Context, key string, result interface{}) error {
	jsonData, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("data not found for key: %s", key)
	}
	if err != nil {
		return fmt.Errorf("could not get data from Redis: %v", err)
	}

	err = json.Unmarshal([]byte(jsonData), result)
	if err != nil {
		return fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	return nil
}
