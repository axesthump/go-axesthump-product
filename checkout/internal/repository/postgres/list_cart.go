package postgres

import (
	"context"
	"route256/checkout/internal/models"
)

func (r *CheckoutRepository) ListCart(ctx context.Context, user int64) ([]models.Item, error) {
	query := `SELECT item_sku, count 
	FROM carts_items ci
	INNER JOIN carts c 
	ON c.user_id = $1 AND c.cart_id = ci.cart_id AND ci.count > 0;`

	rows, err := r.pool.Query(ctx, query, user)
	if err != nil {
		return nil, err
	}

	items := make([]models.Item, 0, 5)
	for rows.Next() {
		var sku uint32
		var count uint32

		err = rows.Scan(&sku, &count)
		if err != nil {
			return nil, err
		}
		item := models.Item{
			Sku:   sku,
			Count: count,
		}
		items = append(items, item)
	}
	rows.Close()
	return items, err
}
