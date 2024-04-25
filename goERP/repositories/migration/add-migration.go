package migration

import (
	db2 "goERP/config/db"
	"goERP/types/migration"
)

const (
	insertSQL string = `
	INSERT INTO MIGRATIONS(
		 MIGRATION_TIMESTAMP
		,MIGRATION_NAME     
		,MIGRATION_SQL      
	)
	values(
		 $1
		,$2
		,$3
	) ON CONFLICT(MIGRATION_NAME, MIGRATION_TIMESTAMP) DO NOTHING`
)

func (m *Migration) Add(migrations migration.Migrations) error {
	db, err := db2.PgConn()

	if err != nil {
		return err
	}
	defer db2.HandleCloseDb(db)

	tx, err := db2.NewTransaction(db)

	if err != nil {
		return err
	}

	for _, migrationData := range migrations {
		_, err := tx.Exec(
			insertSQL,
			migrationData.Timestamp,
			migrationData.Name,
			migrationData.Content,
		)

		if err != nil {
			return db2.HandleRollback(err, tx)
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
