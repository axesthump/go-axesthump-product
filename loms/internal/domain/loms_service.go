package domain

import (
	"context"
	"route256/loms/internal/models"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error) {
	return 42, nil
}

func (s *Service) ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error) {
	items := []models.Item{
		{Sku: 1, Count: 2},
		{Sku: 5, Count: 21},
	}
	return models.OrderInfo{
		Status: "new",
		User:   1,
		Items:  items,
	}, nil
}

func (s *Service) OrderPayed(ctx context.Context, orderID int64) error {
	return nil
}

func (s *Service) CancelOrder(ctx context.Context, orderID int64) error {
	return nil
}

func (s *Service) GetStocks(ctx context.Context, orderID int64) ([]models.Stock, error) {
	return []models.Stock{
		{WarehouseID: 12, Count: 100},
		{WarehouseID: 123, Count: 4},
		{WarehouseID: 112, Count: 1020},
	}, nil
}
