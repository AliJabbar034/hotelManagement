package types

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RoomID       primitive.ObjectID `json:"room_id,omitempty" bson:"room_id,omitempty"`
	UserID       primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CheckInDate  time.Time          `json:"check_in_date,omitempty" bson:"check_in_date,omitempty"`
	CheckOutDate time.Time          `json:"check_out_date,omitempty" bson:"check_out_date,omitempty"`
	Price        float64            `json:"price,omitempty" bson:"price,omitempty"`
}

func ValidateBooking(booking Booking) error {

	if booking.RoomID.IsZero() || booking.UserID.IsZero() || booking.CheckInDate.IsZero() || booking.CheckOutDate.IsZero() || booking.Price == 0 {
		return errors.New("please fill in all fields")
	}
	return nil

}
