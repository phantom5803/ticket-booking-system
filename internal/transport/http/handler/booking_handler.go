package handler

import (
	"errors"

	"airline-booking/internal/domain"
	"airline-booking/internal/repository"
	"airline-booking/internal/service"
	"airline-booking/internal/transport/http/response"
	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	bookingService *service.BookingService
}

func NewBookingHandler(bookingService *service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

func (h *BookingHandler) ListBookings(c *gin.Context) {
	bookings, err := h.bookingService.ListBookings()
	if err != nil {
		response.Error(c, 500, "failed to fetch bookings")
		return
	}
	response.Success(c, 200, bookings)
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var input domain.CreateBookingRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, 400, "invalid json body")
		return
	}

	if input.FlightID == "" || input.PassengerName == "" || input.PassengerEmail == "" || input.SeatCount <= 0 {
		response.Error(c, 400, "flightId, passengerName, passengerEmail and positive seatCount are required")
		return
	}

	booking, err := h.bookingService.CreateBooking(input)
	if err != nil {
		if errors.Is(err, repository.ErrFlightNotFound) {
			response.Error(c, 404, err.Error())
			return
		}
		if errors.Is(err, repository.ErrNotEnoughSeats) {
			response.Error(c, 409, err.Error())
			return
		}
		response.Error(c, 500, "unexpected error")
		return
	}

	response.Success(c, 201, booking)
}
