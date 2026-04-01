package memory

import (
	"sync"

	"airline-booking/internal/domain"
	"airline-booking/internal/repository"
)

type FlightRepository struct {
	mu      sync.Mutex
	flights map[string]domain.Flight
}

func NewFlightRepository() *FlightRepository {
	initialFlights := map[string]domain.Flight{
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

	return &FlightRepository{flights: initialFlights}
}

func (r *FlightRepository) List() ([]domain.Flight, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	flights := make([]domain.Flight, 0, len(r.flights))
	for _, f := range r.flights {
		flights = append(flights, f)
	}
	return flights, nil
}

func (r *FlightRepository) ReserveSeats(flightID string, seats int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	flight, ok := r.flights[flightID]
	if !ok {
		return repository.ErrFlightNotFound
	}
	if flight.AvailableSeats < seats {
		return repository.ErrNotEnoughSeats
	}

	flight.AvailableSeats -= seats
	r.flights[flightID] = flight
	return nil
}
