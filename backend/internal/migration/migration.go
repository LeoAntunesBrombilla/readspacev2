package migration

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func ApplyMigrations(dbConnectionString string) error {

	m, err := migrate.New(
		"file:///backend/migrations", dbConnectionString,
	)

	if err != nil {
		if err == migrate.ErrNilVersion {
			log.Println("The migrations source URL is incorrectly formatted.")
		} else {
			log.Printf("Another error occurred: %s", err)
		}
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Migration error: %+v", err)
		return err
	}

	log.Printf("Success migrating")

	return nil
}
