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

type User struct {
	Event_id string `json:"event_id"`
	User     string `json:"user"`
}

type NG struct {
	Date uint32 `json:"date"`
}

type Coordinate struct {
	User      string  `json:"user"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
