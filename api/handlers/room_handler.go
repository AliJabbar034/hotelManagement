package handlers

import (
	"net/http"

	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/types"
	"github.com/alijabbar034/hotelManagement/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	room_storer dbstores.Room_storer
}

func NewRoomHandler(room_storer dbstores.Room_storer) *RoomHandler {

	return &RoomHandler{
		room_storer: room_storer,
	}
}

func (r *RoomHandler) CreateRoomHandler(c *gin.Context) {
	var room types.Room

	if err := c.BindJSON(&room); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	if err := types.ValidateRoomData(room); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	id, err := r.room_storer.CreateRoom(room)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Created successfully",
	})

}

func (r *RoomHandler) UpdateRoomHandler(c *gin.Context) {
	var room types.Room
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	if err := c.BindJSON(&room); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	updatedCount, err := r.room_storer.UpdateRoom(room, id)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"updatedCount": updatedCount,
		"message":      "Updated successfully",
	})
}

func (r *RoomHandler) GetRoomHandler(c *gin.Context) {
	id := c.Param("id")
	_id, _ := primitive.ObjectIDFromHex(id)
	room, err := r.room_storer.GetRoomById(_id)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, room)
}

func (r *RoomHandler) GetRoomsHandler(c *gin.Context) {

	rooms, err := r.room_storer.GetAllRooms()
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (r *RoomHandler) DeleteRoomHandler(c *gin.Context) {

	id := c.Param("id")
	_id, _ := primitive.ObjectIDFromHex(id)
	count, err := r.room_storer.DeleteRoom(_id)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count":   count,
		"message": "Delete room successfully"})
}
