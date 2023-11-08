package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// connection string details
	dsn := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"

	// trying to connect to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// pinging the database
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Successfully connected to the database!")

	return db, nil
}
