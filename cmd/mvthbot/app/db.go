package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func setupDB(c *configuration, lg *zap.SugaredLogger) (*sqlx.DB, error) {
	lg.Infoln("DB setup")

	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics

	db, err := sqlx.Connect(c.Database.Driver, c.Database.SourceStr)
	if err != nil {
		return nil, err
	}

	lg.Infoln("DB setup success")
	return db, nil
}
