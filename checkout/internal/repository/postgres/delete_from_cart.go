package postgres

import (
	"context"
	"fmt"
)

func (r *CheckoutRepository) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	query := `UPDATE cart_items ci
	SET count = case count-$1 >= 0 WHEN TRUE THEN count-$1 ELSE 0 END
	FROM carts c
	WHERE c.id = ci.cart_id AND c.user_id = $2 AND sku = $3;`

	_, err := r.pool.Exec(ctx, query, count, user, sku)
	if err != nil {
		return fmt.Errorf("postgres DeleteFromCart: %w", err)
	}
	return nil
}
