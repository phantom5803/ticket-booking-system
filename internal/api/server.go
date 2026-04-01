package api

import (
	"errors"

	"airline-booking/internal/models"
	"airline-booking/internal/store"
	"github.com/gin-gonic/gin"
)

type Server struct {
	flightStore  *store.FlightStore
	bookingStore *store.BookingStore
}

func NewServer(flightStore *store.FlightStore, bookingStore *store.BookingStore) *Server {
	return &Server{
		flightStore:  flightStore,
		bookingStore: bookingStore,
	}
}

func (s *Server) Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", s.handleHealth)
	router.GET("/flights", s.handleListFlights)
	router.GET("/bookings", s.handleListBookings)
	router.POST("/bookings", s.handleCreateBooking)

	return router
}

func (s *Server) handleHealth(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func (s *Server) handleListFlights(c *gin.Context) {
	c.JSON(200, s.flightStore.List())
}

func (s *Server) handleListBookings(c *gin.Context) {
	c.JSON(200, s.bookingStore.List())
}

func (s *Server) handleCreateBooking(c *gin.Context) {
	var input models.CreateBookingRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid json body"})
		return
	}

	if input.FlightID == "" || input.PassengerName == "" || input.PassengerEmail == "" || input.SeatCount <= 0 {
		c.JSON(400, gin.H{"error": "flightId, passengerName, passengerEmail and positive seatCount are required"})
		return
	}

	if err := s.flightStore.ReserveSeats(input.FlightID, input.SeatCount); err != nil {
		if errors.Is(err, store.ErrFlightNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, store.ErrNotEnoughSeats) {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": "unexpected error"})
		return
	}

	booking := s.bookingStore.Create(input)
	c.JSON(201, booking)
}
