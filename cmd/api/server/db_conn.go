package server

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func InitDBMigrate() {
	databaseURL := "postgresql://postgres:@localhost:5432/cats_social?sslmode=disable"

	migrationsPath := "db/migrations"

	m, err := migrate.New(
		"file://"+migrationsPath,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Could not create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migrations: %v", err)
	}

	log.Println("Migration successful")

}
