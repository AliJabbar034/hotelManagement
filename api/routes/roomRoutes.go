package routes

import (
	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/api/handlers"
	"github.com/alijabbar034/hotelManagement/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RoomRoutes(routes *gin.RouterGroup, db *mongo.Database) {
	roomHandler := handlers.NewRoomHandler(dbstores.NewRoom_Mongo(db))
	Auth := middleware.NewAuth(dbstores.NewAuthMongo(db))
	r := routes.Group("/room")
	r.POST("/", Auth.Authenticate, middleware.Authorize("admin"), roomHandler.CreateRoomHandler)
	r.GET("/", roomHandler.GetRoomsHandler)
	r.GET("/:id", roomHandler.GetRoomHandler)
	r.PUT("/:id", Auth.Authenticate, middleware.Authorize("admin"), roomHandler.UpdateRoomHandler)
	r.DELETE("/:id", Auth.Authenticate, middleware.Authorize("admin"), roomHandler.DeleteRoomHandler)

}
