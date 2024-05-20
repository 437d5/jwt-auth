package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

// ParseToken func parses the provided tokenString
// and returns *jwt.Token and error
func ParseToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	log.Print("Starting token parsing")
	// TODO change secretKey to config
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	log.Print("Token parsed")
	return token, nil
}

// GetIDFromToken func takes parsed token *jwt.Token
// and returns id string and error
func GetIDFromToken(token *jwt.Token) (id string, err error) {
	if token == nil {
		return "", errors.New("token is nil")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("couldn't get claims from token: %w", err)
	}
	id, ok = claim["jti"].(string)
	if !ok {
		return "", fmt.Errorf("couldn't get id from token")
	}

	return id, nil
}
