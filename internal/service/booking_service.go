package service

import (
	"airline-booking/internal/domain"
	"airline-booking/internal/repository"
)

type BookingService struct {
	flightRepo  repository.FlightRepository
	bookingRepo repository.BookingRepository
}

func NewBookingService(flightRepo repository.FlightRepository, bookingRepo repository.BookingRepository) *BookingService {
	return &BookingService{
		flightRepo:  flightRepo,
		bookingRepo: bookingRepo,
	}
}

func (s *BookingService) ListBookings() ([]domain.Booking, error) {
	return s.bookingRepo.List()
}

func (s *BookingService) CreateBooking(input domain.CreateBookingRequest) (domain.Booking, error) {
	if err := s.flightRepo.ReserveSeats(input.FlightID, input.SeatCount); err != nil {
		return domain.Booking{}, err
	}
	return s.bookingRepo.Create(input)
}
