package handlers

import (
	"net/http"

	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/types"
	"github.com/alijabbar034/hotelManagement/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingHandler struct {
	booking_storer dbstores.Booking_Storer
}

func New_BookingHandler(db dbstores.Booking_Storer) *BookingHandler {

	return &BookingHandler{
		booking_storer: db,
	}
}

func (b *BookingHandler) CreateBookingHandler(c *gin.Context) {

	var booking types.Booking

	if err := c.BindJSON(&booking); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	if err := types.ValidateBooking(booking); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	userid, _ := primitive.ObjectIDFromHex(booking.UserID.Hex())
	roomId, _ := primitive.ObjectIDFromHex(booking.RoomID.Hex())
	booking.UserID = userid
	booking.RoomID = roomId

	id, err := b.booking_storer.CreateBooking(booking)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"message": "Created successfully",
	})
}

func (b *BookingHandler) GetBookingHandler(c *gin.Context) {
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	booking, err := b.booking_storer.GetBooking(id)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, booking)

}

func (b *BookingHandler) GetAllBookingHandler(c *gin.Context) {

	bookigs, err := b.booking_storer.GetAllBookings()
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, bookigs)
}

func (b *BookingHandler) DeleteBookingHandler(c *gin.Context) {
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)
	count, err := b.booking_storer.DeleteBooking(id)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count":   count,
		"message": "Delete booking successfully"})
}
