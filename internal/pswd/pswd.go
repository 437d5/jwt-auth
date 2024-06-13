package pswd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(hash), nil
}

func Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}
