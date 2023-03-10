package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (r *LomsRepository) ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error) {
	const query = `
	SELECT o.status, o.user_id, sku, count
	FROM orders o
	INNER JOIN order_items oi on o.id = oi.order_id
	WHERE o.id = $1;
	`

	rows, err := r.pool.Query(ctx, query, orderID)
	if err != nil {
		return models.OrderInfo{}, fmt.Errorf("postgres ListOrder select: %w", err)
	}
	defer rows.Close()

	items := make([]models.Item, 0, 5)
	orderInfo := models.OrderInfo{
		Items: items,
	}
	for rows.Next() {
		item := models.Item{}
		err = rows.Scan(&orderInfo.Status, &orderInfo.User, &item.Sku, &item.Count)
		if err != nil {
			return models.OrderInfo{}, fmt.Errorf("postgres ListOrder scan: %w", err)
		}
		orderInfo.Items = append(orderInfo.Items, item)
	}

	return orderInfo, nil
}
