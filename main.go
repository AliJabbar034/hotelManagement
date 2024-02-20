package main

import (
	"time"

	"github.com/alijabbar034/hotelManagement/api/routes"
	"github.com/alijabbar034/hotelManagement/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectToDb()
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000/"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := app.Group("/api")
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Gin API",
		})
	})
	routes.UserRoutes(api, db)
	routes.RoomRoutes(api, db)
	routes.BookingRoutes(api, db)

	app.Run(":3000")

}
