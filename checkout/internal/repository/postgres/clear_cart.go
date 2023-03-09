package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (r *CheckoutRepository) ClearCart(ctx context.Context, user int64) (err error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	query := `SELECT cart_id FROM carts WHERE user_id = $1`
	row := tx.QueryRow(ctx, query, user)
	var cartID int64
	err = row.Scan(&cartID)
	if err != nil {
		return err
	}

	queryDeleteCart := `DELETE FROM carts where cart_id = $1`
	queryDeleteCartItems := `DELETE FROM carts_items where cart_id = $1`

	_, err = tx.Exec(ctx, queryDeleteCart, cartID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, queryDeleteCartItems, cartID)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
