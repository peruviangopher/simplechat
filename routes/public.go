package routes

import (
	"github.com/gin-gonic/gin"

	"simplechat/controllers"
	"simplechat/setup"
)

func PublicRoutes(g *gin.RouterGroup, cfg *setup.Config) {

	g.GET("/login", controllers.LoginGetHandler(cfg))
	g.POST("/login", controllers.LoginPostHandler(cfg))
	g.GET("/", controllers.IndexGetHandler(cfg))

}
