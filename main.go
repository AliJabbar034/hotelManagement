package main

import (
	"github.com/alijabbar034/hotelManagement/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDb()
	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	app.Run(":3000")

}
