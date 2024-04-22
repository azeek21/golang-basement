package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"server/core/paging"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination, err := paging.GetPaginationFromQueryParams(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		log.Printf("middleware [pagination]: parsed: %v\n", pagination)
		c.Set(paging.PAGIONATION_KEY, pagination)
		c.Next()
	}
}
