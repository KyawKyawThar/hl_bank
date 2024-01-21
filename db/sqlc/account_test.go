package db

import (
	"context"
	"github.com/hl/hl_bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Currency: util.RandomCurrency(),
		Balance:  int64(util.RandomAmount()),
	}

	acc, err := testStore.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.Balance)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	acc, err := testStore.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc1.ID, acc.ID)
	require.Equal(t, acc1.Owner, acc.Owner)
	require.Equal(t, acc1.Balance, acc.Balance)
	require.Equal(t, acc1.Currency, acc.Currency)

	require.WithinDuration(t, acc1.CreatedAt, acc.CreatedAt, time.Second)
}
