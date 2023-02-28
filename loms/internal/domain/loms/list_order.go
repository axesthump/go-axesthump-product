package loms

import (
	"context"
	"route256/loms/internal/models"
)

func (s *Service) ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error) {
	items := []models.Item{
		{Sku: 1, Count: 2},
		{Sku: 5, Count: 21},
	}
	return models.OrderInfo{
		Status: models.OrderStatus(1),
		User:   1,
		Items:  items,
	}, nil
}
