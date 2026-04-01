package router

import (
	"airline-booking/internal/repository"
	"airline-booking/internal/service"
	"airline-booking/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func New(flightRepo repository.FlightRepository, bookingRepo repository.BookingRepository) *gin.Engine {
	flightService := service.NewFlightService(flightRepo)
	bookingService := service.NewBookingService(flightRepo, bookingRepo)

	healthHandler := handler.NewHealthHandler()
	flightHandler := handler.NewFlightHandler(flightService)
	bookingHandler := handler.NewBookingHandler(bookingService)

	r := gin.Default()
	r.GET("/health", healthHandler.GetHealth)
	r.GET("/flights", flightHandler.ListFlights)
	r.GET("/bookings", bookingHandler.ListBookings)
	r.POST("/bookings", bookingHandler.CreateBooking)

	return r
}
