package loms

import (
	"context"
	"route256/loms/internal/models"
)

func (s *Service) ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error) {
	return s.repository.ListOrder(ctx, orderID)
}
