package models

import "time"

type SystemMessageLog struct {
	ID          string    `json:"id" bson:"id"`
	ReplyID     string    `json:"replyId" bson:"replyId"`
	ReplyUserID string    `json:"userId" bson:"userId"`
	Type        string    `json:"type" bson:"type"`
	Text        string    `json:"text" bson:"text"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
