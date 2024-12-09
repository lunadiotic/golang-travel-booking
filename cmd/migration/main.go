// cmd/migration/main.go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lunadiotic/golang-travel-booking/pkg/config"
	"github.com/lunadiotic/golang-travel-booking/pkg/database"
)

func main() {
	// Flag untuk menentukan action up atau down
	var action string
	flag.StringVar(&action, "action", "up", "migration action (up or down)")
	flag.Parse()

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

	// Run Migration
	if err := database.RunMigration(db, action); err != nil {
		log.Fatal("Cannot run migration:", err)
	}

	fmt.Printf("Migration %s completed successfully\n", action)
}
