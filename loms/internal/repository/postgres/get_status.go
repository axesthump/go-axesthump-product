package postgres

import (
	"context"
	"route256/loms/internal/models"
)

func (r *LomsRepository) GetStatus(ctx context.Context, orderID int64) (models.OrderStatusID, error) {
	query := `SELECT status FROM orders WHERE order_id = $1;`
	row := r.pool.QueryRow(ctx, query, orderID)

	var status int
	err := row.Scan(&status)
	if err != nil {
		return 0, err
	}
	return models.OrderStatusID(status), nil
}
