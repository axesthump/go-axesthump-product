package checkout

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrInsufficientStocks = errors.New("insufficient stocksChecker")
)

func (s *Service) AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	stocks, err := s.stocksChecker.Stocks(ctx, sku)
	if err != nil {
		return fmt.Errorf("checking stocksChecker: %w", err)
	}
	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			err = s.repository.AddToCart(ctx, user, sku, count)
			if err != nil {
				return fmt.Errorf("service AddToCart: %w", err)
			}
			return nil
		}
	}
	return ErrInsufficientStocks
}
