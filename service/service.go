package service

import (
	"context"
	"log"
	"time"

	"goLineBot/models"
	db "goLineBot/mongo"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson"
)

// Save message into the collection messages.
func SaveMessage(message *models.Message, bot *linebot.Client) error {
	messagesCollection := db.DBclient.Database("goLineBot").Collection("messages")
	_, err := messagesCollection.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}

	usersCollection := db.DBclient.Database("goLineBot").Collection("users")

	var user models.User

	filter := bson.M{"userId": message.UserID}
	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		var newUser models.User
		newUser.UserID = message.UserID
		newUser.ChatGptSwitch = false
		ctx := context.Background()
		userProfile, gettingProfieErr := bot.GetProfile(message.UserID).WithContext(ctx).Do()
		if gettingProfieErr != nil {
			return gettingProfieErr
		}
		newUser.Profile = models.Profile(*userProfile)
		newUser.CreatedAt = time.Now()
		NewUser(&newUser)
	}

	return err
}

func SaveSystemMessage(message models.SystemMessageLog) error {
	messagesCollection := db.DBclient.Database("goLineBot").Collection("system_messages_log")
	_, err := messagesCollection.InsertOne(context.Background(), message)
	if err != nil {
		log.Print(err)
		return err
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
	filter := bson.M{"userId": searchId}
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
	filter := bson.M{"userId": user.UserID}
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
func ChangeGptSwitch(userId string) {
	usersCollection := db.DBclient.Database("goLineBot").Collection("users")

	user, err := FindUserById(userId)

	filter := bson.M{"userId": bson.M{"$eq": userId}}
	update := bson.M{"$set": bson.M{"chatGptSwitch": true}}
	if user.ChatGptSwitch {
		update = bson.M{"$set": bson.M{"chatGptSwitch": false}}
	}

	_, err = usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

}

func FindCanMessagesById(id string) (*models.CanMessage, error) {
	collection := db.DBclient.Database("goLineBot").Collection("can_messages")

	var canMessage *models.CanMessage
	filter := bson.M{"id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&canMessage)
	return canMessage, err

}
