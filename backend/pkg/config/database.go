package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseConnStr string
}

func New() *Config {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	return &Config{
		DatabaseConnStr: connStr,
	}
}
