package models

type Flight struct {
	ID             string `json:"id"`
	FlightNumber   string `json:"flightNumber"`
	From           string `json:"from"`
	To             string `json:"to"`
	DepartureTime  string `json:"departureTime"`
	ArrivalTime    string `json:"arrivalTime"`
	AvailableSeats int    `json:"availableSeats"`
}

type Booking struct {
	ID             string `json:"id"`
	FlightID       string `json:"flightId"`
	PassengerName  string `json:"passengerName"`
	PassengerEmail string `json:"passengerEmail"`
	SeatCount      int    `json:"seatCount"`
	Status         string `json:"status"`
}

type CreateBookingRequest struct {
	FlightID       string `json:"flightId"`
	PassengerName  string `json:"passengerName"`
	PassengerEmail string `json:"passengerEmail"`
	SeatCount      int    `json:"seatCount"`
}
