package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (r *CheckoutRepository) AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error {

	queryCart := "SELECT id FROM carts WHERE user_id = $1;"
	var cartID int64
	err := r.pool.QueryRow(ctx, queryCart, user).Scan(&cartID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return r.createCart(ctx, user, sku, count)
		}
		return fmt.Errorf("postgres AddToCart: %w", err)
	}

	return r.addItems(ctx, err, cartID, sku, count)
}

func (r *CheckoutRepository) addItems(
	ctx context.Context,
	err error,
	cartID int64,
	sku uint32,
	count uint32,
) error {
	queryCartItems := "SELECT id FROM cart_items WHERE cart_id = $1 AND sku = $2"
	var cartItemID int64
	err = r.pool.QueryRow(ctx, queryCartItems, cartID, sku).Scan(&cartItemID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return r.insertCartItems(ctx, sku, count, cartID)
		}
		return fmt.Errorf("postgres addItems: %w", err)
	}

	return r.updateItems(ctx, err, count, cartItemID, sku)
}

func (r *CheckoutRepository) updateItems(ctx context.Context, err error, count uint32, cartItemID int64, sku uint32) error {
	queryUpdate := "UPDATE cart_items SET count = count+$1 WHERE id = $2 AND sku = $3;"
	_, err = r.pool.Exec(ctx, queryUpdate, count, cartItemID, sku)
	if err != nil {
		return fmt.Errorf("postgres updateItems: %w", err)
	}
	return nil
}

func (r *CheckoutRepository) createCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	err := r.inTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := "INSERT INTO carts (user_id) VALUES ($1) RETURNING id;"
		var cartID int64
		err := tx.QueryRow(ctx, query, user).Scan(&cartID)
		if err != nil {
			return fmt.Errorf("postgres createCart insert into carts: %w", err)
		}

		err = r.insertCartItemsTx(ctx, tx, sku, count, cartID)
		if err != nil {
			return fmt.Errorf("postgres createCart: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("postgres create cart: %w", err)
	}
	return nil
}

func (r *CheckoutRepository) insertCartItems(ctx context.Context, sku uint32, count uint32, cartID int64) error {
	query := "INSERT INTO cart_items (sku, count, cart_id) VALUES ($1, $2, $3);"
	var err error
	_, err = r.pool.Exec(ctx, query, sku, count, cartID)
	if err != nil {
		return fmt.Errorf("postgres insertCartItems: %w", err)
	}
	return nil
}

func (r *CheckoutRepository) insertCartItemsTx(ctx context.Context, tx pgx.Tx, sku uint32, count uint32, cartID int64) error {
	query := "INSERT INTO cart_items (sku, count, cart_id) VALUES ($1, $2, $3);"
	var err error
	_, err = tx.Exec(ctx, query, sku, count, cartID)
	if err != nil {
		return fmt.Errorf("postgres insertCartItemsTx: %w", err)
	}
	return nil
}
