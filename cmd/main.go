package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tj-beads/internal/db"
)

func main() {
	ctx := context.Background()

	database, err := db.New(ctx, "tj-beads.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Create users table
	if err := database.CreateUserTable(ctx); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	// Seed test user if not exists
	_, err = database.GetUserByUsername(ctx, "test")
	if err != nil {
		if _, err := database.CreateUser(ctx, "test", "test"); err != nil {
			log.Fatalf("Failed to create test user: %v", err)
		}
		fmt.Println("Created test user")
	}

	fmt.Println("tj-beads application")
	fmt.Println("Database connected successfully")
	os.Exit(0)
}