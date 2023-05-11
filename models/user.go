package models

import (
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Profile linebot.UserProfileResponse

type User struct {
	UserID        string    `json:"userId" bson:"userId"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	Profile       Profile
	ChatGptSwitch bool `json:"chatGptSwitch" bson:"chatGptSwitch"`
}

func (user User) CheckGptSwitch() bool {
	return user.ChatGptSwitch
}
