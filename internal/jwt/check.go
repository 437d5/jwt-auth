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
