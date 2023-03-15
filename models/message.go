package models

import "time"

type Message struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Type      string    `json:"type" bson:"type"`
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
