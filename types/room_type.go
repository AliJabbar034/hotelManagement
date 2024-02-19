package types

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Number    string             `json:"number,omitempty" bson:"number,omitempty"`
	Capacity  int                `json:"capacity,omitempty" bson:"capacity,omitempty"`
	Price     float64            `json:"price,omitempty" bson:"price,omitempty"`
	Available bool               `json:"available,omitempty" bson:"available,omitempty"`
}

func ValidateRoomData(RoomData Room) error {

	if RoomData.Number == "" || RoomData.Capacity == 0 || RoomData.Price == 0 {
		return errors.New("Invalid room data")
	}
	return nil
}
