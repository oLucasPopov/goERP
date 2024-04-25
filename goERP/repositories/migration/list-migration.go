package migration

import (
	db2 "goERP/config/db"
	"goERP/types/migration"
)

func (m *Migration) List() (migration.Migrations, error) {
	db, err := db2.PgConn()
	if err != nil {
		return nil, err
	}

	defer db2.HandleCloseDb(db)

	rows, err := db.Query(`
	  SELECT MIGRATION_TIMESTAMP
		    ,MIGRATION_NAME
			,MIGRATION_SQL
		FROM MIGRATIONS 
	   WHERE NOT MIGRATION_EXECUTED
    	 ORDER BY MIGRATION_TIMESTAMP 
	`)

	defer db2.HandleCloseRows(rows)

	var migrations []migration.Migration

	for rows.Next() {
		migrationData := migration.Migration{}

		err := rows.Scan(
			&migrationData.Timestamp,
			&migrationData.Name,
			&migrationData.Content,
		)

		if err != nil {
			return nil, err
		}

		migrations = append(migrations, migrationData)
	}

	if err != nil {
		return nil, err
	}

	return migrations, nil
}
