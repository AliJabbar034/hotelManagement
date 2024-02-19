package dbstores

import (
	"context"

	"github.com/alijabbar034/hotelManagement/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User_Storer interface {
	RegisterUser(user *types.User) (string, error)

	GetUserByEmail(email string) (*types.User, error)
	GetUserById(id primitive.ObjectID) (*types.User, error)
	GetAllUsers() ([]types.User, error)
	DeleteUser(primitive.ObjectID) (int64, error)
	UpdateUser(types.User, string) (int64, error)
}

type User_Mongo struct {
	collect *mongo.Collection
}

func NewUser_Mongo(db *mongo.Database) *User_Mongo {
	return &User_Mongo{
		collect: db.Collection("user"),
	}
}

func (m *User_Mongo) RegisterUser(user *types.User) (string, error) {

	inserted, err := m.collect.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	id := inserted.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (m *User_Mongo) GetUserByEmail(email string) (*types.User, error) {
	var user types.User

	if err := m.collect.FindOne(context.Background(), bson.M{"email": email}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {

			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (m *User_Mongo) GetUserById(id primitive.ObjectID) (*types.User, error) {

	var user types.User

	if err := m.collect.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {

			return nil, nil
		}
		return nil, err
	}
	return nil, nil
}

func (m *User_Mongo) GetAllUsers() ([]types.User, error) {

	var users []types.User
	cursor, err := m.collect.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	if er := cursor.All(context.Background(), &users); er == nil {
		return nil, er
	}
	return users, nil
}

func (m *User_Mongo) DeleteUser(id primitive.ObjectID) (int64, error) {

	filter := bson.M{"_id": id}
	result, err := m.collect.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (m *User_Mongo) UpdateUser(user types.User, id string) (int64, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.D{}
	if user.FirstName != "" {
		update = append(update, bson.E{
			"$set", bson.D{{
				"first_name", user.FirstName,
			}},
		})
	}
	if user.LastName != "" {
		update = append(update, bson.E{
			"$set", bson.D{{
				"last_name", user.LastName,
			}},
		})
	}
	updated, err := m.collect.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}

	return updated.ModifiedCount, nil
}
