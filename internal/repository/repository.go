package repository

import (
	"errors"

	"airline-booking/internal/domain"
)

var ErrFlightNotFound = errors.New("flight not found")
var ErrNotEnoughSeats = errors.New("not enough seats")

type FlightRepository interface {
	List() ([]domain.Flight, error)
	ReserveSeats(flightID string, seats int) error
}

type BookingRepository interface {
	List() ([]domain.Booking, error)
	Create(input domain.CreateBookingRequest) (domain.Booking, error)
}
