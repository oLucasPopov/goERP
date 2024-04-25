package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"goERP/config/env"
)

func PgConn() (*sql.DB, error) {
	pgConnection := env.Configs.Pg

	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		pgConnection.User,
		pgConnection.Password,
		pgConnection.Database,
		pgConnection.Host,
		pgConnection.Port,
	))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		closeErr := db.Close()

		if closeErr != nil {
			return nil, closeErr
		}

		return nil, err
	}

	return db, nil
}
