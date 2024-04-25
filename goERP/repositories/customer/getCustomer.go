package customer

import (
	db2 "goERP/config/db"
	"goERP/types/customer"
	"goERP/types/location"
	"log"
)

type GetCustomer struct {
}

const (
	getSQL string = `
		SELECT id
		   ,nome_fantasia
		   ,razao_social
		   ,cpf_cnpj
		   ,telefone
		   ,celular
		   ,email
		   ,observacoes
		   ,coalesce(id_estado, -1)
		   ,coalesce(id_cidade, -1)
		   ,cep
		   ,rua
		   ,numero
		   ,bairro
		   ,referencia
		   ,coalesce(tp_codigo, -1)
	FROM CLIENTES WHERE ID = $1`
)

func (e *GetCustomer) Get(id int64) (*customer.Customer, error) {
	db, err := db2.PgConn()
	returnedCustomer := customer.Customer{}
	returnedCustomer.Location = location.FullLocation{}

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

	if res.Next() {
		err := res.Scan(
			&returnedCustomer.Id,
			&returnedCustomer.CompanyName,
			&returnedCustomer.Name,
			&returnedCustomer.CpfCnpj,
			&returnedCustomer.Phone,
			&returnedCustomer.Cellphone,
			&returnedCustomer.Email,
			&returnedCustomer.Obs,
			&returnedCustomer.Location.StateId,
			&returnedCustomer.Location.CityId,
			&returnedCustomer.Location.Zipcode,
			&returnedCustomer.Location.Address,
			&returnedCustomer.Location.Number,
			&returnedCustomer.Location.Neighbourhood,
			&returnedCustomer.Location.Reference,
			&returnedCustomer.PriceTableCode,
		)
		if err != nil {
			return nil, err
		}
	}

	if returnedCustomer.Id == 0 {
		return nil, nil
	}

	return &returnedCustomer, nil
}
