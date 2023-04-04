package postgres

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
)

func (r *LomsRepository) SaveFailedSendsOrders(ctx context.Context, orders []models.ErrOrder) error {
	const query = `
	UPDATE outbox_orders
	SET err_message = msgs.msg
	FROM (
	    SELECT * FROM 
	    unnest($1::text[]) as msg,
	    unnest($2::bigint[]) as s_id
	) msgs
	WHERE id = msgs.s_id;
	`

	errorsMessages := make([]string, len(orders))
	ids := make([]int64, len(orders))

	for i, order := range orders {
		errorsMessages[i] = order.Err.Error()
		ids[i] = order.Order.ID
	}

	_, err := r.pool.Exec(ctx, query, errorsMessages, ids)
	if err != nil {
		return fmt.Errorf("SaveFailedSendsOrders: %w", err)
	}
	return nil

}
