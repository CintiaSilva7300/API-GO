package main

import (
	"API-GO/config"
	// "API-GO/domain"
	// "API-GO/mock"
	"API-GO/service/message"
	// "context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pool, err := config.GetDBConnectionPool()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer pool.Close()

	fmt.Println("Successfully connected to the database")

	// Insert mock users into the database (if needed)
	// users := mock.MockUsers()
	// for _, user := range users {
	// 	err = domain.InsertUser(context.Background(), pool, user)
	// 	if err != nil {
	// 		log.Fatalf("Error inserting user: %v", err)
	// 	}
	// }

	fmt.Println("Mock users created successfully")

	// Start the message listener
	message.StartMessageListener(pool)
}
