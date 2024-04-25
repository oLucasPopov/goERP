package employeeRepository

import (
	"database/sql"
	db2 "goERP/config/db"
	employeeType "goERP/types/employee"
	"goERP/types/location"
	"log"
)

type AddEmployee struct {
}

const (
	insertSQL string = `
	insert into funcionarios(
		nome,
		cpf_cnpj,
		telefone,
		celular,
		email,
		observacoes,
		id_estado,
		id_cidade,
		cep,
		rua,
		numero,
		bairro,
		referencia,
		tipo_salario,
		salario)
	values(
		 COALESCE($1, '')
		,COALESCE($2, '')
		,COALESCE($3, '')
		,COALESCE($4, '')
		,COALESCE($5, '')
		,COALESCE($6, '')
		,$7
		,$8
		,COALESCE($9,  '')		
		,COALESCE($10, '')
		,COALESCE($11, '')
		,COALESCE($12, '')
		,COALESCE($13, '')
		,$14
		,$15) returning *`
)

func (e *AddEmployee) Add(addEmployee employeeType.AddEmployee) (employeeType.Employee, error) {
	db, err := db2.PgConn()
	createdEmployee := employeeType.Employee{}
	createdEmployee.Location = location.FullLocation{}

	if err != nil {
		return createdEmployee, err
	}
	defer db2.HandleCloseDb(db)

	stm, err := db.Prepare(insertSQL)

	if err != nil {
		log.Println("error adding employee: ", err.Error())
		return createdEmployee, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("error closing db connection at addEmployee:", err.Error())
		}
	}(db)

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
