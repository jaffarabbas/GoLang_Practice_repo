package controllers

import (
	"go-jwt-mongo-api/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Register()

func Login()

func GetUsers()

func GetUser()
