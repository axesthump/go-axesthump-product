package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	"github.com/jackc/pgx/v4"
)

func (r *LomsRepository) saveInOutbox(ctx context.Context, tx pgx.Tx, orderID int64, status models.OrderStatusID) error {
	const query = `
	INSERT INTO outbox_orders(order_id, status)
	VALUES ($1, $2);
	`

	_, err := tx.Exec(ctx, query, orderID, status)
	if err != nil {
		return fmt.Errorf("saveInOutbox: %w", err)
	}
	return nil
}
