package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (r *LomsRepository) GetOutboxOrders(ctx context.Context) ([]models.OutboxOrder, error) {
	const query = `
	SELECT id, order_id, status
	FROM outbox_orders
	WHERE is_send = false
	ORDER BY id;
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("GetOutboxOrders: %w", err)
	}
	outboxOrders := make([]models.OutboxOrder, 0, 10)
	for rows.Next() {
		outboxOrder := models.OutboxOrder{}
		err = rows.Scan(&outboxOrder.ID, &outboxOrder.OrderID, &outboxOrder.Status)
		if err != nil {
			return nil, fmt.Errorf("GetOutboxOrders Scan: %w", err)
		}
		outboxOrders = append(outboxOrders, outboxOrder)
	}
	return outboxOrders, nil
}
