package handlers

import (
	"github.com/gin-gonic/gin"

	"server/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	engine := gin.New()
	engine.LoadHTMLFiles("./assets/templates/restaurants.html")

	h.RegisterRoutes(engine)
	return engine
}
