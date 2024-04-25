package migration

import (
	"context"
	"database/sql"
	db2 "goERP/config/db"
	"goERP/types/migration"
	"log"
)

func (m *Migration) Persist(migrations migration.Migrations) error {
	db, err := db2.PgConn()

	if err != nil {
		return err
	}
	defer db2.HandleCloseDb(db)

	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, migrationData := range migrations {
		log.Println("executing migration", migrationData.Name)
		_, err := tx.Exec(migrationData.Content)

		if err != nil {
			return db2.HandleRollback(err, tx)
		}

		_, err = tx.Exec(`
			UPDATE MIGRATIONS 
			   SET MIGRATION_EXECUTED = TRUE 
			 WHERE MIGRATION_NAME = $1 
			   AND MIGRATION_TIMESTAMP = $2
		`,
			migrationData.Name,
			migrationData.Timestamp)

		if err != nil {
			return db2.HandleRollback(err, tx)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
