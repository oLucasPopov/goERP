package db

import (
	"database/sql"
	"log"
)

func HandleRollback(err error, tx *sql.Tx) error {
	log.Println("rolling back due to error", err.Error())
	if txError := tx.Rollback(); txError != nil {
		log.Println("error while rolling back due to error", txError.Error())
		return txError
	}

	return err
}
