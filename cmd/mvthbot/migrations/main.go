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
