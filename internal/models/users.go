package models

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Name     string
	Email    string
	Password string
}

type UserModel struct {
	DB *mongo.Collection
}
