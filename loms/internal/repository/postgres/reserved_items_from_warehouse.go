package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type reservedItemsInWarehouses struct {
	itemSku     int64
	warehouseID int64
	itemCount   int64
}

func (r *LomsRepository) getReservedItemsInWarehousesFromOrder(
	ctx context.Context,
	tx pgx.Tx,
	orderID int64,
) ([]reservedItemsInWarehouses, error) {
	query := `SELECT wi.sku, wi.warehouse_id, osciw.count
	FROM warehouses_items wi
	INNER JOIN order_items oi on oi.order_id = $1
	INNER JOIN order_items_count_in_warehouse osciw on oi.id = osciw.order_items_id
	WHERE wi.sku = oi.sku AND osciw.warehouse_id = wi.warehouse_id;`

	rows, err := tx.Query(ctx, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("postgres getReservedItemsInWarehousesFromOrder select: %w", err)
	}

	reservedItemsInWarehousesRes := make([]reservedItemsInWarehouses, 0, 5)
	for rows.Next() {
		var reservedItems reservedItemsInWarehouses
		err = rows.Scan(&reservedItems.itemSku, &reservedItems.warehouseID, &reservedItems.itemCount)
		if err != nil {
			return nil, fmt.Errorf("postgres getReservedItemsInWarehousesFromOrder scan: %w", err)
		}
		reservedItemsInWarehousesRes = append(reservedItemsInWarehousesRes, reservedItems)
	}
	return reservedItemsInWarehousesRes, nil
}
