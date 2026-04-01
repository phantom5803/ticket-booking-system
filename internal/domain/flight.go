package domain

type Flight struct {
	ID             string `json:"id"`
	FlightNumber   string `json:"flightNumber"`
	From           string `json:"from"`
	To             string `json:"to"`
	DepartureTime  string `json:"departureTime"`
	ArrivalTime    string `json:"arrivalTime"`
	AvailableSeats int    `json:"availableSeats"`
}
