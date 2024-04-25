package customers

import (
	"github.com/gorilla/mux"
	"goERP/controllers"
	customerService "goERP/services/customers"
	"goERP/types/restResponses"
	"net/http"
	"strconv"
)

type GetCustomer struct {
	customerService customerService.CustomerService
}

func MakeGetCustomer(employeeService customerService.CustomerService) *GetCustomer {
	return &GetCustomer{
		customerService: employeeService,
	}
}

func (ae *GetCustomer) Handle(w http.ResponseWriter, r *http.Request) {
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

	res, rErr := ae.customerService.Get(id)
	if rErr != nil {
		controllers.ErrorResponse(w, rErr.Status, *rErr)
		return
	}

	controllers.SendResponse(w, http.StatusOK, res)
}
