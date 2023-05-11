package models

import "time"

type SystemMessageLog struct {
	ReplyID     string    `json:"id" bson:"_id"`
	ReplyUserID string    `json:"userId" bson:"userId"`
	Type        string    `json:"type" bson:"type"`
	Text        string    `json:"text" bson:"text"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
