package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"route256/loms/internal/models"
)

func (r *LomsRepository) CreateOrder(ctx context.Context, order models.OrderData) (orderID int64, err error) {
	err = r.inTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queryOrder := "INSERT INTO orders (status, user_id) VALUES ($1, $2) RETURNING id"
		err = tx.QueryRow(ctx, queryOrder, models.New, order.User).Scan(&orderID)
		if err != nil {
			return err
		}
		err = r.insertInOrderItems(ctx, tx, order, err, orderID)
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

func (r *LomsRepository) insertInOrderItems(ctx context.Context, tx pgx.Tx, order models.OrderData, err error, orderID int64) error {
	for _, item := range order.Items {
		queryOrderSet := "INSERT INTO order_items(sku, count, order_id) VALUES ($1, $2, $3)"
		_, err = tx.Exec(ctx, queryOrderSet, item.Sku, item.Count, orderID)
		if err != nil {
			return fmt.Errorf("postgres insertInOrderItems: %w", err)
		}
	}
	return nil
}
