package checkout

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (s *Service) ListCart(ctx context.Context, user int64) (models.CartInfo, error) {
	skus := []uint32{
		1076963,
		1148162,
		1625903,
		2618151,
		2956315,
		2958025,
		3596599,
		3618852,
		4288068,
		4465995,
	}

	products, err := s.productsChecker.GetProducts(ctx, skus)
	if err != nil {
		return models.CartInfo{}, fmt.Errorf("get products: %w", err)
	}
	items := make([]models.Item, 0, len(products))
	cartInfo := models.CartInfo{Items: items}
	for _, product := range products {
		cartInfo.TotalPrice += product.Price
		cartInfo.Items = append(
			cartInfo.Items,
			models.Item{
				Sku:   1,
				Count: 2,
				Name:  product.Name,
				Price: product.Price,
			},
		)
	}
	return cartInfo, nil

}
