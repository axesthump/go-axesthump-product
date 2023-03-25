package productinfogetter

import (
	"context"
	"route256/checkout/internal/models"
	"route256/checkout/internal/pool"
)

type ProductsChecker interface {
	GetProduct(ctx context.Context, sku uint32) (models.Product, error)
}

type ProductInfoGetter struct {
	productChecker ProductsChecker
}

func New(checker ProductsChecker) *ProductInfoGetter {
	return &ProductInfoGetter{
		productChecker: checker,
	}
}

func (p *ProductInfoGetter) GetProductsInfo(ctx context.Context, items []models.Item) ([]models.Item, error) {
	workerPool := pool.New(ctx, p.productChecker, 5)
	workerPool.Submit(items)
	fillItems := make([]models.Item, 0, len(items))

	for i := 0; i < len(items); i++ {
		select {
		case item := <-workerPool.Out:
			fillItems = append(fillItems, item)
		case err := <-workerPool.Err:
			return nil, err
		}
	}
	return fillItems, nil
}
