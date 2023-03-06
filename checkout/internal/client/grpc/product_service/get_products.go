package product_service

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/models"
	productServiceAPI "route256/checkout/pkg/product_service_v1"
)

func (c *Client) GetProducts(ctx context.Context, skus []uint32) ([]models.Product, error) {
	products := make([]models.Product, len(skus))
	for i, sku := range skus {
		var req = &productServiceAPI.GetProductRequest{
			Token: config.ConfigData.Token,
			Sku:   sku,
		}
		response, err := c.productServiceClient.GetProduct(ctx, req)
		if err != nil {
			return nil, err
		}
		products[i] = models.Product{
			Name:  response.GetName(),
			Price: response.GetPrice(),
		}
	}
	return products, nil
}
