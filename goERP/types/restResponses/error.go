package restResponses

type Error struct {
	Status  int    `json:"-"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}
