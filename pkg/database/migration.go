package database

import (
	"fmt"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB, action string) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting underlying sql.DB: %v", err)
	}

	fmt.Println("Starting migration process...")

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating postgres driver: %v", err)
	}

	// Get absolute path
	absPath, err := filepath.Abs("database/migrations")
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}
	fmt.Printf("Migration path: file://%s\n", absPath)

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", absPath),
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %v", err)
	}

	switch action {
	case "up":
		fmt.Println("Running UP migration...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("error running migration up: %v", err)
		}
		fmt.Println("Migration UP executed")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("error running migration down: %v", err)
		}
	default:
		return fmt.Errorf("invalid action: %s", action)
	}

	return nil
}
