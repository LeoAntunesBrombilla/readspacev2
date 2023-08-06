package database

import (
	"context"
	"fmt"
	"readspacev2/pkg/config"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	maxRetries = 5
	retryDelay = time.Second * 5
)

func Connect() (*pgxpool.Pool, error) {

	var dbpool *pgxpool.Pool
	var err error
	cfg := config.New()

	for i := 1; i <= maxRetries; i++ {
		dbpool, err = pgxpool.Connect(context.Background(), cfg.DatabaseConnStr)
		if err == nil {
			fmt.Println("Successfully connected!")
			return dbpool, nil
		}

		fmt.Printf("Unable to connect to database: %v. Retry attempt %d\n", err, i)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("unable to connect to database: %v", err)
}
