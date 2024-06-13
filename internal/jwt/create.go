// Package jwt is a package to create and sing new tokens
package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"time"
)

type TokenClaims struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	jwt.RegisteredClaims
}

// CreateToken func creates new JWT token with provided id of type primitive.ObjectID
// It returns JWT of string type and error if it occurs
func CreateToken(id primitive.ObjectID) (string, error) {
	log.Print("Creating token")
	if id == primitive.NilObjectID {
		log.Print("empty id")
		return "", errors.New("id is required")
	}

	now := time.Now()
	// TODO change exp time to config
	exp := now.Add(time.Hour)

	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		return "", errors.New("cannot get SECRET_KEY variable")
	}
	mySecretKey := []byte(secretKey)
	claims := TokenClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Print("New JWT Token was generated")

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		log.Print(err)
		return "", err
	}
	log.Print("New JWT Token was generated and signed")

	return tokenString, nil
}
