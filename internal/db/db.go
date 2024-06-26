package db

import (
	"context"
	"errors"
	"github.com/437d5/jwt-auth/internal/pswd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repo struct {
	DB *mongo.Client
}

func (r *Repo) CreateNewUser(ctx context.Context, user User) error {
	hashedPswd, err := pswd.Hash(user.Password)
	if err != nil {
		log.Print(err)
		return err
	}
	user.Password = hashedPswd

	collection := r.DB.Database("users").Collection("users")
	_, err = collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
	})
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (User, error) {
	collection := r.DB.Database("users").Collection("users")
	var user User
	err := collection.FindOne(ctx, bson.D{
		{Key: "name", Value: username},
	}).Decode(&user)
	if err != nil {
		log.Print(err)
		return User{}, err
	}
	return user, nil
}

func (r *Repo) UserExists(ctx context.Context, username, email string) (bool, error) {
	var user User
	collection := r.DB.Database("users").Collection("users")
	err := collection.FindOne(ctx, bson.M{"$or": []bson.M{
		{"username": username},
		{"email": email},
	}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
