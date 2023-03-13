package checkout

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (s *Service) Purchase(ctx context.Context, user int64) (int64, error) {
	items, err := s.repository.ListCart(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("purchase: %w", err)
	}
	itemsOrder := make([]models.CreateOrderItem, len(items))
	for i := range items {
		itemsOrder[i] = models.CreateOrderItem{
			Sku:   items[i].Sku,
			Count: items[i].Count,
		}
	}
	res, err := s.createOrderChecker.CreateOrder(ctx, user, itemsOrder)
	if err != nil {
		return res, fmt.Errorf("purchase: %w", err)
	}
	err = s.repository.ClearCart(ctx, user)
	if err != nil {
		return res, fmt.Errorf("purchase: %w", err)
	}

	return res, nil
}
