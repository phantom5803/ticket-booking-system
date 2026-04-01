package service

import (
	"airline-booking/internal/domain"
	"airline-booking/internal/repository"
)

type FlightService struct {
	flightRepo repository.FlightRepository
}

func NewFlightService(flightRepo repository.FlightRepository) *FlightService {
	return &FlightService{flightRepo: flightRepo}
}

func (s *FlightService) ListFlights() ([]domain.Flight, error) {
	return s.flightRepo.List()
}
