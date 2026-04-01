package handler

import (
	"airline-booking/internal/service"
	"airline-booking/internal/transport/http/response"
	"github.com/gin-gonic/gin"
)

type FlightHandler struct {
	flightService *service.FlightService
}

func NewFlightHandler(flightService *service.FlightService) *FlightHandler {
	return &FlightHandler{flightService: flightService}
}

func (h *FlightHandler) ListFlights(c *gin.Context) {
	flights, err := h.flightService.ListFlights()
	if err != nil {
		response.Error(c, 500, "failed to fetch flights")
		return
	}
	response.Success(c, 200, flights)
}
