package bot

import (
	"github.com/gin-gonic/gin"

	"simplechat/bot/middleware"
	"simplechat/bot/routes"
	"simplechat/setup"
)

func RunServer(cfg *setup.Config) {
	router := gin.Default()

	private := router.Group("/")
	private.Use(middleware.AuthRequired(cfg))
	routes.PrivateRoutes(private, cfg)

	router.Run(cfg.BotPort())
}
