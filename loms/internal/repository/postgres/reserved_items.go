package postgres

import (
	"context"
	"log"
	"route256/loms/internal/models"

	"github.com/jackc/pgx/v4"
)

type warehouseReservedItems struct {
	warehouseID int64
	count       int64
	itemSku     int64
	orderSetID  int64
}

type orderSetData struct {
	sku        int64
	count      int64
	orderSetID int64
}

type warehouseItemData struct {
	available   int64
	warehouseID int64
}

func (r *LomsRepository) ReservedItems(ctx context.Context, orderID int64) {
	tx, err := r.pool.Begin(ctx)
	if err != nil { // todo ошибки пока не обрабатываются, так как будет этот метод исполнятся в другом потоке (реализация после следующего воркшопа)
		return
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	ordersSetData, err := r.getOrdersSetData(ctx, tx, orderID)
	if err != nil {
		return
	}

	possibleReservedItemsInfo, err := r.getPossibleReservedItemsInfo(ctx, tx, ordersSetData, orderID)
	if err != nil {
		return
	}
	if possibleReservedItemsInfo == nil {
		return
	}

	for _, wInfo := range possibleReservedItemsInfo {
		err = r.updateWarehousesItems(ctx, tx, wInfo)
		if err != nil {
			return
		}

		err = r.insertOrdersSetCountInWarehouse(ctx, tx, wInfo)
		if err != nil {
			return
		}
	}

	err = r.updateStatus(ctx, tx, models.AwaitingPayment, orderID)
	if err != nil {
		return
	}
	err = tx.Commit(ctx)
	if err != nil {
		return
	}
}

func (r *LomsRepository) getOrdersSetData(ctx context.Context, tx pgx.Tx, orderID int64) ([]orderSetData, error) {
	queryOrderSet := `SELECT item_sku, items_count, order_set_id
	FROM orders_set
	WHERE order_id = $1;`

	rows, err := tx.Query(ctx, queryOrderSet, orderID)
	if err != nil {
		log.Println("Query error on orders_set:", err.Error())
		return nil, err
	}

	ordersSetData := make([]orderSetData, 0)
	for rows.Next() {
		var orderSetInfo orderSetData

		err = rows.Scan(&orderSetInfo.sku, &orderSetInfo.count, &orderSetInfo.orderSetID)
		if err != nil {
			return nil, err
		}
		ordersSetData = append(ordersSetData, orderSetInfo)
	}
	rows.Close()
	return ordersSetData, nil
}

func (r *LomsRepository) getPossibleReservedItemsInfo(
	ctx context.Context,
	tx pgx.Tx,
	ordersSetData []orderSetData,
	orderID int64,
) ([]warehouseReservedItems, error) {
	warehouseReservedItemsInfo := make([]warehouseReservedItems, 0)

	for _, orderSetInfo := range ordersSetData {
		queryWarehouseItems := `SELECT available, warehouse_id
		FROM warehouses_items
		WHERE item_sku = $1 AND available > 0
		ORDER BY available DESC;`

		warehouseItemsRows, err := tx.Query(ctx, queryWarehouseItems, orderSetInfo.sku)
		if err != nil {
			log.Println("Query error on warehouses_items:", err.Error())
			return nil, err
		}
		for warehouseItemsRows.Next() {
			var warehouseInfo warehouseItemData
			err = warehouseItemsRows.Scan(&warehouseInfo.available, &warehouseInfo.warehouseID)
			if err != nil {
				log.Println("Query error on warehouses_items scan:", err.Error())
				return nil, err
			}

			if warehouseInfo.available-orderSetInfo.count >= 0 {
				warehouseReservedItemsInfo = append(
					warehouseReservedItemsInfo,
					warehouseReservedItems{
						warehouseID: warehouseInfo.warehouseID,
						count:       orderSetInfo.count,
						itemSku:     orderSetInfo.sku,
						orderSetID:  orderSetInfo.orderSetID,
					},
				)
				warehouseItemsRows.Close()
				orderSetInfo.count = 0
				break
			}
			warehouseReservedItemsInfo = append(warehouseReservedItemsInfo, warehouseReservedItems{
				warehouseID: warehouseInfo.warehouseID,
				count:       warehouseInfo.available,
				itemSku:     orderSetInfo.sku,
				orderSetID:  orderSetInfo.orderSetID,
			})
			orderSetInfo.count -= warehouseInfo.available
		}
		warehouseItemsRows.Close()

		if orderSetInfo.count > 0 {
			log.Println("Not enough items")
			err = r.updateStatus(ctx, tx, models.Failed, orderID)
			if err != nil {
				return nil, err
			}
			err = tx.Commit(ctx)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	}
	return warehouseReservedItemsInfo, nil
}

func (r *LomsRepository) updateWarehousesItems(
	ctx context.Context,
	tx pgx.Tx,
	wInfo warehouseReservedItems,
) error {
	queryWarehouseItems := `UPDATE warehouses_items
		SET available = available - $1, reserved = reserved + $1
		WHERE warehouse_id = $2 AND item_sku = $3`
	_, err := tx.Exec(ctx, queryWarehouseItems, wInfo.count, wInfo.warehouseID, wInfo.itemSku)
	if err != nil {
		log.Println("Query error on warehouses_items update:", err.Error())
		return err
	}
	return nil
}

func (r *LomsRepository) insertOrdersSetCountInWarehouse(
	ctx context.Context,
	tx pgx.Tx,
	wInfo warehouseReservedItems,
) error {
	queryOrderSetCount := `INSERT INTO orders_set_count_in_warehouse(item_count, warehouse_id, order_set_id)
		VALUES ($1, $2, $3)`
	_, err := tx.Exec(ctx, queryOrderSetCount, wInfo.count, wInfo.warehouseID, wInfo.orderSetID)
	if err != nil {
		log.Println("Query error on orders_set_count_in_warehouse:", err.Error())
		return err
	}
	return nil
}

func (r *LomsRepository) updateStatus(ctx context.Context, tx pgx.Tx, status models.OrderStatusID, orderID int64) error {
	queryOrder := `UPDATE orders
	SET status = $1
	WHERE order_id = $2;`

	_, err := tx.Exec(ctx, queryOrder, status, orderID)
	if err != nil {
		log.Println("Query error on orders update status:", err.Error())
		return err
	}
	return nil
}
