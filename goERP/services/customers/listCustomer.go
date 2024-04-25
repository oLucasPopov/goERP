package customerService

import (
	"goERP/types/customer"
	"goERP/types/restResponses"
	"net/http"
)

func (cs *CustomerService) List(page int64) (*customer.Customers, *restResponses.Error) {
	returnedCustomers, err := cs.listCustomerRepository.List(page)
	if err != nil {
		return nil, &restResponses.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if returnedCustomers == nil {
		return nil, &restResponses.Error{
			Status: http.StatusNotFound,
		}
	}

	return returnedCustomers, nil
}
