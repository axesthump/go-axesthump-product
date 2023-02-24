package loms

import (
	"context"
	"fmt"
	"route256/checkout/internal/config"
	"route256/checkout/internal/models"
	"route256/libs/clientwrapper"
)

type Client struct {
}

const (
	stocksURL      = "/stocks"
	createOrderURL = "/createOrder"
)

func New() *Client {
	return &Client{}
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
	response, err := clientwrapper.SendRequest[StocksRequest, StocksResponse](
		ctx,
		request,
		config.ConfigData.Services.Loms+stocksURL,
	)
	if err != nil {
		return nil, fmt.Errorf("get stocks: %w", err)
	}
	return response.Stocks, nil
}

func (c *Client) CreateOrder(
	ctx context.Context,
	user int64,
	items []models.CreateOrderItem,
) (int64, error) {
	request := CreateOrderRequest{User: user, Items: items}
	res, err := clientwrapper.SendRequest[CreateOrderRequest, CreateOrderResponse](
		ctx,
		request,
		config.ConfigData.Services.Loms+createOrderURL,
	)
	if err != nil {
		return -1, fmt.Errorf("create order: %w", err)
	}
	return res.OrderID, nil
}
