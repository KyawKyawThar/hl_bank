package db

import (
	"context"
	"github.com/hl/hl_bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {

	password := util.RandomString(7)

	hashPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword)

	arg := CreateUserParams{
		Username: util.RandomOwner(),
		Password: hashPassword,
		Email:    util.RandomEmailAccount(),
		FullName: util.RandomOwner(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.FullName, user.FullName)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestQueries_GetUser(t *testing.T) {
	createUser := createRandomUser(t)

	user, err := testStore.GetUser(context.Background(), createUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createUser.Username, user.Username)
	require.Equal(t, createUser.Email, user.Email)
	require.Equal(t, createUser.Password, user.Password)
	require.Equal(t, createUser.FullName, user.FullName)

	require.WithinDuration(t, createUser.CreatedAt, user.CreatedAt, time.Second)
	require.WithinDuration(t, createUser.PasswordChangedAt, user.PasswordChangedAt, time.Second)
}
