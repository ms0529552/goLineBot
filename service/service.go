package service

import (
	"context"
	//"log"
	"time"

	"goLineBot/models"
	db "goLineBot/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

// Save message into the collection messages.
func SaveMessage(message *models.Message) error {
	messagesCollection := db.DBclient.Database("goLineBot").Collection("messages")
	_, err := messagesCollection.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}

	usersCollection := db.DBclient.Database("goLineBot").Collection("users")

	var user models.User
	filter := bson.M{"user_id": message.UserID}
	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		newUser := models.User{UserID: message.UserID, CreatedAt: time.Now()}
		NewUser(&newUser)
	}

	return err
}

// Insert new user into the collection users.
func NewUser(user *models.User) error {
	collection := db.DBclient.Database("goLineBot").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return err
}

func FindUserById(searchId string) (*models.User, error) {
	usersCollection := db.DBclient.Database("goLineBot").Collection("users")
	var user *models.User
	filter := bson.M{"user_id": searchId}
	err := usersCollection.FindOne(context.Background(), filter).Decode(&user)
	return user, err

}

// Get all the user in tho collection users
func GetUsersList() ([]models.User, error) {
	usersCollection := db.DBclient.Database("goLineBot").Collection("users")

	var usersList []models.User
	filter := bson.M{}
	cursor, err := usersCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return usersList, err
}

func GetMessagesByUser(user *models.User) ([]models.Message, error) {
	messagesCollection := db.DBclient.Database("goLineBot").Collection("messages")

	var messagesByUser []models.Message
	filter := bson.M{"user_id": user.UserID}
	cursor, err := messagesCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messagesByUser = append(messagesByUser, message)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return messagesByUser, err

}

func GetAllMessages() ([]models.Message, error) {
	messagesCollection := db.DBclient.Database("goLineBot").Collection("messages")

	var messagesList []models.Message
	filter := bson.M{}
	cursor, err := messagesCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messagesList = append(messagesList, message)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return messagesList, err

}
