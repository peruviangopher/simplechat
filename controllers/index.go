package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"simplechat/helpers"
	"simplechat/setup"
)

func IndexGetHandler(cfg *setup.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(cfg.UserKey())
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "Log in to start!",
			"user":    user,
		})
	}
}

func DashboardGetHandler(cfg *setup.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(cfg.UserKey())
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content":  "Chat rooms:",
			"user":     user,
			"rooms":    helpers.GetChatRoomsForView(cfg.Rooms()),
			"chatport": cfg.Port(),
		})
	}
}


