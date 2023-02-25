package stocks

import (
	"context"
	"log"
	"route256/loms/internal/handlers"
	"route256/loms/internal/models"
)

type getStocks interface {
	GetStocks(ctx context.Context, orderID int64) ([]models.Stock, error)
}

type (
	Request struct {
		Sku int64 `json:"sku"`
	}

	Response struct {
		Stocks []models.Stock `json:"stocks"`
	}
)

func (r Request) Validate() error {
	if r.Sku == 0 {
		return handlers.ErrEmptyOrderID
	}
	return nil
}

type Handler struct {
	service getStocks
}

func New(service getStocks) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("stocks: %+v", req)
	stocks, err := h.service.GetStocks(ctx, req.Sku)
	var response Response
	if err != nil {
		return response, err
	}
	response.Stocks = stocks
	return response, err
}
