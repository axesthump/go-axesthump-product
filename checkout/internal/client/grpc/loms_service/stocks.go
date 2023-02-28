package loms_service

import (
	"context"
	"route256/checkout/internal/models"
	lomsAPI "route256/loms/pkg/loms_v1"
)

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]models.Stock, error) {
	req := &lomsAPI.StocksRequest{Sku: sku}
	stocks, err := c.lomsClient.Stocks(ctx, req)
	if err != nil {
		return nil, err
	}
	res := toStocks(stocks)
	return res, nil
}

func toStocks(stocks *lomsAPI.StocksResponse) []models.Stock {
	convertedStocks := make([]models.Stock, len(stocks.GetStocks()))
	for i, stock := range stocks.GetStocks() {
		convertedStocks[i] = models.Stock{
			WarehouseID: stock.GetWarehouseID(),
			Count:       stock.GetCount(),
		}
	}
	return convertedStocks
}
