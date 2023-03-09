package postgres

import (
	"context"
	"route256/loms/internal/models"
)

func (r *LomsRepository) CreateOrder(ctx context.Context, order models.OrderData) (orderID int64, err error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	queryOrder := "INSERT INTO orders (status, user_id) VALUES ($1, $2) RETURNING order_id"
	row := tx.QueryRow(ctx, queryOrder, models.New, order.User)

	err = row.Scan(&orderID)
	if err != nil {
		return 0, err
	}
	for _, item := range order.Items {
		queryOrderSet := "INSERT INTO orders_set(item_sku, items_count, order_id) VALUES ($1, $2, $3)"
		_, err = tx.Exec(ctx, queryOrderSet, item.Sku, item.Count, orderID)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}
