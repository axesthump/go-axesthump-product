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
	WHERE send_status = $1
	ORDER BY id;
	`

	rows, err := r.pool.Query(ctx, query, models.Open)
	if err != nil {
		return nil, fmt.Errorf("GetOutboxOrders: %w", err)
	}
	defer rows.Close()
	outboxOrders := make([]models.OutboxOrder, 0, 10)
	for rows.Next() {
		outboxOrder := models.OutboxOrder{}
		err = rows.Scan(&outboxOrder.ID, &outboxOrder.OrderID, &outboxOrder.Status)
		if err != nil {
			return nil, fmt.Errorf("GetOutboxOrders Scan: %w", err)
		}
		outboxOrders = append(outboxOrders, outboxOrder)
	}
	ids := make([]int64, len(outboxOrders))
	for i, order := range outboxOrders {
		ids[i] = order.ID
	}
	err = r.UpdateOutboxOrders(ctx, ids, models.InProgress)
	if err != nil {
		return nil, fmt.Errorf("UpdateOutboxOrders: %w", err)
	}
	return outboxOrders, rows.Err()
}
