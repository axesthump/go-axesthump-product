package product_service

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/models"
	productServiceAPI "route256/checkout/pkg/product_service_v1"
)

func (c *Client) GetProduct(ctx context.Context, sku uint32) (models.Product, error) {
	var req = &productServiceAPI.GetProductRequest{
		Token: config.ConfigData.Token,
		Sku:   sku,
	}
	response, err := c.productServiceClient.GetProduct(ctx, req)
	if err != nil {
		return models.Product{}, err
	}
	product := models.Product{
		Name:  response.GetName(),
		Price: response.GetPrice(),
	}
	return product, nil
}
