package domain

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
