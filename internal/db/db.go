package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InitDB initializes the database connection.
func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err = db.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}

// CloseDB closes the database connection.
func CloseDB(db *sql.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("failed to close the database: %v", err)
		} else {
			fmt.Println("database connection closed")
		}
	}
}
