package middleware

import (
	"fmt"

	"github.com/azeek21/blog/models"
	"github.com/gin-gonic/gin"
)

// TODO:
func NewPagingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paging := models.NewDefaultPagingIncoming()
		_ = ctx.ShouldBind(paging)

		fmt.Printf("GOT PAGING INCOMING: %+v\n", paging)
		ctx.Set(models.PAGING_MODEL_NAME, paging)
		ctx.Next()
	}
}
