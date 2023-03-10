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
	skus := make([]uint32, len(items))
	for i := range items {
		skus[i] = items[i].Sku
	}

	products, err := s.productsChecker.GetProducts(ctx, skus)
	if err != nil {
		return models.CartInfo{}, fmt.Errorf("get products: %w", err)
	}
	return s.getCartInfo(items, products)

}

func (s *Service) getCartInfo(items []models.Item, products []models.Product) (models.CartInfo, error) {
	cartInfo := models.CartInfo{Items: items}
	for i, product := range products {
		cartInfo.TotalPrice += product.Price * items[i].Count
		cartInfo.Items[i].Name = product.Name
		cartInfo.Items[i].Price = product.Price
	}
	return cartInfo, nil
}
