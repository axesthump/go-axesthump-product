package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	"time"

	"github.com/jackc/pgx/v4"
)

func (r *LomsRepository) CreateOrder(ctx context.Context, order models.OrderData) (orderID int64, err error) {
	err = r.inTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		const query = `
		INSERT INTO orders (status, create_at, user_id) 
		VALUES ($1, $2, $3) 
		RETURNING id;
		`

		err = tx.QueryRow(ctx, query, models.New, time.Now(), order.User).Scan(&orderID)
		if err != nil {
			return err
		}
		err = r.insertInOrderItems(ctx, tx, order, orderID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("postgres create order: %w", err)
	}
	return orderID, nil
}

func (r *LomsRepository) insertInOrderItems(ctx context.Context, tx pgx.Tx, order models.OrderData, orderID int64) error {
	for _, item := range order.Items {
		const query = `
		INSERT INTO order_items(sku, count, order_id)
		VALUES ($1, $2, $3);
		`

		_, err := tx.Exec(ctx, query, item.Sku, item.Count, orderID)
		if err != nil {
			return fmt.Errorf("postgres insertInOrderItems: %w", err)
		}
	}
	return nil
}
