package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"simplechat/chat"
	"simplechat/setup"
)

func Room(r *chat.Room, cfg *setup.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(cfg.UserKey())
		username := user.(string)

		r.ServeHTTP(c.Writer, c.Request, username, cfg)
	}
}
