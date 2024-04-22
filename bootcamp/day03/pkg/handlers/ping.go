package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/types/models"
)

func (h *Handler) Ping(c *gin.Context) {
	_, _ = h.services.Restaurant.Create(models.Restaurant{})
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
