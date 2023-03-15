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

var DBclient *mongo.Client

func ConnetDB(dbAdress string) {
	var err error
	DBclient, err = mongo.NewClient(options.Client().ApplyURI(dbAdress))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = DBclient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check if successfully connected
	err = DBclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongoDB connected at: " + dbAdress)
}

//For now there's no need to use below function, however, keep it for future.

// func GetDBClient(dbAdress string) *mongo.Client {
// 	if DBclient == nil {
// 		ConnetDB(dbAdress)
// 	}
// 	return DBclient
// }
