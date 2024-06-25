package db

import (
	"context"
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

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (User, error) {
	collection := r.DB.Database("users").Collection("users")
	var user User
	err := collection.FindOne(ctx, bson.D{
		{Key: "email", Value: email},
	}).Decode(&user)
	if err != nil {
		log.Print(err)
		return User{}, err
	}
	return user, nil
}

func (r *Repo) UserExists(ctx context.Context, username, email string) bool {
	var usernameFlag, emailFlag bool
	_, err := r.GetUserByUsername(ctx, username)
	if err == nil {
		usernameFlag = true
	}

	_, err = r.GetUserByEmail(ctx, email)
	if err == nil {
		emailFlag = true
	}

	return usernameFlag || emailFlag
}
