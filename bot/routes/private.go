package routes

import (
	"github.com/gin-gonic/gin"

	"simplechat/bot/controllers"
	"simplechat/setup"
)

func PrivateRoutes(g *gin.RouterGroup, cfg *setup.Config) {
	g.GET("/stock", controllers.StockHandler(cfg))
}
