package models

import "time"

type User struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	//RecentInteract time.Time `json:"recent_interact" bson:"recent_interact"`
}
