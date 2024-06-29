package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


func goEnvVar(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func OpenDB() (*sql.DB, error) {
	var db *sql.DB
	fmt.Println(goEnvVar("POSTGRES_PORT"))
	port, err := strconv.Atoi(goEnvVar("POSTGRES_PORT"))
	if err != nil {
		return db, err
	}

	return sql.Open("postgres", fmt.Sprintf(
			"host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable",
			port,
			goEnvVar("POSTGRES_USER"),
			goEnvVar("POSTGRES_PASSWORD"),
			goEnvVar("POSTGRES_DB"),
		),
	)
}