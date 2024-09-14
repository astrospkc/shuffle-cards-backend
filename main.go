package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/backend/config"

	"github.com/backend/router"
	"github.com/gorilla/handlers"
)

// var ctx = context.Background()

// func storeUser(user model.User1) {
//     // Store user as a hash in Redis
//     userdetails := []string{
//         "name" ,user.Name,
//         "college" ,user.College,
//     }

//     // err := config.Rdb.HSet(ctx, "user:"+user.ID, map[string]interface{}{
//     //     "name":    user.Name,
//     //     "college": user.College,
//     // }).Err()

//     err := config.Rdb.HSet(ctx, "user:"+user.ID, userdetails).Err()

//     if err != nil {
//         log.Fatalf("Error saving user to Redis: %v", err)
//     }
//     fmt.Printf("Stored user: %v\n", user)
// }

// func storeMarks(marks model.Marks) {
//     // Store marks as a hash in Redis
//     err := config.Rdb.HSet(ctx, "marks:"+marks.UserID+":"+marks.Subject, map[string]interface{}{
//         "score": marks.Score,
//     }).Err()
//     if err != nil {
//         log.Fatalf("Error saving marks to Redis: %v", err)
//     }
//     fmt.Printf("Stored marks for user %s in subject %s: %d\n", marks.UserID, marks.Subject, marks.Score)
// }



func main() {
    fmt.Println("Mongodb API")

    // Initialize the Redis client
    config.InitializeRedis()

//   user := model.User1{ID: "1", Name: "John Doe", College: "Example University"}
//     marks := model.Marks{UserID: "1", Subject: "Math", Score: 95}

//     // Store user and marks in Redis
//     storeUser(user)
//     storeMarks(marks)


    // Initialize the router
    r := router.Router()

    // Set up CORS middleware
    r.Use(handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
    ))

    
    fmt.Println("Redis connected successfully")

    // Start the HTTP server
    fmt.Println("Server is getting started")
    err := http.ListenAndServe(":3000", r)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }

    // This line will never be reached because ListenAndServe is blocking
    fmt.Println("Server is up and running")
}