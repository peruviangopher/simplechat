package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"simplechat/chat"
	"simplechat/controllers"
	"simplechat/globals"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup, roomsQuantity *string) {
	g.GET("/logout", controllers.LogoutGetHandler())

	rooms, err := strconv.Atoi(*roomsQuantity)
	if err != nil {
		rooms = globals.DefaultRooms
	}

	for i := 1; i <= rooms; i++ {
		r := chat.NewRoom()
		g.GET(fmt.Sprintf("/room/%v", i), controllers.Room(r))
		go r.Run()
	}

	g.GET("/dashboard", controllers.DashboardGetHandler(rooms))
}
