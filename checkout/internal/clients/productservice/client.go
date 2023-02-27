package productservice

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/models"
	"route256/libs/clientwrapper"
)

type Client struct {
}

const getProductURL = "/get_product"

func New() *Client {
	return &Client{}
}

type (
	GetProductRequest struct {
		Token string `json:"token"`
		Sku   uint32 `json:"sku"`
	}

	GetProductResponse struct {
		models.Product
	}
)

func (c *Client) GetProducts(ctx context.Context, skus []uint32) ([]models.Product, error) {
	products := make([]models.Product, len(skus))
	for i, sku := range skus {
		request := GetProductRequest{Token: config.ConfigData.Token, Sku: sku}
		response, err := clientwrapper.SendRequest[GetProductRequest, GetProductResponse](
			ctx,
			request,
			config.ConfigData.Services.ProductService+getProductURL,
		)
		if err != nil {
			return nil, err
		}
		products[i] = response.Product
	}
	return products, nil
}
