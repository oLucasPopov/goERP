package customerService

import (
	"goERP/types/customer"
	"goERP/types/restResponses"
	"net/http"
)

func (cs *CustomerService) Get(id int64) (*customer.Customer, *restResponses.Error) {
	returnedCustomer, err := cs.getCustomerRepository.Get(id)
	if err != nil {
		return nil, &restResponses.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if returnedCustomer == nil {
		return nil, &restResponses.Error{
			Status: http.StatusNotFound,
		}
	}

	return returnedCustomer, nil
}
