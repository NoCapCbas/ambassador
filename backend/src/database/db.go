package database

import (
  "database/sql"
  "log"
  "fmt"

  _ "github.com/lib/pq"

	_ "ambassador/src/models"
)

const (
  host = "db"
  user = "postgres"
  password = "root"
  dbname = "ambassador"
  port = "5432"
  sslmode = "disable"
  timezone = "UTC"
)

var DB *sql.DB

func Connect() {
	var err error

	// Connect to PostgreSQL database
   dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
    log.Fatalf("Error connecting to database: %v\n", err)
	}

  err = DB.Ping() 
  if err != nil {
    log.Fatalf("Error pinging database: %v\n", err)
  }
}

func AutoMigrate() {
	// Auto migrate models
	_, err := DB.Exec(`

  CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password BYTEA NOT NULL,
    is_ambassador BOOLEAN NOT NULL
  );
  `) 
	if err != nil {
    log.Fatalf("Error creating users table: %v\n", err)
	}
}

func CloseDB() {
  err := DB.Close()
  if err != nil {
    log.Fatalf("Error closing database connection: %v\n", err)
  }
}
