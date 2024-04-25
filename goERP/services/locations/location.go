package locationService

import "goERP/repositories/city"

type LocationService struct {
	getCityRepository city.GetCities
}

func MakeGetCityService(getCityRepository city.GetCities) *LocationService {
	return &LocationService{
		getCityRepository: getCityRepository,
	}
}
