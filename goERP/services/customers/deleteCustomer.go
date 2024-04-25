package customerService

import (
	"goERP/types/restResponses"
	"net/http"
)

func (cs *CustomerService) Delete(id int64) *restResponses.Error {
	err := cs.deleteCustomerRepository.Delete(id)
	if err != nil {
		return &restResponses.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}
