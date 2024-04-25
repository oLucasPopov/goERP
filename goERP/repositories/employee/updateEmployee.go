package employeeRepository

import (
	db2 "goERP/config/db"
	employeeType "goERP/types/employee"
	"goERP/types/location"
)

type UpdateEmployee struct {
}

const (
	updateSQL string = `
	update funcionarios
	   set nome         = coalesce($1, nome)
		    ,cpf_cnpj     = coalesce($2, cpf_cnpj)
		    ,telefone     = coalesce($3, telefone)
		    ,celular      = coalesce($4, celular)
		    ,email        = coalesce($5, email)
		    ,observacoes  = coalesce($6, observacoes)
		    ,id_estado    = coalesce($7, id_estado)
		    ,id_cidade    = coalesce($8, id_cidade)
		    ,cep          = coalesce($9, cep)
		    ,rua          = coalesce($10, rua)
		    ,numero       = coalesce($11, numero)
		    ,bairro       = coalesce($12, bairro)
		    ,referencia   = coalesce($13, referencia)
		    ,tipo_salario = coalesce($14, tipo_salario)
		    ,salario      = coalesce($15, salario)
		where id = $16
		returning *`
)

func (e *UpdateEmployee) Update(addEmployee employeeType.AddEmployee, id int64) (employeeType.Employee, error) {
	createdEmployee := employeeType.Employee{}
	createdEmployee.Location = location.FullLocation{}

	db, err := db2.PgConn()
	if err != nil {
		return createdEmployee, err
	}
	defer db2.HandleCloseDb(db)

	stm, err := db.Prepare(updateSQL)
	if err != nil {
		return createdEmployee, err
	}
	defer db2.HandleCloseStmt(stm)

	res, err := stm.Query(
		addEmployee.Name,
		addEmployee.CpfCnpj,
		addEmployee.Phone,
		addEmployee.Cellphone,
		addEmployee.Email,
		addEmployee.Obs,
		addEmployee.Location.StateId,
		addEmployee.Location.CityId,
		addEmployee.Location.Zipcode,
		addEmployee.Location.Address,
		addEmployee.Location.Number,
		addEmployee.Location.Neighbourhood,
		addEmployee.Location.Reference,
		addEmployee.SalaryType,
		addEmployee.Salary,
		id,
	)
	if err != nil {
		return createdEmployee, err
	}
	defer db2.HandleCloseRows(res)

	if res.Next() {
		err := res.Scan(
			&createdEmployee.Id,
			&createdEmployee.Name,
			&createdEmployee.CpfCnpj,
			&createdEmployee.Phone,
			&createdEmployee.Cellphone,
			&createdEmployee.Email,
			&createdEmployee.Obs,
			&createdEmployee.Location.StateId,
			&createdEmployee.Location.CityId,
			&createdEmployee.Location.Zipcode,
			&createdEmployee.Location.Address,
			&createdEmployee.Location.Number,
			&createdEmployee.Location.Neighbourhood,
			&createdEmployee.Location.Reference,
			&createdEmployee.SalaryType,
			&createdEmployee.Salary)
		if err != nil {
			return createdEmployee, err
		}
	}

	return createdEmployee, nil
}
