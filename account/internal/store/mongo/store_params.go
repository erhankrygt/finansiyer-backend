package mongostore

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Password    string
	CreatedAt   primitive.DateTime
}