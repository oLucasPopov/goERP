package locationService

import (
	"goERP/types/location"
	"goERP/types/restResponses"
	"net/http"
)

func (cs *LocationService) Get(id int64) (*location.Cities, *restResponses.Error) {
	returnedCustomer, err := cs.getCityRepository.Get(id)
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
