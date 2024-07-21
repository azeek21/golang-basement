package middleware

import (
	"net/http"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/service"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

func isAborted(c *gin.Context, err error, strict bool) bool {
	if err != nil && strict {
		c.String(http.StatusUnauthorized, "unauthorized")
		c.Abort()
		return true
	}
	return false
}

// strict: true, every unauthed request will get 401
// strict: false, requests still pass, but when they are authenticated, a user will be attached to context
// TODO: refactor dependency injection
func AuthMiddleware(userService service.UserService, jwtService service.JwtService, secret string, strict bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// NOTE: this is used if middleware is called inside an already authed request. Prevents verifying and querying multiple times
		if utils.IsAuthed(c) {
			c.Next()
			return
		}

		tokenCookie, err := c.Request.Cookie(models.AUTH_COOKIE_NAME)
		if isAborted(c, err, strict) {
			return
		}

		token := tokenCookie.Value

		user_id, err := jwtService.VerifyJwt(token)
		if isAborted(c, err, strict) {
			return
		}

		if err == nil {
			new_id := uint(user_id)
			user, err := userService.GetUserById(new_id)
			if isAborted(c, err, strict) {
				return
			}
			c.Set(models.USER_MODEL_NAME, user)
		}

		c.Next()
	}
}

// TODO: finish this up. maybe integrate into auth middleware
func RoleMiddleware(allowRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _user, ok := c.Get(models.USER_MODEL_NAME); ok {
			user := _user.(models.User)
			if user.RoleCode != allowRole {
				c.String(http.StatusUnauthorized, "You don't have enough privilages to complete this action")
				c.Abort()
				return
			}
			c.Next()
			return
		}
		c.String(http.StatusUnauthorized, "You don't have enough privilages to complete this action")
		c.Abort()
	}
}
