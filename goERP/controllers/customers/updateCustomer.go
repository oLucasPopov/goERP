package customers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goERP/controllers"
	customerService "goERP/services/customers"
	"goERP/types/customer"
	"goERP/types/restResponses"
	"net/http"
	"strconv"
)

type UpdateCustomer struct {
	customerService customerService.CustomerService
}

func MakeUpdateCustomer(employeeService customerService.CustomerService) *UpdateCustomer {
	return &UpdateCustomer{
		customerService: employeeService,
	}
}

func (ae *UpdateCustomer) Handle(w http.ResponseWriter, r *http.Request) {
	var addCustomer customer.AddCustomer
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

	if err := json.NewDecoder(r.Body).Decode(&addCustomer); err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: err.Error()})
		return
	}

	res, rErr := ae.customerService.Update(addCustomer, id)
	if rErr != nil {
		controllers.ErrorResponse(w, rErr.Status, *rErr)
		return
	}

	controllers.SendResponse(w, http.StatusOK, res)
}
