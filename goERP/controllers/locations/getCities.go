package locations

import (
	"github.com/gorilla/mux"
	"goERP/controllers"
	locationService "goERP/services/locations"
	"goERP/types/restResponses"
	"net/http"
	"strconv"
)

type GetCities struct {
	locationService locationService.LocationService
}

func MakeGetCities(locationService locationService.LocationService) *GetCities {
	return &GetCities{
		locationService: locationService,
	}
}

func (ae *GetCities) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	stringId, ok := vars["ibge-id"]
	if !ok {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: "the parameter ibge-id is required"})
		return
	}

	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		controllers.ErrorResponse(w, http.StatusBadRequest, restResponses.Error{Message: "parameter id provided incorrectly"})
		return
	}

	res, rErr := ae.locationService.Get(id)
	if rErr != nil {
		controllers.ErrorResponse(w, rErr.Status, *rErr)
		return
	}

	controllers.SendResponse(w, http.StatusOK, res)
}
