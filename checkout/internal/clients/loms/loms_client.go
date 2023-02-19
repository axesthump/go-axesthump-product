package loms

import (
	"context"
	"route256/checkout/internal/models"
	"route256/libs/clientwrapper"

	"github.com/pkg/errors"
)

type Client struct {
	url            string
	urlStocks      string
	urlCreateOrder string
}

func New(url string) *Client {
	return &Client{
		url:            url,
		urlCreateOrder: url + "/createOrder",
		urlStocks:      url + "/stocks",
	}
}

type (
	StocksRequest struct {
		SKU uint32 `json:"sku"`
	}
	StocksResponse struct {
		Stocks []models.Stock `json:"stocks"`
	}
)

type (
	CreateOrderRequest struct {
		User  int64                    `json:"user"`
		Items []models.CreateOrderItem `json:"items"`
	}
	CreateOrderResponse struct {
		OrderID int64 `json:"orderID"`
	}
)

func (c *Client) GetStocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	request := StocksRequest{SKU: sku}
	response, err := clientwrapper.SendRequest[StocksRequest, StocksResponse](ctx, request, c.urlStocks)
	if err != nil {
		return nil, errors.WithMessage(err, "get stocks")
	}
	return response.Stocks, nil
}

func (c *Client) CreateOrder(ctx context.Context, user int64, items []models.CreateOrderItem) error {
	request := CreateOrderRequest{User: user, Items: items}
	_, err := clientwrapper.SendRequest[CreateOrderRequest, CreateOrderResponse](ctx, request, c.urlCreateOrder)
	if err != nil {
		return errors.WithMessage(err, "create order")
	}
	return nil
}
