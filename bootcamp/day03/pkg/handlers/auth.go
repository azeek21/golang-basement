package handlers

import (
	"net/http"
	"server/types"
	"server/types/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogIn(c *gin.Context) {
	res, err := h.services.Auth.LogIn(types.LoginDTO{
		Email:    "login",
		Password: "pp",
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meeesage": res})
}

func (h *Handler) LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"meeesage": "OK"})
}

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	id, err := h.services.Auth.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"created": id})
}
