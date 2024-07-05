package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func StartAt() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("start_at", time.Now().UnixMilli())
		c.Request = c.Request.WithContext(c)
		c.Next()
	}
}
