package db

import (
	"context"
	"database/sql"
)

func NewTransaction(db *sql.DB) (*sql.Tx, error) {
	return db.BeginTx(context.Background(), &sql.TxOptions{})
}
