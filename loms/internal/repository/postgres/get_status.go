package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (r *LomsRepository) GetStatus(ctx context.Context, orderID int64) (models.OrderStatusID, error) {
	query := `SELECT status FROM orders WHERE id = $1;`
	var status int
	err := r.pool.QueryRow(ctx, query, orderID).Scan(&status)
	if err != nil {
		return 0, fmt.Errorf("postgres GetStatus: %w", err)
	}
	return models.OrderStatusID(status), nil
}
