// package main

// import (
// 	"github.com/backend/config"
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// 	// app.Get(route, handler function)

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})
// 	config.ConnectDB()
// 	config.ConnectRedis()

// 	app.Listen(":3000")

// }

package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
    ctx := context.Background()
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        Password: "",           // No password set
        DB: 0,                  // Use default DB
    })

    pong, err := client.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(pong) // Should print "PONG"
}