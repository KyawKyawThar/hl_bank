package db

import (
	"context"
	"github.com/hl/hl_bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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
func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entry := createRandomEntry(t, acc.ID)

	entry1, err := testStore.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.Equal(t, entry.Amount, entry1.Amount)

	require.WithinDuration(t, entry.CreatedAt, entry1.CreatedAt, time.Second)

}

func TestEntryList(t *testing.T) {

	acc := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc.ID)
	}

	arg := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testStore.ListEntries(context.Background(), arg)

	require.Len(t, entries, 5)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
	argErr := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     -1,
		Offset:    -1,
	}

	_, errEntry := testStore.ListEntries(context.Background(), argErr)

	require.Error(t, errEntry)
}
