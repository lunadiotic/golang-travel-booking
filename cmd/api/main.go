// cmd/api/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print some env vars to verify
	fmt.Printf("Database Host: %s\n", os.Getenv("DB_HOST"))
	fmt.Printf("App Environment: %s\n", os.Getenv("APP_ENV"))
}
