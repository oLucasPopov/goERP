package customer

import (
	db2 "goERP/config/db"
	"goERP/types/customer"
	"goERP/types/location"
	"log"
)

type ListCustomer struct {
}

const (
	listSQL string = `
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
	FROM CLIENTES 
	ORDER BY id
	OFFSET ($1 - 1) * 20 LIMIT 20`
)

func (e *ListCustomer) List(page int64) (*customer.Customers, error) {
	db, err := db2.PgConn()

	if err != nil {
		return nil, err
	}
	defer db2.HandleCloseDb(db)

	stm, err := db.Prepare(listSQL)
	if err != nil {
		log.Println("error listing employee: ", err.Error())
		return nil, err
	}
	defer db2.HandleCloseStmt(stm)

	res, err := stm.Query(page)
	if err != nil {
		return nil, err
	}
	defer db2.HandleCloseRows(res)

	var customers customer.Customers

	for res.Next() {
		returnedCustomer := customer.Customer{}
		returnedCustomer.Location = location.FullLocation{}
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

		customers = append(customers, returnedCustomer)
	}

	return &customers, nil
}
