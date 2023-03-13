package loms

import (
	"context"
	"route256/loms/internal/models"
)

func (s *Service) Stocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	return s.repository.Stocks(ctx, sku)
}
