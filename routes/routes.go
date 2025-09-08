package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simplechat/chat"
	"simplechat/globals"
	"strconv"

	controllers "simplechat/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup, roomsQuantity *string) {

	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())

	loop, err := strconv.Atoi(*roomsQuantity)
	if err != nil {
		loop = globals.DefaultRooms
	}

	for i := 1; i <= loop; i++ {
		r := chat.NewRoom()
		g.GET(fmt.Sprintf("/room/%v", i), controllers.Room(r))
		go r.Run()
	}
}
