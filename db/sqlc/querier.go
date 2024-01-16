// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Accounts, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Accounts, error)
	UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Accounts, error)
}

var _ Querier = (*Queries)(nil)
