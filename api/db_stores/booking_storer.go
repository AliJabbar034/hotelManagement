package dbstores

import (
	"context"
	"errors"

	"github.com/alijabbar034/hotelManagement/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Booking_Storer interface {
	CreateBooking(booking types.Booking) (string, error)
	GetBooking(primitive.ObjectID) (*types.Booking, error)

	DeleteBooking(primitive.ObjectID) (int64, error)
	GetAllBookings() ([]types.Booking, error)
}

type Mongo_Booking struct {
	collection *mongo.Collection
	roomColl   *mongo.Collection
}

func NewMongo_Booking(db *mongo.Database) *Mongo_Booking {
	return &Mongo_Booking{
		collection: db.Collection("bookings"),
		roomColl:   db.Collection("rooms"),
	}
}

func (m *Mongo_Booking) CreateBooking(booking types.Booking) (string, error) {
	var room types.Room
	filter := bson.M{"_id": booking.RoomID}

	err := m.roomColl.FindOne(context.Background(), filter).Decode(&room)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("No Room found with this id")
		}
		return "", err
	}
	inserted, err := m.collection.InsertOne(context.Background(), booking)
	if err != nil {
		return "", err
	}
	_, er := m.roomColl.UpdateOne(context.Background(), bson.M{"_id": booking.RoomID}, bson.M{"$set": bson.M{"Available": false}})
	if er != nil {
		return "", er
	}
	id := inserted.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (m *Mongo_Booking) GetAllBookings() ([]types.Booking, error) {

	var bookings []types.Booking

	cursor, err := m.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (m *Mongo_Booking) GetBooking(id primitive.ObjectID) (*types.Booking, error) {
	var bookings types.Booking

	if err := m.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&bookings); err != nil {
		return nil, err
	}

	return &bookings, nil
}

func (m *Mongo_Booking) DeleteBooking(id primitive.ObjectID) (int64, error) {
	var booking types.Booking
	err := m.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&booking)
	if err != nil {
		return 0, err
	}
	count, er := m.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if er != nil {
		return 0, er
	}
	_, er = m.roomColl.UpdateOne(context.Background(), bson.M{"_id": booking.RoomID}, bson.M{"$set": bson.M{"Available": true}})

	if er != nil {
		return 0, er
	}
	return count.DeletedCount, nil
}
