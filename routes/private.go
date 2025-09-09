package routes

import (
	"fmt"
	"simplechat/chat/consumer"

	"github.com/gin-gonic/gin"

	"simplechat/chat"
	"simplechat/controllers"
	"simplechat/setup"
)

func PrivateRoutes(g *gin.RouterGroup, cfg *setup.Config) {
	g.GET("/logout", controllers.LogoutGetHandler(cfg))

	for i := 1; i <= cfg.Rooms(); i++ {
		r := chat.NewRoom(fmt.Sprintf("%v", i))
		go consumer.Consume(r)
		g.GET(fmt.Sprintf("/room/%v", i), controllers.Room(r, cfg))
		go r.Run()
	}

	g.GET("/dashboard", controllers.DashboardGetHandler(cfg))
}
