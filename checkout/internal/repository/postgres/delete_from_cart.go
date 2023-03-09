package postgres

import "context"

func (r *CheckoutRepository) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	query := `UPDATE carts_items ci
	SET count = case count-$1 >= 0 WHEN TRUE THEN count-$1 ELSE 0 END
	FROM carts c
	WHERE c.cart_id = ci.cart_id AND c.user_id = $2 AND item_sku = $3;`

	_, err := r.pool.Exec(ctx, query, count, user, sku)
	return err
}
