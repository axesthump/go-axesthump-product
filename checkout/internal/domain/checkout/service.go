package checkout

//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i ProductInfoGetter -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i CreateOrderChecker -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i StocksChecker -o ./mocks/ -s "_minimock.go"

import (
	"context"
	"route256/checkout/internal/models"
)

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]models.Stock, error)
}

type CreateOrderChecker interface {
	CreateOrder(ctx context.Context, user int64, items []models.CreateOrderItem) (int64, error)
}

type ProductInfoGetter interface {
	GetProductsInfo(ctx context.Context, items []models.Item) ([]models.Item, error)
}

type Repository interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error
	ListCart(ctx context.Context, user int64) ([]models.Item, error)
	ClearCart(ctx context.Context, user int64) error
}

type Service struct {
	stocksChecker      StocksChecker
	createOrderChecker CreateOrderChecker
	repository         Repository
	productInfoGetter  ProductInfoGetter
}

func New(
	stocksChecker StocksChecker,
	createOrderChecker CreateOrderChecker,
	productInfoGetter ProductInfoGetter,
	repository Repository,
) *Service {
	return &Service{
		stocksChecker:      stocksChecker,
		createOrderChecker: createOrderChecker,
		productInfoGetter:  productInfoGetter,
		repository:         repository,
	}

}
