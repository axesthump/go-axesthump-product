package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type reservedItemsInWarehouses struct {
	warehouseItemID int64
	itemCount       int64
}

func (r *LomsRepository) getReservedItemsInWarehousesFromOrder(
	ctx context.Context,
	tx pgx.Tx,
	orderID int64,
) ([]reservedItemsInWarehouses, error) {
	query := `SELECT wi.warehouse_items_id, osciw.item_count
	FROM warehouses_items wi
	INNER JOIN orders_set os on os.order_id = $1
	INNER JOIN orders_set_count_in_warehouse osciw on os.order_set_id = osciw.order_set_id
	WHERE wi.item_sku = os.item_sku AND osciw.warehouse_id = wi.warehouse_id;`

	rows, err := tx.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}

	reservedItemsInWarehousesRes := make([]reservedItemsInWarehouses, 0, 5)
	for rows.Next() {
		var reservedItems reservedItemsInWarehouses
		err = rows.Scan(&reservedItems.warehouseItemID, &reservedItems.itemCount)
		if err != nil {
			return nil, err
		}
		reservedItemsInWarehousesRes = append(reservedItemsInWarehousesRes, reservedItems)
	}
	return reservedItemsInWarehousesRes, nil
}
