package employees

import (
	"encoding/json"
	"goERP/controllers"
	employeesService "goERP/services/employees"
	employeeType "goERP/types/employee"
	"goERP/types/restResponses"
	"net/http"
)

type AddEmployee struct {
	employeeService employeesService.EmployeeService
}

func MakeAddEmployee(employeeService employeesService.EmployeeService) *AddEmployee {
	return &AddEmployee{
		employeeService: employeeService,
	}
}

func (ae *AddEmployee) Handle(w http.ResponseWriter, r *http.Request) {
	var addEmployee employeeType.AddEmployee

	if err := json.NewDecoder(r.Body).Decode(&addEmployee); err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: err.Error()})
		return
	}

	res, err := ae.employeeService.Add(addEmployee)
	if err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, *err)
		return
	}

	controllers.SendResponse(w, http.StatusCreated, res)
}
