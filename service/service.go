package service

import (
	"context"

	"goLineBot/models"
	db "goLineBot/mongo"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveMessage(message *models.Message) error {
	collection := db.DBclient.Database("goLineBot").Collection("messages")
	_, err := collection.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
