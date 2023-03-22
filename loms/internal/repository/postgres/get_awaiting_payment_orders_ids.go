package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (r *LomsRepository) GetAwaitingPaymentOrdersIDs(ctx context.Context) ([]models.OrderTimestamp, error) {
	const query = `
	SELECT id, create_at
	FROM orders
	WHERE status = $1;
	`

	rows, err := r.pool.Query(ctx, query, models.AwaitingPayment)
	if err != nil {
		return nil, fmt.Errorf("postgres GetAwaitingPaymentOrdersIDs select: %w", err)
	}
	defer rows.Close()

	orders := make([]models.OrderTimestamp, 0, 5)
	for rows.Next() {
		var orderInfo models.OrderTimestamp
		err = rows.Scan(&orderInfo.ID, &orderInfo.CreateAt)
		if err != nil {
			return nil, fmt.Errorf("postgres GetAwaitingPaymentOrdersIDs scan: %w", err)
		}
		orders = append(orders, orderInfo)
	}

	return orders, rows.Err()
}
