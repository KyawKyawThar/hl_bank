package db

import (
	"context"
	"fmt"
	"github.com/hl/hl_bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {

	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
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

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		Balance: util.RandomAmount(),
		ID:      acc1.ID,
	}

	acc2, err := testStore.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, arg.ID, acc2.ID)
	require.Equal(t, arg.Balance, acc2.Balance)

	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Currency, acc2.Currency)

	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)

}

func TestListAccount(t *testing.T) {

	//var lastAccount Account
	//
	//for i := 0; i < 10; i++ {
	//	lastAccount = createRandomAccount(t)
	//	//fmt.Printf("last account owner for %s\n", lastAccount.Owner)
	//}

	// fmt.Printf("last account owner %s\n", lastAccount.Owner)
	arg := ListAccountsParams{
		//Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accLists, err := testStore.ListAccounts(context.Background(), arg)

	fmt.Println("accLists..........", accLists)
	require.NoError(t, err)
	require.NotEmpty(t, accLists)

	//for _, acc := range accLists {
	//	 require.Equal(t, lastAccount.Owner, acc.Owner)
	//	require.Equal(t, lastAccount.Balance, acc.Balance)
	//
	//	require.NotEmpty(t, acc)
	//}
	_, queryErr := testStore.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  -1,
		Offset: -1,
	})

	require.Error(t, queryErr)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)

	err := testStore.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)

	acc2, err := testStore.GetAccount(context.Background(), acc.ID)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, acc2)
	require.Error(t, err)
}
