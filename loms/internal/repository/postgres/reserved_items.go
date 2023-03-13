package postgres

import (
	"context"
	"fmt"
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

func (r *LomsRepository) ReservedItems(ctx context.Context, orderID int64) error {
	err := r.inTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		ordersSetData, err := r.getOrdersSetData(ctx, tx, orderID)
		if err != nil {
			return err
		}

		possibleReservedItemsInfo, err := r.getPossibleReservedItemsInfo(ctx, tx, ordersSetData, orderID)
		if err != nil {
			return err
		}
		if possibleReservedItemsInfo == nil {
			return err
		}

		for _, wInfo := range possibleReservedItemsInfo {
			err = r.updateWarehousesItems(ctx, tx, wInfo)
			if err != nil {
				return err
			}

			err = r.insertOrdersSetCountInWarehouse(ctx, tx, wInfo)
			if err != nil {
				return err
			}
		}

		err = r.updateStatus(ctx, tx, models.AwaitingPayment, orderID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("reserved items: %w", err)
	}
	return nil
}

func (r *LomsRepository) getOrdersSetData(ctx context.Context, tx pgx.Tx, orderID int64) ([]orderSetData, error) {
	const query = `
	SELECT sku, count, id
	FROM order_items
	WHERE order_id = $1;
	`

	rows, err := tx.Query(ctx, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("postgres getOrdersSetData select: %w", err)
	}
	defer rows.Close()
	ordersSetData := make([]orderSetData, 0)
	for rows.Next() {
		var orderSetInfo orderSetData

		err = rows.Scan(&orderSetInfo.sku, &orderSetInfo.count, &orderSetInfo.orderSetID)
		if err != nil {
			return nil, fmt.Errorf("postgres getOrdersSetData scan: %w", err)
		}
		ordersSetData = append(ordersSetData, orderSetInfo)
	}
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
		const query = `
		SELECT available, warehouse_id
		FROM warehouses_items
		WHERE sku = $1 AND available > 0
		ORDER BY available DESC;
		`

		warehouseItemsRows, err := tx.Query(ctx, query, orderSetInfo.sku)
		if err != nil {
			return nil, fmt.Errorf("postgres getPossibleReservedItemsInfo select: %w", err)
		}
		for warehouseItemsRows.Next() {
			var warehouseInfo warehouseItemData
			err = warehouseItemsRows.Scan(&warehouseInfo.available, &warehouseInfo.warehouseID)
			if err != nil {
				return nil, fmt.Errorf("postgres getPossibleReservedItemsInfo scan: %w", err)
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
			err = r.updateStatus(ctx, tx, models.Failed, orderID)
			if err != nil {
				return nil, fmt.Errorf("postgres getPossibleReservedItemsInfo: %w", err)
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
	const query = `
	UPDATE warehouses_items
	SET available = available - $1, reserved = reserved + $1
	WHERE warehouse_id = $2 AND sku = $3;
	`

	_, err := tx.Exec(ctx, query, wInfo.count, wInfo.warehouseID, wInfo.itemSku)
	if err != nil {
		return fmt.Errorf("postgres updateWarehousesItems: %w", err)
	}
	return nil
}

func (r *LomsRepository) insertOrdersSetCountInWarehouse(
	ctx context.Context,
	tx pgx.Tx,
	wInfo warehouseReservedItems,
) error {
	const query = `
	INSERT INTO order_items_count_in_warehouse(count, warehouse_id, order_items_id)
	VALUES ($1, $2, $3);
	`

	_, err := tx.Exec(ctx, query, wInfo.count, wInfo.warehouseID, wInfo.orderSetID)
	if err != nil {
		return fmt.Errorf("postgres insertOrdersSetCountInWarehouse: %w", err)
	}
	return nil
}

func (r *LomsRepository) updateStatus(ctx context.Context, tx pgx.Tx, status models.OrderStatusID, orderID int64) error {
	const query = `
	UPDATE orders
	SET status = $1
	WHERE id = $2;
	`

	_, err := tx.Exec(ctx, query, status, orderID)
	if err != nil {
		return fmt.Errorf("postgres updateStatus: %w", err)
	}
	return nil
}
