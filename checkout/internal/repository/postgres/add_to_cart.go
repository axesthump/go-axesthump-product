package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (r *CheckoutRepository) AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	queryCart := "SELECT cart_id FROM carts WHERE user_id = $1;"

	row := r.pool.QueryRow(ctx, queryCart, user)
	var cartID int64
	err := row.Scan(&cartID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return r.createCart(ctx, user, sku, count)
		} else {
			return err
		}
	}

	queryCartItems := "SELECT cart_item_id FROM carts_items WHERE cart_id = $1 AND item_sku = $2"
	var cartItemID int64
	row = r.pool.QueryRow(ctx, queryCartItems, cartID, sku)
	err = row.Scan(&cartItemID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return r.insertCart(ctx, sku, count, cartID, nil)
		} else {
			return err
		}
	}

	queryUpdate := "UPDATE carts_items SET count = count+$1 WHERE cart_item_id = $2 AND item_sku = $3;"
	_, err = r.pool.Exec(ctx, queryUpdate, count, cartItemID, sku)
	return err
}

func (r *CheckoutRepository) createCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	query := "INSERT INTO carts (user_id) VALUES ($1) RETURNING cart_id;"
	row := tx.QueryRow(ctx, query, user)
	if err != nil {
		return err
	}
	var cartID int64
	err = row.Scan(&cartID)
	if err != nil {
		return err
	}

	err = r.insertCart(ctx, sku, count, cartID, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *CheckoutRepository) insertCart(ctx context.Context, sku uint32, count uint32, cartID int64, tx pgx.Tx) error {
	query := "INSERT INTO carts_items (item_sku, count, cart_id) VALUES ($1, $2, $3);"
	var err error
	if tx != nil {
		_, err = tx.Exec(ctx, query, sku, count, cartID)
	} else {
		_, err = r.pool.Exec(ctx, query, sku, count, cartID)
	}
	return err
}
