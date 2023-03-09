package loms

import (
	"context"
	"errors"
	"route256/loms/internal/models"
)

func (s *Service) OrderPayed(ctx context.Context, orderID int64) error {
	status, err := s.repository.GetStatus(ctx, orderID)
	if err != nil {
		return err
	}
	if status != models.AwaitingPayment {
		return errors.New("wrong status for pay")
	}
	return s.repository.OrderPayed(ctx, orderID)
}
