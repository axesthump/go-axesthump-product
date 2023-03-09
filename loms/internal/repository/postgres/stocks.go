package postgres

import (
	"context"
	"errors"
	"route256/loms/internal/models"

	"github.com/jackc/pgx/v4"
)

func (r *LomsRepository) Stocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	query := `SELECT warehouse_id, available 
	FROM warehouses_items 
	WHERE item_sku = $1 AND available > 0;`

	rows, err := r.pool.Query(ctx, query, sku)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Stock{}, nil
		}
		return nil, err
	}

	stocks := make([]models.Stock, 0, 5)
	for rows.Next() {
		var warehouseID int64
		var count uint64

		err = rows.Scan(&warehouseID, &count)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, models.Stock{
			WarehouseID: warehouseID,
			Count:       count,
		})
	}
	rows.Close()
	return stocks, nil
}
