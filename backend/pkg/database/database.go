package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	maxRetries = 5
	retryDelay = time.Second * 5
)

func Connect() (*pgxpool.Pool, error) {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	var dbpool *pgxpool.Pool
	var err error
	for i := 1; i <= maxRetries; i++ {
		dbpool, err = pgxpool.Connect(context.Background(), connStr)
		if err == nil {
			fmt.Println("Successfully connected!")
			return dbpool, nil
		}

		fmt.Printf("Unable to connect to database: %v. Retry attempt %d\n", err, i)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("unable to connect to database: %v", err)
}
