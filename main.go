package main

import (
	"API-GO/config"
	"API-GO/domain"
	"API-GO/mock"
	"API-GO/service/message"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Get the connection string from the config
	// dbconnStr := config.GetDBConnectionString()

	// // Attempt to open a connection to the database
	// db, err := sql.Open("postgres", dbconnStr)
	// if err != nil {
	// 	log.Fatalf("Could not connect to database: %v", err)
	// }

	// // Check if the connection is actually alive
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("Could not establish a connection: %v", err)
	// }

	// defer db.Close()

	// fmt.Println("Successfully connected to the database")

	// // Create a new user
	// users := mock.MockUsers()

	// // Insert mock users into the database
	// for _, user := range users {
	// 	err = domain.InsertUser(db, user)
	// 	if err != nil {
	// 		log.Fatalf("Error inserting user: %v", err)
	// 	}
	// }

	// fmt.Println("Mock users created successfully")

	// Get the connection string from the config
	dbconnStr := config.GetDBConnectionString()

	// Attempt to open a connection to the database
	db, err := sql.Open("postgres", dbconnStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Check if the connection is actually alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not establish a connection: %v", err)
	}

	defer db.Close()

	fmt.Println("Successfully connected to the database")

	// Insert mock users into the database (if needed)
	users := mock.MockUsers()
	for _, user := range users {
		err = domain.InsertUser(db, user)
		if err != nil {
			log.Fatalf("Error inserting user: %v", err)
		}
	}

	fmt.Println("Mock users created successfully")

	// Start the message listener
	message.StartMessageListener(db)
}
