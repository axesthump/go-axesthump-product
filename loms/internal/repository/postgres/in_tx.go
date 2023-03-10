package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (r *LomsRepository) inTx(ctx context.Context, f func(ctx context.Context, tx pgx.Tx) error) error {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	err = f(ctx, tx)
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return nil
}
