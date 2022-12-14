package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	MsgId       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MessageBody string             `json:"message,omitempty" bson:"message,omitempty"`
	PhoneNumber string             `json:"phone_number,omitempty"  bson:"phone_number,omitempty"`
	TimeStamp   string             `json:"timeStamp,omitempty" bson:"timeStamp,omitempty"`
}
