package db

import (
	"context"
	"golabs/config"
	"log"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func ConnectRedis() {
	Rdb = createRedisConnection()
}

func createRedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetEnv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	// Check if the Redis connection is working
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("Connected to Redis")

	return rdb
}
