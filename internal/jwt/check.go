package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// ValidToken func get tokenString of type string, parses it and
// returns true if token is valid or false
func ValidToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	log.Print("Starting token parsing")

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Print("invalid signing method in token")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}

	if !token.Valid {
		log.Print("Invalid token")
		return nil, errors.New("invalid token")
	}

	// Type assert the claims to TokenClaims
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		log.Print("Invalid token claims")
		return nil, errors.New("could not parse claims")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		log.Print("Token expired")
		return nil, errors.New("token has expired")
	}

	log.Print("Token parsed successfully")
	return token, nil
}

// GetIDFromToken func takes parsed token *jwt.Token
// and returns id string and error
func GetIDFromToken(token *jwt.Token) (primitive.ObjectID, error) {
	if token == nil {
		log.Print("Token is nil")
		return primitive.NilObjectID, errors.New("token is nil")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		log.Print("Invalid token claims")
		return primitive.NilObjectID, fmt.Errorf("couldn't get claims from token")
	}
	id := claims.ID

	return id, nil
}
