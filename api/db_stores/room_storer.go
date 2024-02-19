package dbstores

import (
	"context"

	"github.com/alijabbar034/hotelManagement/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Room_storer interface {
	CreateRoom(types.Room) (string, error)
	UpdateRoom(types.Room, primitive.ObjectID) (int64, error)
	GetRoomById(primitive.ObjectID) (*types.Room, error)
	GetAllRooms() ([]types.Room, error)
	DeleteRoom(primitive.ObjectID) (int64, error)
}

type MongoDB_Room struct {
	coll *mongo.Collection
}

func NewRoom_Mongo(db *mongo.Database) *MongoDB_Room {
	return &MongoDB_Room{
		coll: db.Collection("rooms"),
	}
}

func (m *MongoDB_Room) CreateRoom(room types.Room) (string, error) {
	room.Available = true
	inserted, err := m.coll.InsertOne(context.Background(), &room)
	if err != nil {
		return "", err

	}
	id := inserted.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (m *MongoDB_Room) UpdateRoom(room types.Room, id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	update := bson.D{}

	if room.Number != "" {
		update = append(update, bson.E{
			"$set", bson.D{{
				"name", room.Number,
			}},
		})
	}
	if room.Capacity != 0 {
		update = append(update, bson.E{
			"$set", bson.D{{
				"capacity", room.Capacity,
			}},
		})
	}
	if room.Price > 0 {
		update = append(update, bson.E{
			"$set", bson.D{{
				"price", room.Price,
			}},
		})
	}
	if room.Available {
		update = append(update, bson.E{
			"$set", bson.D{{
				"available", room.Available,
			}},
		})
	}

	result, err := m.coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil

}

func (m *MongoDB_Room) GetRoomById(id primitive.ObjectID) (*types.Room, error) {

	var room types.Room

	if err := m.coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(&room); err != nil {
		return nil, err
	}

	return &room, nil
}

func (m *MongoDB_Room) GetAllRooms() ([]types.Room, error) {
	var rooms []types.Room

	cursor, err := m.coll.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err

	}
	if err := cursor.All(context.Background(), &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (m *MongoDB_Room) DeleteRoom(id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	result, err := m.coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
