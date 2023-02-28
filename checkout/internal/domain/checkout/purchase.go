package checkout

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (s *Service) Purchase(ctx context.Context, user int64) (int64, error) {
	items := []models.CreateOrderItem{
		{Sku: 1, Count: 10},
		{Sku: 2, Count: 20},
	}
	res, err := s.createOrderChecker.CreateOrder(ctx, user, items)
	if err != nil {
		return res, fmt.Errorf("purchase: %w", err)
	}
	return res, nil
}
