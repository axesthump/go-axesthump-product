package postgres

import (
	"context"
	"errors"
	"fmt"
	"route256/loms/internal/models"

	"github.com/jackc/pgx/v4"
)

func (r *LomsRepository) Stocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	const query = `
	SELECT warehouse_id, available 
	FROM warehouses_items 
	WHERE sku = $1 AND available > 0;
	`

	rows, err := r.pool.Query(ctx, query, sku)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Stock{}, nil
		}
		return nil, fmt.Errorf("postgres Stocks select: %w", err)
	}

	stocks := make([]models.Stock, 0, 5)
	for rows.Next() {
		var warehouseID int64
		var count uint64

		err = rows.Scan(&warehouseID, &count)
		if err != nil {
			return nil, fmt.Errorf("postgres Stocks scan: %w", err)
		}
		stocks = append(stocks, models.Stock{
			WarehouseID: warehouseID,
			Count:       count,
		})
	}
	return stocks, nil
}
