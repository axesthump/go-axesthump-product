package loms_service

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	lomsAPI "route256/loms/pkg/loms_v1"
)

func (c *Client) CreateOrder(
	ctx context.Context,
	user int64,
	items []models.CreateOrderItem,
) (int64, error) {

	descItems := toItemsDesc(items)
	request := &lomsAPI.CreateOrderRequest{
		User:  user,
		Items: descItems,
	}
	res, err := c.lomsClient.CreateOrder(ctx, request)
	if err != nil {
		return 0, fmt.Errorf("create order: %w", err)
	}
	return res.GetOrderID(), nil
}

func toItemsDesc(items []models.CreateOrderItem) []*lomsAPI.Item {
	descItems := make([]*lomsAPI.Item, len(items))
	for i, item := range items {
		descItems[i] = &lomsAPI.Item{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}
	return descItems
}
