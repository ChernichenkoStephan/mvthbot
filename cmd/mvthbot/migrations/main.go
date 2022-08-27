package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	if os.Getenv("DRIVER") == `` {
		log.Fatal(`Need DRIVER in environment variables`)
	}

	if os.Getenv("DATA_SOURCE_NAME") == `` {
		log.Fatal(`Need DATA_SOURCE_NAME in environment variables`)
	}

	if os.Getenv("MIGRATIONS_DIR") == `` {
		log.Fatal(`Need MIGRATIONS_DIR in environment variables`)
	}

	driver_str := os.Getenv("DRIVER")
	dataSourceName := os.Getenv("DATA_SOURCE_NAME")
	migPath := `file://` + os.Getenv("MIGRATIONS_DIR")

	db, err := sql.Open(driver_str, dataSourceName) //"")
	if err != nil {
		log.Fatalf("Error during open: %s", err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error during making the driver x: %s", err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		migPath,
		driver_str, driver)
	if err != nil {
		log.Fatalf("Error during db instance setup: %s", err.Error())
	}
	err = m.Up() // set all of migrations to run
	if err != nil {
		log.Fatalf("Error during migrations: %s", err.Error())
	} else {
		log.Println(`Migrate success`)
	}
	// m.Step(2) // or m.Step(2) if you want to explicitly set the number of migrations to run
}
