package config

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func ConnectRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        panic("Failed to connect to Redis")
    }
}