package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(c *gin.Engine) {
	c.GET("/ping", h.Ping)
	c.GET("/buy-candy", h.BuyCandy)
}
