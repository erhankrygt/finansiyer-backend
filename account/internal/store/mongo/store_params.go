package mongostore

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Password    string
	Token       string
	IsActive    bool
	IsDeleted   bool
	CreatedAt   primitive.DateTime
}

type LoginFilter struct {
	PhoneNumber string
	Password    string
}
