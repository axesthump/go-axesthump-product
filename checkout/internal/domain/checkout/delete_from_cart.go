package checkout

import (
	"context"
	"fmt"
)

func (s *Service) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	err := s.repository.DeleteFromCart(ctx, user, sku, count)
	if err != nil {
		return fmt.Errorf("service DeleteFromCart: %w", err)
	}
	return nil
}
