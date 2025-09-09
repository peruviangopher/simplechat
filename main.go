package main

import (
	"flag"
	"strconv"
	//"html/template"
	//"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/filesystem"
	"github.com/gin-gonic/gin"

	"simplechat/globals"
	"simplechat/middleware"
	"simplechat/routes"
)

func main() {
	var roomsQuantity = flag.String("rooms", strconv.Itoa(globals.DefaultRooms), "Number of rooms available in app")
	flag.Parse() // parse the flags

	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	var sessionPath = "/tmp/"
	store := filesystem.NewStore(sessionPath, []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private, roomsQuantity)

	router.Run(":8080")
}
