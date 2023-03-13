package loms

import (
	"context"
	"route256/loms/internal/models"
)

func (s *Service) CreateOrder(ctx context.Context, order models.OrderData) (int64, error) {
	orderID, err := s.repository.CreateOrder(ctx, order)
	if err != nil {
		return 0, err
	}

	err = s.repository.ReservedItems(ctx, orderID)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}
