package handlers

import (
	"github.com/gin-gonic/gin"

	"server/pkg/middleware"
)

func (h *Handler) RegisterRoutes(c *gin.Engine) {
	v1Group := c.Group("/v1")
	c.GET("/ping", h.Ping)
	{
		restaurantsGroup := v1Group.Group("/restaurants")
		{
			// restaurants rest
			restaurantsGroup.GET("/", middleware.PaginationMiddleware(), h.GetAllRestaurants)
			restaurantsGroup.POST("/", middleware.AuthMiddleware(h.services.User), h.CraeteRestaunrat)
			restaurantsGroup.GET("/:id", h.GetRestaurantById)
			restaurantsGroup.PUT("/:id", middleware.AuthMiddleware(h.services.User), h.UpdateRestaurant)
			restaurantsGroup.GET("/search/", middleware.PaginationMiddleware(), h.SearchRestaurans)
			restaurantsGroup.DELETE("/:id", middleware.AuthMiddleware(h.services.User), h.DeleteRestaurant)
			// single restaurant item rest
			restaurant := restaurantsGroup.Group("/:id/")
			{
				restaurant.GET("/closest", middleware.PaginationMiddleware(), h.GetClosest)
			}
		}

		authGroup := v1Group.Group("/auth")
		{
			authGroup.POST("/register", h.Register)
			authGroup.POST("/login", h.LogIn)
			authGroup.POST("/logout", middleware.AuthMiddleware(h.services.User), h.LogOut)
		}
	}
}
