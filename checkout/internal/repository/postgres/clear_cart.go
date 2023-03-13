package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (r *CheckoutRepository) ClearCart(ctx context.Context, user int64) error {
	err := r.inTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		const query = `
		SELECT id
		FROM carts 
		WHERE user_id = $1;
		`

		var cartID int64
		err := tx.QueryRow(ctx, query, user).Scan(&cartID)
		if err != nil {
			return err
		}

		err = r.deleteCart(ctx, tx, cartID)
		if err != nil {
			return err
		}

		err = r.deleteCartItems(ctx, tx, cartID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("postgres clear cart: %w", err)
	}
	return nil
}

func (r *CheckoutRepository) deleteCartItems(ctx context.Context, tx pgx.Tx, cartID int64) error {
	const query = `
	DELETE FROM cart_items 
	WHERE cart_id = $1;
	`

	_, err := tx.Exec(ctx, query, cartID)
	if err != nil {
		return fmt.Errorf("postgres deleteCartItems: %w", err)
	}
	return nil
}

func (r *CheckoutRepository) deleteCart(ctx context.Context, tx pgx.Tx, cartID int64) error {
	const query = `
	DELETE FROM carts 
	WHERE id = $1;
	`

	_, err := tx.Exec(ctx, query, cartID)
	if err != nil {
		return fmt.Errorf("postgres deleteCart: %w", err)
	}
	return nil
}
