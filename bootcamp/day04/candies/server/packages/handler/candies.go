package handler

import (
	"candies/server/types/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BuyCandy(c *gin.Context) {
	body := dtos.BuyCandyRequestDTO{}
	if err := c.ShouldBindJSON(&body); err != nil {
		resp := dtos.BuyCandyBadRequestResponseDTO{Error: "bad request: " + err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		c.Abort()
		return
	}
	resp, err := h.services.Candy.BuyCandy(&body)
	if err != nil {
		c.JSON(err.Error().HttpStatus, gin.H{"error": err.Error().Message})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, resp)
}
