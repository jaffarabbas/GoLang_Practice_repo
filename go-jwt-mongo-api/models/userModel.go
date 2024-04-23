package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Username     *string            `json:"username" validate:"required" bson:"username"`
	Email        *string            `json:"email" validate:"email,required" bson:"email"`
	Password     *string            `json:"password" bson:"password"`
	Token        *string            `json:"token" bson:"token"`
	User_Type    *string            `json:"user_type" bson:"user_type"`
	RefreshToken *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	User_Id      *string            `json:"user_id" bson:"user_id"`
}
