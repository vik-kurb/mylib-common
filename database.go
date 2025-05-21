package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbName     = "DB_NAME"
)

func SetupDB(envPath string) (*sql.DB, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	dbURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv(dbUser), os.Getenv(dbPassword), os.Getenv(dbHost), os.Getenv(dbPort), os.Getenv(dbName))
	return sql.Open("postgres", dbURL)
}

func SetupDBByURL(envPath string, dbURL string) (*sql.DB, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", os.Getenv(dbURL))
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("Failed to close db: %v", err)
	}
}

func CloseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Printf("Failed to close rows: %v", err)
	}
}
