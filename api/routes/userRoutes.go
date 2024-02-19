package routes

import (
	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/api/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(routes *gin.RouterGroup, db *mongo.Database) {
	userHandler := handlers.NewUserHandler(dbstores.NewUser_Mongo(db))
	r := routes.Group("/user")
	{
		r.POST("/", userHandler.RegisterUser)
		r.POST("/login", userHandler.LoginUserHandler)
		r.GET("/:id", userHandler.GetByIdHandler)
		r.GET("/", userHandler.GetAllUserHandler)
		r.DELETE("/:id", userHandler.DeleteUserHandler)
		r.GET("/logout", userHandler.LoginUserHandler)
		r.PUT("/me", userHandler.UpdateUserHandler)
		r.GET("/me", userHandler.GetProfileHandler)
	}
}
