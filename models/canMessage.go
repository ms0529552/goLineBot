package models

import "time"

type CanMessage struct {
	ID          string    `json:"id" bson:"id"`
	Description string    `json:"description" bson:"description"`
	Content     string    `json:"contetnt" bson:"content"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
