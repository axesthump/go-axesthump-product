package productservice

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/models"
	"route256/libs/clientwrapper"
)

type Client struct {
	url           string
	urlGetProduct string
}

func New(url string) *Client {
	return &Client{
		url:           url,
		urlGetProduct: url + "/get_product",
	}
}

type (
	GetProductRequest struct {
		Token string `json:"token"`
		SKU   uint32 `json:"sku"`
	}

	GetProductResponse struct {
		models.Product
	}
)

func (c *Client) GetProducts(ctx context.Context, skus []uint32) ([]models.Product, error) {
	products := make([]models.Product, len(skus))
	for i, sku := range skus {
		request := GetProductRequest{Token: config.ConfigData.Token, SKU: sku}
		response, err := clientwrapper.SendRequest[GetProductRequest, GetProductResponse](ctx, request, c.urlGetProduct)
		if err != nil {
			return nil, err
		}
		products[i] = response.Product
	}
	return products, nil
}
