package store

import (
	"errors"
	"fmt"
	"sync"

	"airline-booking/internal/models"
)

var ErrFlightNotFound = errors.New("flight not found")
var ErrNotEnoughSeats = errors.New("not enough seats")

type FlightStore struct {
	mu      sync.Mutex
	flights map[string]models.Flight
}

func NewFlightStore() *FlightStore {
	initialFlights := map[string]models.Flight{
		"F1": {
			ID:             "F1",
			FlightNumber:   "AI-101",
			From:           "DEL",
			To:             "BOM",
			DepartureTime:  "2026-04-02T09:00:00Z",
			ArrivalTime:    "2026-04-02T11:10:00Z",
			AvailableSeats: 60,
		},
		"F2": {
			ID:             "F2",
			FlightNumber:   "AI-202",
			From:           "BLR",
			To:             "DEL",
			DepartureTime:  "2026-04-02T14:15:00Z",
			ArrivalTime:    "2026-04-02T16:50:00Z",
			AvailableSeats: 40,
		},
	}

	return &FlightStore{flights: initialFlights}
}

func (s *FlightStore) List() []models.Flight {
	s.mu.Lock()
	defer s.mu.Unlock()

	flights := make([]models.Flight, 0, len(s.flights))
	for _, f := range s.flights {
		flights = append(flights, f)
	}
	return flights
}

func (s *FlightStore) ReserveSeats(flightID string, seats int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	flight, ok := s.flights[flightID]
	if !ok {
		return ErrFlightNotFound
	}
	if flight.AvailableSeats < seats {
		return ErrNotEnoughSeats
	}

	flight.AvailableSeats -= seats
	s.flights[flightID] = flight
	return nil
}

type BookingStore struct {
	mu       sync.Mutex
	bookings map[string]models.Booking
	counter  int
}

func NewBookingStore() *BookingStore {
	return &BookingStore{
		bookings: map[string]models.Booking{},
		counter:  0,
	}
}

func (s *BookingStore) List() []models.Booking {
	s.mu.Lock()
	defer s.mu.Unlock()

	bookings := make([]models.Booking, 0, len(s.bookings))
	for _, b := range s.bookings {
		bookings = append(bookings, b)
	}
	return bookings
}

func (s *BookingStore) Create(input models.CreateBookingRequest) models.Booking {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++
	id := fmt.Sprintf("B%d", s.counter)
	booking := models.Booking{
		ID:             id,
		FlightID:       input.FlightID,
		PassengerName:  input.PassengerName,
		PassengerEmail: input.PassengerEmail,
		SeatCount:      input.SeatCount,
		Status:         "CONFIRMED",
	}
	s.bookings[id] = booking
	return booking
}
