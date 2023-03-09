package postgres

import (
	"context"
	"route256/loms/internal/models"
)

func (r *LomsRepository) ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error) {
	query := `SELECT o.status, o.user_id, item_sku, items_count
	FROM orders o
	INNER JOIN orders_set os on o.order_id = os.order_id
	WHERE o.order_id = $1;`

	rows, err := r.pool.Query(ctx, query, orderID)
	if err != nil {
		return models.OrderInfo{}, err
	}

	items := make([]models.Item, 0, 5)
	orderInfo := models.OrderInfo{
		Items: items,
	}
	for rows.Next() {
		item := models.Item{}
		err = rows.Scan(&orderInfo.Status, &orderInfo.User, &item.Sku, &item.Count)
		if err != nil {
			return models.OrderInfo{}, err
		}
		orderInfo.Items = append(orderInfo.Items, item)
	}
	rows.Close()

	return orderInfo, nil
}
