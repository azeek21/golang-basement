package middleware

import (
	"time"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/gin-gonic/gin"

	"github.com/JGLTechnologies/gin-rate-limit"
)

func getRequestIdentifier(c *gin.Context) string {
	return c.ClientIP()
}

func handleLimitExeededError(c *gin.Context, _ ratelimit.Info) {
	utils.RenderTempl(c, 200, components.AlertsContainer(models.ALERT_LEVELS.WARNING, "Relaax, Too many requests in a small timeframe."))
	c.Abort()
}

func RateLimitMiddleware() func(*gin.Context) {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 3,
	})

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: handleLimitExeededError,
		KeyFunc:      getRequestIdentifier,
	})
	return mw
}
