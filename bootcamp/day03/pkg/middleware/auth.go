package middleware

import (
	"errors"
	"net/http"
	"server/pkg/service"
	"server/types"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const AUTH_SLUG = "Authorization"

func AuthMiddleware(userService service.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		_token := c.Request.Header.Get(AUTH_SLUG)
		token := strings.Split(_token, " ")[1]
		jwtParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unauthorized")
			}

			return []byte(types.GLOABAL_CONFIG.Secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil || !jwtParsed.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, ok := jwtParsed.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		user, err := userService.GetById(claims["sub"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
