package db

import (
	"context"
	"fmt"
)

// execTx create function within a transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(q *Queries) error) error {

	tx, err := store.connPool.Begin(ctx)

	if err != nil {
		return nil
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error %v, rb error %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}
