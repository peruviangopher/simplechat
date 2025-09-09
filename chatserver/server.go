package chatserver

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/filesystem"
	"github.com/gin-gonic/gin"

	"simplechat/middleware"
	"simplechat/routes"
	"simplechat/setup"
)

func Run(cfg *setup.Config) {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	var sessionPath = "/tmp/"
	store := filesystem.NewStore(sessionPath, cfg.Secret())
	router.Use(sessions.Sessions("mysession", store))

	public := router.Group("/")
	routes.PublicRoutes(public, cfg)

	private := router.Group("/")
	private.Use(middleware.AuthRequired(cfg))
	routes.PrivateRoutes(private, cfg)

	router.Run(cfg.Port())
}
