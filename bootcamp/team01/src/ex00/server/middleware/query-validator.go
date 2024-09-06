package middleware

import (
	"net/http"
	"replication/models"

	"github.com/gin-gonic/gin"
)

// must accept value as string and return error describing the problem if value is not valid
type ValidatorFunc = func(value string) error

func ParamValidatorMiddleware(key string, validator ValidatorFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param(key)
		err := validator(param)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewBadRequestResponse(err))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
