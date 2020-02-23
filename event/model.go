package event

// Event ...
type Event struct {
	ID          uint32 `json:"event_id"`
	Date        uint32 `json:"date"`
	Name        string `json:"event_name"`
	MaxPeople   uint32 `json:"max_people"`
	Description string `json:"description"`
}

type NG struct {
	Date uint32 `json:"date"`
}
