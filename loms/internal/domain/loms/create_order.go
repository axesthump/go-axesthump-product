package loms

import (
	"context"
	"route256/loms/internal/models"
)

func (s *Service) CreateOrder(ctx context.Context, order models.OrderData) (int64, error) {
	return 42, nil
}
