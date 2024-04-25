package city

import (
	db2 "goERP/config/db"
	"goERP/types/location"
	"log"
)

type GetCities struct {
}

const (
	getSQL string = `
  SELECT C.id
        ,C.id_estado
        ,C.cidade
	FROM CIDADES C
	INNER JOIN ESTADOS E ON E.ID = C.id_estado
   WHERE E.id = $1`
)

func (e *GetCities) Get(id int64) (*location.Cities, error) {
	db, err := db2.PgConn()
	var cidades location.Cities

	if err != nil {
		return nil, err
	}
	defer db2.HandleCloseDb(db)

	stm, err := db.Prepare(getSQL)
	if err != nil {
		log.Println("error getting employee: ", err.Error())
		return nil, err
	}
	defer db2.HandleCloseStmt(stm)

	res, err := stm.Query(id)
	if err != nil {
		return nil, err
	}
	defer db2.HandleCloseRows(res)

	for res.Next() {
		var cidade location.City

		err := res.Scan(
			&cidade.Id,
			&cidade.StateId,
			&cidade.Name,
		)
		if err != nil {
			return nil, err
		}

		cidades = append(cidades, cidade)
	}

	return &cidades, nil
}
