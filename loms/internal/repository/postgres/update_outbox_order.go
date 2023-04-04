package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (r *LomsRepository) UpdateOutboxOrders(ctx context.Context, outboxID []int64, status models.SendStatus) error {
	const query = `
	UPDATE outbox_orders
	SET send_status = $1
	WHERE id = ANY($2);
	`

	_, err := r.pool.Exec(ctx, query, status, outboxID)
	if err != nil {
		return fmt.Errorf("UpdateOutboxOrder: %w", err)
	}
	return nil
}
