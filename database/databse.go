package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb() *mongo.Database {

	if err := godotenv.Load(); err != nil {
		panic("could not load database environment")
	}

	uri := os.Getenv("MONGO_URL")

	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	database := conn.Database("HotelManager")
	fmt.Println("Connection established with database ")
	return database

}
