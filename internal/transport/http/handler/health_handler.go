package handler

import (
	"airline-booking/internal/transport/http/response"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetHealth(c *gin.Context) {
	response.Success(c, 200, gin.H{"status": "ok"})
}
