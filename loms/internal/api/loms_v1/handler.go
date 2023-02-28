package loms_v1

import (
	"context"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/loms_v1"
)

type Service interface {
	CreateOrder(ctx context.Context, order models.OrderData) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error)
	OrderPayed(ctx context.Context, orderID int64) error
	CancelOrder(ctx context.Context, orderID int64) error
	Stocks(ctx context.Context, sku uint32) ([]models.Stock, error)
}

type Handler struct {
	desc.UnimplementedLomsV1Server
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}
