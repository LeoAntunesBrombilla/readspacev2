package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	maxRetries = 5
	retryDelay = time.Second * 5
)

func Connect() (*pgxpool.Pool, error) {
	connStr := "host=db user=postgres password=postgres dbname=readspacev2 port=5432 sslmode=disable"

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
