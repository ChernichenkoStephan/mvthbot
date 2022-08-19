package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func setupDB(lg *zap.SugaredLogger) (*sqlx.DB, error) {
	lg.Infoln("DB setup")

	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	// TODO change to Viper
	db, err := sqlx.Connect("postgres", "user=admin dbname=mvthdb sslmode=disable")
	if err != nil {
		return nil, err
	}

	lg.Infoln("DB setup success")
	return db, nil
}
