package loms_v1

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/loms_v1"
)

func (h *Handler) CreateOrder(ctx context.Context, r *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	orderInfo, err := toOrderData(r)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	orderID, err := h.service.CreateOrder(ctx, orderInfo)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &desc.CreateOrderResponse{OrderID: orderID}, nil
}

func toOrderData(r *desc.CreateOrderRequest) (models.OrderData, error) {
	descItems := r.GetItems()
	if descItems == nil {
		return models.OrderData{}, fmt.Errorf("no items")
	}
	items := make([]models.Item, len(descItems))
	for i, item := range descItems {
		items[i] = models.Item{Sku: item.GetSku(), Count: item.GetCount()}
	}
	return models.OrderData{
		User:  r.GetUser(),
		Items: items,
	}, nil
}
