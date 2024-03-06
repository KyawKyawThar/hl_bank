package util

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {

	password := RandomString(7)
	hashPassword1, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashPassword1)

	err = CompareHahPassword(hashPassword1, password)

	require.NoError(t, err)

	wrongPassword := RandomString(7)
	err = CompareHahPassword(hashPassword1, wrongPassword)
	fmt.Println("err is:", err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword2)
	require.NotEqual(t, hashPassword1, hashPassword2)
}
