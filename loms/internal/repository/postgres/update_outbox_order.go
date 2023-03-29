package postgres

import (
	"context"
	"fmt"
)

func (r *LomsRepository) UpdateOutboxOrder(ctx context.Context, outboxID int64) error {
	const query = `
	UPDATE outbox_orders
	SET is_send = true
	WHERE id = $1;
	`

	_, err := r.pool.Exec(ctx, query, outboxID)
	if err != nil {
		return fmt.Errorf("UpdateOutboxOrder: %w", err)
	}
	return nil
}
