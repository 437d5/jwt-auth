package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	DB *mongo.Client
}

func (r *Repo) CreateNewUser(ctx context.Context, user User) error {
	collection := r.DB.Database("users").Collection("users")
	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
	})
	if err != nil {
		return err
	}
	return nil
}
