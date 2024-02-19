package routes

import (
	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/api/handlers"
	"github.com/alijabbar034/hotelManagement/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(routes *gin.RouterGroup, db *mongo.Database) {
	userHandler := handlers.NewUserHandler(dbstores.NewUser_Mongo(db))
	Auth := middleware.NewAuth(dbstores.NewAuthMongo(db))
	r := routes.Group("/user")
	{
		r.POST("/", userHandler.RegisterUser)
		r.POST("/login", userHandler.LoginUserHandler)
		r.GET("/:id", Auth.Authenticate, middleware.Authorize("admin"), userHandler.GetByIdHandler)
		r.GET("/", Auth.Authenticate, middleware.Authorize("admin"), userHandler.GetAllUserHandler)
		r.DELETE("/:id", Auth.Authenticate, middleware.Authorize("admin"), userHandler.DeleteUserHandler)
		r.GET("/logout", Auth.Authenticate, userHandler.LoginUserHandler)
		r.PUT("/me", Auth.Authenticate, userHandler.UpdateUserHandler)
		r.GET("/me", Auth.Authenticate, userHandler.GetProfileHandler)
	}
}
