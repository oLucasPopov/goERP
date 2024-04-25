package customers

import (
	"github.com/gorilla/mux"
	"goERP/controllers"
	customerService "goERP/services/customers"
	"goERP/types/restResponses"
	"net/http"
	"strconv"
)

type ListCustomer struct {
	customerService customerService.CustomerService
}

func MakeListCustomer(employeeService customerService.CustomerService) *ListCustomer {
	return &ListCustomer{
		customerService: employeeService,
	}
}

func (ae *ListCustomer) Handle(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		page int64 = 1
	)

	vars := mux.Vars(r)

	pageNumber, ok := vars["page"]
	if ok {
		page, err = strconv.ParseInt(pageNumber, 10, 64)
	}

	if err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: "parameter page provided incorrectly"})
		return
	}

	res, rErr := ae.customerService.List(page)
	if rErr != nil {
		controllers.ErrorResponse(w, rErr.Status, *rErr)
		return
	}

	controllers.SendResponse(w, http.StatusOK, res)
}
