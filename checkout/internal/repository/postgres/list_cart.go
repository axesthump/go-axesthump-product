package postgres

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (r *CheckoutRepository) ListCart(ctx context.Context, user int64) ([]models.Item, error) {
	const query = `
	SELECT sku, count 
	FROM cart_items ci
	INNER JOIN carts c 
	ON c.user_id = $1 AND c.id = ci.cart_id AND ci.count > 0;
	`

	rows, err := r.pool.Query(ctx, query, user)
	if err != nil {
		return nil, fmt.Errorf("postgres ListCart select: %w", err)
	}
	defer rows.Close()

	items := make([]models.Item, 0, 5)
	for rows.Next() {
		var sku uint32
		var count uint32

		err = rows.Scan(&sku, &count)
		if err != nil {
			return nil, fmt.Errorf("postgres ListCart scan: %w", err)
		}
		item := models.Item{
			Sku:   sku,
			Count: count,
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
