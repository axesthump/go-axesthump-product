package loms

import (
	"context"
	"route256/loms/internal/models"
)

type Repository interface {
	CreateOrder(ctx context.Context, order models.OrderData) (int64, error)
	Stocks(ctx context.Context, sku uint32) ([]models.Stock, error)
	ReservedItems(ctx context.Context, orderID int64)
	ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error)
	OrderPayed(ctx context.Context, orderID int64) error
	CancelOrder(ctx context.Context, orderID int64) error
	GetStatus(ctx context.Context, orderID int64) (models.OrderStatusID, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{repository: repository}
}
