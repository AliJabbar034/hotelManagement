package routes

import (
	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/api/handlers"
	"github.com/alijabbar034/hotelManagement/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func BookingRoutes(routes *gin.RouterGroup, db *mongo.Database) {
	bookingHandler := handlers.New_BookingHandler(dbstores.NewMongo_Booking(db))
	Auth := middleware.NewAuth(dbstores.NewAuthMongo(db))
	r := routes.Group("/booking")
	r.POST("/", Auth.Authenticate, middleware.Authorize("admin"), bookingHandler.CreateBookingHandler)
	r.GET("/", Auth.Authenticate, middleware.Authorize("admin"), bookingHandler.GetAllBookingHandler)
	r.GET("/:id", Auth.Authenticate, middleware.Authorize("admin"), bookingHandler.GetBookingHandler)
	r.DELETE("/:id", Auth.Authenticate, middleware.Authorize("admin"), bookingHandler.DeleteBookingHandler)
}
