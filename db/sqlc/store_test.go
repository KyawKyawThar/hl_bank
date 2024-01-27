package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferTx(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	fmt.Println(">> before:", acc1.Balance, acc2.Balance)

	//run n concurrent transaction --> for debug transaction deadlock use 2
	n := 5

	amount := int64(5)

	//check result

	existed := make(map[int]bool)

	results := make(chan TransferTxResult)
	errs := make(chan error)

	for i := 0; i < n; i++ {

		//txName: for testing deadLock
		// txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			// ctx := context.WithValue(context.Background(), txKey, txName)
			// result, err := testStore.TransferTx(ctx, TransferTxParams{
			result, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: acc1.ID,
				ToAccountId:   acc2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check errors
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfers
		require.NotEmpty(t, transfer)
		require.Equal(t, acc1.ID, transfer.FromAccountID)
		require.Equal(t, acc2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry.ID)
		require.Equal(t, acc1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)

		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry.ID)
		require.Equal(t, acc1.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)

		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount.ID)
		require.Equal(t, acc1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount.ID)
		require.Equal(t, acc2.ID, toAccount.ID)

		//check account's balance
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)

		diff1 := acc1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - acc2.Balance
		diff3 := acc2.Balance - toAccount.Balance

		fmt.Println("different acc", diff1, diff2, diff3)
		fmt.Println("amount is:", amount)
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount,2*amount,3*amount,....,n*amount

		k := int(diff1 / amount)
		fmt.Println("k is:", k)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balance
	updatedAccount1, err := testStore.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, acc1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, acc2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	fmt.Println(">> before:", acc1.Balance, acc2.Balance)

	//run n concurrent transaction --> for debug transaction deadlock use 2
	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {

		fromAccountID := acc1.ID
		toAccountID := acc2.ID

		if i%2 == 1 {
			fromAccountID = acc2.ID
			toAccountID = acc1.ID
		}

		go func() {

			_, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccountID,
				ToAccountId:   toAccountID,
				Amount:        amount,
			})

			errs <- err

		}()
	}
	fmt.Println("code running...")
	// check errors
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

	}

	// check the final updated balance
	updatedAccount1, err := testStore.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, acc1.Balance, updatedAccount1.Balance)
	require.Equal(t, acc2.Balance, updatedAccount2.Balance)
}
