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

	fmt.Println("tj-beads application")
	fmt.Println("Database connected successfully")
	os.Exit(0)
}