package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"simplechat/chat"
	"simplechat/config"
	"simplechat/controllers"
)

func PublicRoutes(g *gin.RouterGroup, cfg *config.Config) {

	g.GET("/login", controllers.LoginGetHandler(cfg))
	g.POST("/login", controllers.LoginPostHandler(cfg))
	g.GET("/", controllers.IndexGetHandler(cfg))

}

func PrivateRoutes(g *gin.RouterGroup, cfg *config.Config) {
	g.GET("/logout", controllers.LogoutGetHandler(cfg))

	for i := 1; i <= cfg.Rooms(); i++ {
		r := chat.NewRoom()
		g.GET(fmt.Sprintf("/room/%v", i), controllers.Room(r, cfg))
		go r.Run()
	}

	g.GET("/dashboard", controllers.DashboardGetHandler(cfg))
}
