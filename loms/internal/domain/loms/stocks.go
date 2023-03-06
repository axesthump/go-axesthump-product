package loms

import (
	"context"
	"route256/loms/internal/models"
)

func (s *Service) Stocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	return []models.Stock{
		{
			WarehouseID: 1,
			Count:       2,
		},
		{
			WarehouseID: 2,
			Count:       5,
		},
	}, nil
}
