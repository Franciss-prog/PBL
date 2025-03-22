package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx" // Optional, but makes handling SQL easier
	_ "github.com/lib/pq"     // PostgreSQL driver
)

var DB *sqlx.DB

func ConnectDB() {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL is not set in environment variables")
	}
	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	fmt.Println("Connected to PostgreSQL!")
}
