package db

import (
	"context"
	"fmt"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnetDB(dbAdress string) {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(dbAdress))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check if successfully connected
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongoDB connected at: " + dbAdress)
}

func GetDBClient(dbAdress string) *mongo.Client {
	if client == nil {
		ConnetDB(dbAdress)
	}
	return client
}
