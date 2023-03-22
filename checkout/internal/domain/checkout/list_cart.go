package checkout

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

//skus := []uint32{
//1076963,
//1148162,
//1625903,
//2618151,
//2956315,
//2958025,
//3596599,
//3618852,
//4288068,
//4465995,
//}

func (s *Service) ListCart(ctx context.Context, user int64) (models.CartInfo, error) {
	items, err := s.repository.ListCart(ctx, user)
	if err != nil {
		return models.CartInfo{}, fmt.Errorf("get products: %w", err)
	}

	fillItems, err := s.productInfoGetter.GetProductsInfo(ctx, items)
	if err != nil {
		return models.CartInfo{}, fmt.Errorf("get products fill items: %w", err)
	}
	return s.getCartInfo(fillItems)
}

func (s *Service) getCartInfo(items []models.Item) (models.CartInfo, error) {
	cartInfo := models.CartInfo{Items: items}
	for _, item := range items {
		cartInfo.TotalPrice += item.Price * item.Count
	}
	return cartInfo, nil
}
