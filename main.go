package main

import (
	"github.com/alijabbar034/hotelManagement/api/routes"
	"github.com/alijabbar034/hotelManagement/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectToDb()
	app := gin.Default()
	api := app.Group("/api")
	routes.UserRoutes(api, db)

	app.Run(":3000")

}
