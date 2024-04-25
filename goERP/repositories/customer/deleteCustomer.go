package customer

import (
	db2 "goERP/config/db"
	"log"
)

type DeleteCustomer struct {
}

const (
	deleteSQL string = `DELETE FROM CLIENTES WHERE ID = $1`
)

func (e *DeleteCustomer) Delete(id int64) error {
	db, err := db2.PgConn()
	defer db2.HandleCloseDb(db)

	_, err = db.Exec(deleteSQL, id)
	if err != nil {
		log.Println("error deleting customer: ", err.Error())
		return err
	}

	return nil
}
