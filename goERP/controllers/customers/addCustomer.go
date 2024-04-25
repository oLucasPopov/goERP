package customers

import (
	"encoding/json"
	"goERP/controllers"
	customerService "goERP/services/customers"
	"goERP/types/customer"
	"goERP/types/restResponses"
	"net/http"
)

type AddCustomer struct {
	customerService customerService.CustomerService
}

func MakeAddCustomer(employeeService customerService.CustomerService) *AddCustomer {
	return &AddCustomer{
		customerService: employeeService,
	}
}

func (ae *AddCustomer) Handle(w http.ResponseWriter, r *http.Request) {
	var addCustomer customer.AddCustomer

	if err := json.NewDecoder(r.Body).Decode(&addCustomer); err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: err.Error()})
		return
	}

	res, err := ae.customerService.Add(addCustomer)
	if err != nil {
		controllers.ErrorResponse(w, err.Status, *err)
		return
	}

	controllers.SendResponse(w, http.StatusCreated, res)
}
