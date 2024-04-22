package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"server/core/paging"
	"server/types"
	"server/types/models"
)

func (h *Handler) GetAllRestaurants(c *gin.Context) {
	anyPage, _ := c.Get(paging.PAGIONATION_KEY)
	pagination := anyPage.(types.PagingIncoming)
	restaurants, err := h.services.Restaurant.GetAll(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

func (h *Handler) GetRestaurantById(c *gin.Context) {
	id := c.Param("id")
	res, err := h.services.Restaurant.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) CraeteRestaunrat(c *gin.Context) {
	var restaurant models.Restaurant
	if err := c.BindJSON(&restaurant); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.Restaurant.Create(restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"created": id})
}

func (h *Handler) DeleteRestaurant(c *gin.Context) {
	id := c.Param("id")
	deleteSuccess, err := h.services.Restaurant.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": deleteSuccess})
}

func (h *Handler) SearchRestaurans(c *gin.Context) {
	_page, _ := c.Get(paging.PAGIONATION_KEY)
	pagination := _page.(types.PagingIncoming)
	query := c.DefaultQuery("q", "")

	restaurants, err := h.services.Search(query, pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, restaurants)
}

func (h *Handler) UpdateRestaurant(c *gin.Context) {
	id := c.Param("id")
	restaurant := models.Restaurant{}
	if err := c.Bind(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	restaurant.ID = id

	if _, err := h.services.Restaurant.Update(restaurant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": restaurant.ID})
}

func (h *Handler) GetClosest(c *gin.Context) {
	id := c.Param("id")
	_page, _ := c.Get(paging.PAGIONATION_KEY)
	pagination := _page.(types.PagingIncoming)

	_distance := c.DefaultQuery("distance", "500")
	distance, err := strconv.Atoi(_distance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "distance must be integer"})
		c.Abort()
		return
	}

	res, err := h.services.Restaurant.GetClosest(id, pagination, distance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, res)
}
