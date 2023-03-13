package loms

import (
	"context"
	"errors"
	"fmt"
	"route256/loms/internal/models"
)

func (s *Service) CancelOrder(ctx context.Context, orderID int64) error {
	status, err := s.repository.GetStatus(ctx, orderID)
	if err != nil {
		return fmt.Errorf("service CancelOrder: %w", err)
	}
	if status == models.New || status == models.AwaitingPayment {
		return s.repository.CancelOrder(ctx, orderID)
	}
	return errors.New("wrong status for cancel")
}
