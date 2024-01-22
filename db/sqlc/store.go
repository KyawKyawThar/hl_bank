package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// composition is the prefer ways to extend struct functionally in golang
// instead of inheritance

// SQLStore provide all functionally to execute db queries and transaction
type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

// NewStore create a new store
func NewStore(connPool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
