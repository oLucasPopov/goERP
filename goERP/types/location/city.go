package location

type City struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	StateId int64  `json:"state_id"`
}

type Cities []City
