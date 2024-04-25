package customers

import (
	"github.com/gorilla/mux"
	"goERP/controllers"
	customerService "goERP/services/customers"
	"goERP/types/restResponses"
	"net/http"
	"strconv"
)

type DeleteCustomer struct {
	customerService customerService.CustomerService
}

func MakeDeleteCustomer(employeeService customerService.CustomerService) *DeleteCustomer {
	return &DeleteCustomer{
		customerService: employeeService,
	}
}

func (ae *DeleteCustomer) Handle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stringId, ok := params["id"]

	if !ok {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{
			Message: "the field id is required",
		})
		return
	}

	id, err := strconv.ParseInt(stringId, 10, 64)

	rErr := ae.customerService.Delete(id)
	if err != nil {
		controllers.ErrorResponse(w, rErr.Status, *rErr)
		return
	}

	controllers.SendResponse(w, http.StatusOK, r)
}
