package checkout

import (
	"context"
	"route256/checkout/internal/models"
)

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]models.Stock, error)
}

type ProductsChecker interface {
	GetProduct(ctx context.Context, sku uint32) (models.Product, error)
}

type CreateOrderChecker interface {
	CreateOrder(ctx context.Context, user int64, items []models.CreateOrderItem) (int64, error)
}

type Repository interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error
	ListCart(ctx context.Context, user int64) ([]models.Item, error)
	ClearCart(ctx context.Context, user int64) error
}

type Service struct {
	stocksChecker      StocksChecker
	productChecker     ProductsChecker
	createOrderChecker CreateOrderChecker
	repository         Repository
}

func New(
	stocksChecker StocksChecker,
	productsChecker ProductsChecker,
	createOrderChecker CreateOrderChecker,
	repository Repository,
) *Service {
	return &Service{
		stocksChecker:      stocksChecker,
		productChecker:     productsChecker,
		createOrderChecker: createOrderChecker,
		repository:         repository,
	}

}
