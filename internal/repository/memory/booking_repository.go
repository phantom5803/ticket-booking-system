package memory

import (
	"fmt"
	"sync"

	"airline-booking/internal/domain"
)

type BookingRepository struct {
	mu       sync.Mutex
	bookings map[string]domain.Booking
	counter  int
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		bookings: map[string]domain.Booking{},
		counter:  0,
	}
}

func (r *BookingRepository) List() ([]domain.Booking, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	bookings := make([]domain.Booking, 0, len(r.bookings))
	for _, b := range r.bookings {
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *BookingRepository) Create(input domain.CreateBookingRequest) (domain.Booking, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	id := fmt.Sprintf("B%d", r.counter)
	booking := domain.Booking{
		ID:             id,
		FlightID:       input.FlightID,
		PassengerName:  input.PassengerName,
		PassengerEmail: input.PassengerEmail,
		SeatCount:      input.SeatCount,
		Status:         "CONFIRMED",
	}
	r.bookings[id] = booking
	return booking, nil
}
