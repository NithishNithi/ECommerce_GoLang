// This file is to define the connection strings
// Port
// Datbase conenctivity
package config

import (
	"context"
	"fmt"

	//"fmt"
	"log"
	"time"

	"github.com/kishorens18/ecommerce/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func ConnectDataBase() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	mongoConnection := options.Client().ApplyURI(constants.ConnectionString)
	MongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if err := MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	fmt.Println("Database Connected")
	return MongoClient, nil
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
