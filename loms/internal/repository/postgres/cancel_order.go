package postgres

import (
	"context"
	"route256/loms/internal/models"

	"github.com/jackc/pgx/v4"
)

func (r *LomsRepository) CancelOrder(ctx context.Context, orderID int64) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	reservedItems, err := r.getReservedItemsInWarehousesFromOrder(ctx, tx, orderID)
	if err != nil {
		return err
	}
	for _, reservedItem := range reservedItems {
		err = r.updateReservedItems(ctx, tx, reservedItem)
		if err != nil {
			return err
		}
	}
	err = r.updateStatus(ctx, tx, models.Cancelled, orderID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *LomsRepository) updateReservedItems(
	ctx context.Context,
	tx pgx.Tx,
	reservedItem reservedItemsInWarehouses,
) error {
	query := `UPDATE warehouses_items 
	SET reserved = reserved - $1, available = available + $1
	WHERE warehouse_items_id = $2;`

	_, err := tx.Exec(ctx, query, reservedItem.itemCount, reservedItem.warehouseItemID)
	return err
}
