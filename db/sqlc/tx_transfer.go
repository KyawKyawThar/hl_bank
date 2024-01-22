package db

import (
	"context"
	"fmt"
)

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of transfer transaction
type TransferTxResult struct {
	Transfers   Transfer `json:"transfers"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entries"`
	ToEntry     Entry    `json:"to_entries"`
}

var txKey = struct{}{}

// TransferTx performs money transaction from one account to another
// It creates transfer record,add account entries,update accounts balance
// within a single db transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		txName := ctx.Value(txKey)
		fmt.Println("txName from tx_transfer", txName)

		fmt.Println(txName, "create transfer")
		result.Transfers, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		// update Account
		fmt.Println(txName, "get account 1")
		acc1, err := q.GetAccountForUpdate(ctx, arg.FromAccountId)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")
		result.FromAccount, err = q.UpdateAccounts(ctx, UpdateAccountsParams{
			Balance: acc1.Balance - arg.Amount,
			ID:      acc1.ID,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "get account 2")
		acc2, err := q.GetAccountForUpdate(ctx, arg.ToAccountId)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 2")
		result.ToAccount, err = q.UpdateAccounts(ctx, UpdateAccountsParams{
			Balance: acc2.Balance + arg.Amount,
			ID:      acc2.ID,
		})

		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
