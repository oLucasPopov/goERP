package employees

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goERP/controllers"
	employeesService "goERP/services/employees"
	employeeType "goERP/types/employee"
	"goERP/types/restResponses"
	"net/http"
	"strconv"
)

type UpdateEmployee struct {
	employeeService employeesService.EmployeeService
}

func MakeUpdateEmployee(employeeService employeesService.EmployeeService) *UpdateEmployee {
	return &UpdateEmployee{
		employeeService: employeeService,
	}
}

func (ae *UpdateEmployee) Handle(w http.ResponseWriter, r *http.Request) {
	var updateEmployee employeeType.AddEmployee

	vars := mux.Vars(r)

	stringId, ok := vars["id"]
	if !ok {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: "the parameter id is required"})
		return
	}

	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: "parameter id provided incorrectly"})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updateEmployee); err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: err.Error()})
		return
	}

	res, errResponse := ae.employeeService.Update(updateEmployee, id)
	if err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, *errResponse)
		return
	}

	controllers.SendResponse(w, http.StatusCreated, res)
}
