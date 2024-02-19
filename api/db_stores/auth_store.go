package dbstores

import (
	"context"

	"github.com/alijabbar034/hotelManagement/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth_store interface {
	FindUser(string) (*types.User, error)
}

type AuthMongo struct {
	coll *mongo.Collection
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{
		coll: db.Collection("user"),
	}
}

func (auth *AuthMongo) FindUser(id string) (*types.User, error) {

	_id, _ := primitive.ObjectIDFromHex(id)
	var user types.User
	err := auth.coll.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
