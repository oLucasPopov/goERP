package customer

import (
	db2 "goERP/config/db"
	"goERP/types/customer"
	"goERP/types/location"
	"log"
)

type AddCustomer struct {
}

const (
	insertSQL string = `
		INSERT INTO clientes (
			nome_fantasia
		   ,razao_social
		   ,cpf_cnpj
		   ,telefone
		   ,celular
		   ,email
		   ,observacoes
		   ,id_estado
		   ,id_cidade
		   ,cep
		   ,rua
		   ,numero
		   ,bairro
		   ,referencia
		   ,tp_codigo) 
		VALUES($1, $2, $3, $4, $5, $6, $7, NULLIF($8, 0), NULLIF($9, 0), $10, $11, $12, $13, $14, NULLIF($15, 0))
		returning 
		    id
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
		   ,coalesce(tp_codigo, -1);`
)

func (e *AddCustomer) Add(addCustomer customer.AddCustomer) (customer.Customer, error) {
	db, err := db2.PgConn()
	createdCustomer := customer.Customer{}
	createdCustomer.Location = location.FullLocation{}

	if err != nil {
		return createdCustomer, err
	}
	defer db2.HandleCloseDb(db)

	stm, err := db.Prepare(insertSQL)
	if err != nil {
		log.Println("error adding employee: ", err.Error())
		return createdCustomer, err
	}
	defer db2.HandleCloseStmt(stm)

	res, err := stm.Query(
		addCustomer.Name,
		addCustomer.CompanyName,
		addCustomer.CpfCnpj,
		addCustomer.Phone,
		addCustomer.Cellphone,
		addCustomer.Email,
		addCustomer.Obs,
		addCustomer.Location.StateId,
		addCustomer.Location.CityId,
		addCustomer.Location.Zipcode,
		addCustomer.Location.Address,
		addCustomer.Location.Number,
		addCustomer.Location.Neighbourhood,
		addCustomer.Location.Reference,
		addCustomer.PriceTableCode,
	)
	if err != nil {
		return createdCustomer, err
	}

	defer db2.HandleCloseRows(res)

	if res.Next() {
		err := res.Scan(
			&createdCustomer.Id,
			&createdCustomer.CompanyName,
			&createdCustomer.Name,
			&createdCustomer.CpfCnpj,
			&createdCustomer.Phone,
			&createdCustomer.Cellphone,
			&createdCustomer.Email,
			&createdCustomer.Obs,
			&createdCustomer.Location.StateId,
			&createdCustomer.Location.CityId,
			&createdCustomer.Location.Zipcode,
			&createdCustomer.Location.Address,
			&createdCustomer.Location.Number,
			&createdCustomer.Location.Neighbourhood,
			&createdCustomer.Location.Reference,
			&addCustomer.PriceTableCode,
		)
		if err != nil {
			return createdCustomer, err
		}
	}

	return createdCustomer, nil
}
