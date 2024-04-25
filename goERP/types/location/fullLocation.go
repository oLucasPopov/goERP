package location

type FullLocation struct {
	CityId     int64  `json:"city_id"`
	CityIbgeId int32  `json:"city_ibge_id,omitempty"`
	CityName   string `json:"city_name,omitempty"`

	StateId     int64  `json:"state_id"`
	StateIbgeId int32  `json:"state_ibge_id,omitempty"`
	StateName   string `json:"state_name,omitempty"`

	Zipcode       string `json:"zipcode"`
	Address       string `json:"address"`
	Number        string `json:"number"`
	Neighbourhood string `json:"neighbourhood"`
	Reference     string `json:"reference"`
}
