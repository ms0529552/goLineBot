package service

import (
	"context"

	"goLineBot/models"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveMessage(db *mongo.Database, message *models.Message) error {
	collection := db.Collection("messages")
	_, err := collection.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
