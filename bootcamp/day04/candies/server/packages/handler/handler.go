package handler

import (
	"candies/server/packages/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	engine := gin.New()
	h.RegisterRoutes(engine)
	return engine
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
