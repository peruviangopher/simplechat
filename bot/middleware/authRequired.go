package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simplechat/setup"
)

func AuthRequired(cfg *setup.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if cfg.BotAPIKey() != c.GetHeader(cfg.BotAPIKeyName()) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or missing token LML",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
