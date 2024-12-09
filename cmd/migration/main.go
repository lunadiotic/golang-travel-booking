// cmd/migration/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lunadiotic/golang-travel-booking/pkg/config"
	"github.com/lunadiotic/golang-travel-booking/pkg/database"
)

func main() {
	// Flag untuk menentukan action up atau down
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go migrate <up|down>")
	}

	action := os.Args[2] // Ambil "up" atau "down" dari argumen

	fmt.Printf("Running migration %s\n", action)

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Setup Database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	if err := database.TestConnection(db); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Run Migration
	if err := database.RunMigration(db, action); err != nil {
		log.Fatal("Cannot run migration:", err)
	}

	fmt.Printf("Migration %s completed successfully\n", action)
}
