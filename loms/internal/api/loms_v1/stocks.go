package loms_v1

import (
	"context"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/loms_v1"
)

func (h *Handler) Stocks(ctx context.Context, r *desc.StocksRequest) (*desc.StocksResponse, error) {
	stocks, err := h.service.Stocks(ctx, r.GetSku())
	if err != nil {
		return nil, err
	}
	res := ToStocksDesc(stocks)
	return res, nil
}

func ToStocksDesc(stocks []models.Stock) *desc.StocksResponse {
	descStocks := make([]*desc.Stock, len(stocks))
	for i, stock := range stocks {
		descStocks[i] = &desc.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		}
	}
	return &desc.StocksResponse{
		Stocks: descStocks,
	}
}
