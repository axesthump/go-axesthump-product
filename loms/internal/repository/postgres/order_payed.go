package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"

	"github.com/jackc/pgx/v4"
)

func (r *LomsRepository) OrderPayed(ctx context.Context, orderID int64) error {
	err := r.inTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		reservedItems, err := r.getReservedItemsInWarehousesFromOrder(ctx, tx, orderID)
		if err != nil {
			return err
		}
		for _, reservedItem := range reservedItems {
			err = r.updateBoughtItems(ctx, tx, reservedItem)
			if err != nil {
				return err
			}
		}
		err = r.updateStatus(ctx, tx, models.Payed, orderID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("order payed: %w", err)
	}
	return nil
}

func (r *LomsRepository) updateBoughtItems(
	ctx context.Context,
	tx pgx.Tx,
	reservedItem reservedItemsInWarehouses,
) error {
	query := `UPDATE warehouses_items 
	SET reserved = reserved - $1, bought = bought + $1
	WHERE warehouse_id = $2 AND sku = $3;`

	_, err := tx.Exec(ctx, query, reservedItem.itemCount, reservedItem.warehouseID, reservedItem.itemSku)
	if err != nil {
		return fmt.Errorf("postgres updateBoughtItems: %w", err)
	}
	return err
}
