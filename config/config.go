// config/config.go
package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var ctx = context.Background()

// InitializeRedis initializes the Redis client
func InitializeRedis() {
    // Initialize the Redis client using your configuration values
Rdb = redis.NewClient(&redis.Options{
        Addr:     "redis-17587.c281.us-east-1-2.ec2.redns.redis-cloud.com:17587", 
        // Redis server address
        Password: "PjYvtI6BeUckZMenY5PfJexm0Oydn5Du",                       // Redis password
        DB:       0,                                                        // Use default DB
    })

    // Ping the Redis server to check connection
    _, err := Rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatal("Could not connect to Redis: %v", err)
    }

    // _, err = rdb.Set(ctx, "foo1", "barr1", 0).Result()
    // if err != nil{
    //     log.Fatal(err)
    // }
    // err = rdb.HSet(ctx, game.UserID,map[string]interface{}{
    //     "id"
    // }).Result()
	// if err != nil{
	// 	log.Fatal(err)
	// }

    // fmt.Println("result: ", result)

    // _, err = rdb.Get(ctx, "foo").Result()
    // if err != nil{
    //     log.Fatal(err)
    // }


    log.Println("Redis connected successfully")
}