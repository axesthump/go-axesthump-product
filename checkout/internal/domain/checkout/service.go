package checkout

import (
	"context"
	"route256/checkout/internal/models"
)

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]models.Stock, error)
}

type ProductsChecker interface {
	GetProducts(ctx context.Context, skus []uint32) ([]models.Product, error)
}

type CreateOrderChecker interface {
	CreateOrder(ctx context.Context, user int64, items []models.CreateOrderItem) (int64, error)
}

type Service struct {
	stocksChecker      StocksChecker
	productsChecker    ProductsChecker
	createOrderChecker CreateOrderChecker
}

func New(stocksChecker StocksChecker, productsChecker ProductsChecker, createOrderChecker CreateOrderChecker) *Service {
	return &Service{
		stocksChecker:      stocksChecker,
		productsChecker:    productsChecker,
		createOrderChecker: createOrderChecker,
	}

}
