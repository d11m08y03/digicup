package auth_test

import (
	"testing"

	"github.com/d11m08y03/digicup/auth"
	"github.com/stretchr/testify/assert"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "securepassword"

	hashedPassword, err := auth.HashPassword(password)
	assert.NoError(t, err, "Expected no error while hashing password")

	isMatch := auth.CheckPasswordHash(password, hashedPassword)
	assert.True(t, isMatch, "Expected the password to match the hashed password")

	wrongPassword := "wrongpassword"
	isMatch = auth.CheckPasswordHash(wrongPassword, hashedPassword)
	assert.False(t, isMatch, "Expected the password not to match the hashed password")
}
