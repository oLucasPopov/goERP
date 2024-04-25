package customer

import (
	db2 "goERP/config/db"
	"goERP/types/customer"
	"goERP/types/location"
	"log"
)

type UpdateCustomer struct {
}

const (
	updateSQL string = `
		UPDATE clientes 
		   SET nome_fantasia = $1
		   ,razao_social = $2
		   ,cpf_cnpj = $3
		   ,telefone = $4
		   ,celular = $5
		   ,email = $6
		   ,observacoes = $7
		   ,id_estado = NULLIF($8, 0)
		   ,id_cidade = NULLIF($9, 0)
		   ,cep = $10
		   ,rua = $11
		   ,numero = $12
		   ,bairro = $13
		   ,referencia = $14
		   ,tp_codigo = NULLIF($15, 0)
		where id = $16
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

func (e *UpdateCustomer) Update(addCustomer customer.AddCustomer, id int64) (customer.Customer, error) {
	db, err := db2.PgConn()
	updatedCustomer := customer.Customer{}
	updatedCustomer.Location = location.FullLocation{}

	if err != nil {
		return updatedCustomer, err
	}
	defer db2.HandleCloseDb(db)

	stm, err := db.Prepare(updateSQL)
	if err != nil {
		log.Println("error adding employee: ", err.Error())
		return updatedCustomer, err
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
		id,
	)
	if err != nil {
		return updatedCustomer, err
	}

	defer db2.HandleCloseRows(res)

	if res.Next() {
		err := res.Scan(
			&updatedCustomer.Id,
			&updatedCustomer.CompanyName,
			&updatedCustomer.Name,
			&updatedCustomer.CpfCnpj,
			&updatedCustomer.Phone,
			&updatedCustomer.Cellphone,
			&updatedCustomer.Email,
			&updatedCustomer.Obs,
			&updatedCustomer.Location.StateId,
			&updatedCustomer.Location.CityId,
			&updatedCustomer.Location.Zipcode,
			&updatedCustomer.Location.Address,
			&updatedCustomer.Location.Number,
			&updatedCustomer.Location.Neighbourhood,
			&updatedCustomer.Location.Reference,
			&addCustomer.PriceTableCode,
		)
		if err != nil {
			return updatedCustomer, err
		}
	}

	return updatedCustomer, nil
}
