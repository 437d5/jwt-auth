package validations

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

// ValidateEmail checks if the provided email is valid
func ValidateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// validateUsername checks if the provided username is valid
func validateUsername(username string) bool {
	if len(username) < 5 || len(username) > 16 {
		return false
	}

	var validUsernameChars = regexp.MustCompile(`^[a-zA-Z0-9._]+$`)
	if !validUsernameChars.MatchString(username) {
		return false
	}

	if strings.HasPrefix(username, "_") || strings.HasPrefix(username, ".") ||
		strings.HasSuffix(username, "_") || strings.HasSuffix(username, ".") {
		return false
	}

	if strings.Contains(username, "__") || strings.Contains(username, "..") ||
		strings.Contains(username, "._") || strings.Contains(username, "_.") {
		return false
	}

	return true
}

// validatePassword checks if the provided password is valid
func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLetter := false
	hasDigit := false

	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	return hasLetter && hasDigit
}

func ValidatePasswordUsername(username, password string) bool {
	ok := validatePassword(password)
	if !ok {
		log.Print("Password validation failed")
		return false
	}
	ok = validateUsername(username)
	if !ok {
		log.Print("Username validation failed")
	}
	return ok
}
