package migration

import (
	"database/sql"
	"time"
)

type Migration struct {
	transaction *sql.Tx
	Timestamp   time.Time
	Name        string
	Content     string
}

type Migrations = []Migration
