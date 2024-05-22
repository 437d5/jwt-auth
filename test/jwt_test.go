package test

import (
	myjwt "github.com/437d5/jwt-auth/internal/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTokenValid(t *testing.T) {
	actual1, err := myjwt.CreateToken("5")
	assert.NoError(t, err)

	actual2, err := myjwt.CreateToken("5")
	assert.NoError(t, err)

	// Checking that output token is not null
	assert.NotEmpty(t, actual1)
	assert.NotEmpty(t, actual2)

	// Checking idempotency of CreateToken function
	assert.Equal(t, actual1, actual2)
}

func TestCreateTokenInvalid(t *testing.T) {
	_, err := myjwt.CreateToken("")
	assert.Error(t, err)

	_, err = myjwt.CreateToken("5adf")
	assert.Error(t, err)
}

func TestParseTokenValid(t *testing.T) {
	token, _ := myjwt.CreateToken("5")
	// TODO secretkey
	parsedToken, err := myjwt.ParseToken(token, []byte(myjwt.SecretKey))
	assert.NoError(t, err)

	// Checking that parsedToken is not nil
	assert.NotNil(t, parsedToken)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("could not parse claims")
	}
	parsedID, ok := claims["jti"].(string)
	if !ok {
		t.Error("could not parse id")
	}

	// Checking if id in token equals id that we passed to CreateToken function
	assert.Equal(t, "5", parsedID)
}

func TestGetIDFromTokenValid(t *testing.T) {
	token, _ := myjwt.CreateToken("5")
	// TODO secret key
	parsedToken, _ := myjwt.ParseToken(token, []byte(myjwt.SecretKey))

	id, err := myjwt.GetIDFromToken(parsedToken)
	assert.NoError(t, err)

	assert.Equal(t, "5", id)

	id1, err := myjwt.GetIDFromToken(parsedToken)
	assert.NoError(t, err)

	assert.Equal(t, id, id1)
}

func TestGetIDFromTokenInvalid(t *testing.T) {
	id, err := myjwt.GetIDFromToken(nil)
	assert.Error(t, err)
	assert.Equal(t, "", id)
}
