package db

import (
	"database/sql"
	"log"
)

func HandleCloseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Println("error closing db statement:", err.Error())
	}
}

func HandleCloseStmt(rows *sql.Stmt) {
	err := rows.Close()
	if err != nil {
		log.Println("error closing db statement:", err.Error())
	}
}

func HandleCloseDb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println("error closing db:", err.Error())
	}
}
