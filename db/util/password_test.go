package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword1, err := HashPassword(password)
	require.NotEmpty(t, hashedPassword1)
	require.NoError(t, err)

	err = CheckHashedPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckHashedPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(password)
	require.NotEmpty(t, hashedPassword2)
	require.NoError(t, err)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
