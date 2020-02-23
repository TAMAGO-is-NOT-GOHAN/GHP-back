package event

// Event ...
type Event struct {
	ID          string       `json:"event_id"`
	Date        uint32       `json:"date"`
	Name        string       `json:"event_name"`
	MaxPeople   uint32       `json:"max_people"`
	Description string       `json:"description"`
	Coors       []Coordinate `json:"coordinate"`
}

type NG struct {
	Date uint32 `json:"date"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
