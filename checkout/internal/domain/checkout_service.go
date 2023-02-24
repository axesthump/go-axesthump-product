package domain

import (
	"context"
	"errors"
	"fmt"
	"route256/checkout/internal/models"
)

type StocksChecker interface {
	GetStocks(ctx context.Context, sku uint32) ([]models.Stock, error)
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

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

func New(
	stocksChecker StocksChecker,
	productsChecker ProductsChecker,
	createOrderChecker CreateOrderChecker,
) *Service {
	return &Service{
		stocksChecker:      stocksChecker,
		productsChecker:    productsChecker,
		createOrderChecker: createOrderChecker,
	}
}

func (s *Service) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	stocks, err := s.stocksChecker.GetStocks(ctx, sku)
	if err != nil {
		return fmt.Errorf("checking stocks: %w", err)
	}

	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			return nil
		}
	}

	return ErrInsufficientStocks
}

func (s *Service) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	return nil
}

func (s *Service) ListCart(ctx context.Context, user int64) (*models.CartInfo, error) {
	skus := []uint32{
		1076963,
		1148162,
		1625903,
		2618151,
		2956315,
		2958025,
		3596599,
		3618852,
		4288068,
		4465995,
	}

	products, err := s.productsChecker.GetProducts(ctx, skus)
	if err != nil {
		return nil, fmt.Errorf("get products: %w", err)
	}
	items := make([]models.Item, 0, len(products))
	cartInfo := models.CartInfo{Items: items}
	for _, product := range products {
		cartInfo.TotalPrice += product.Price
		cartInfo.Items = append(
			cartInfo.Items,
			models.Item{
				SKU:   1,
				Count: 2,
				Name:  product.Name,
				Price: product.Price,
			},
		)
	}
	return &cartInfo, nil
}

func (s *Service) Purchase(ctx context.Context, user int64) (int64, error) {
	items := []models.CreateOrderItem{
		{Sku: 1, Count: 10},
		{Sku: 2, Count: 20},
	}
	res, err := s.createOrderChecker.CreateOrder(ctx, user, items)
	if err != nil {
		return res, fmt.Errorf("purchase: %w", err)
	}
	return res, nil
}
