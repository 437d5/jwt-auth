// Package jwt is a package to create and sing new tokens
package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strconv"
	"time"
)

// TODO delete secretKey
const SecretKey = "secret"

// CreateToken func creates new JWT token with provided id of type string
// It returns JWT of string type and error if it occurs
func CreateToken(id string) (string, error) {
	log.Print("Creating token")
	if id == "" {
		return "", errors.New("id is required")
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}
	now := time.Now()
	// TODO change exp time to config
	exp := now.Add(time.Hour)
	// TODO change secret to config secret
	mySecretKey := []byte(SecretKey)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Print("New JWT Token was generated")

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}
	log.Print("New JWT Token was generated and signed")

	return tokenString, nil
}
