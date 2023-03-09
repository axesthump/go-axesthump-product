package loms

import (
	"context"
	"errors"
	"route256/loms/internal/models"
)

func (s *Service) CancelOrder(ctx context.Context, orderID int64) error {
	status, err := s.repository.GetStatus(ctx, orderID)
	if err != nil {
		return err
	}
	if status == models.New || status == models.AwaitingPayment {
		return s.repository.CancelOrder(ctx, orderID)
	} else {
		return errors.New("wrong status for cancel")
	}

}
