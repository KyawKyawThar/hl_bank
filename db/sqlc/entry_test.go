package db

import (
	"context"
	"github.com/hl/hl_bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEntry(t *testing.T, accID int64) Entry {

	arg := CreateEntryParams{Amount: util.RandomAmount(), AccountID: accID}
	entry, err := testStore.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc.ID)
}
