package loms_v1

import (
	"context"
	"route256/loms/internal/models"
	desc "route256/loms/pkg/loms_v1"
)

func (h *Handler) ListOrder(ctx context.Context, r *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	orderInfo, err := h.service.ListOrder(ctx, r.GetOrderID())
	if err != nil {
		return nil, err
	}
	res := toOrderInfoDesc(orderInfo)
	return res, nil
}

func toOrderInfoDesc(info models.OrderInfo) *desc.ListOrderResponse {
	status := getStatusDesc(info.Status)
	items := toItemDesc(info.Items)
	return &desc.ListOrderResponse{
		Status: status,
		User:   info.User,
		Items:  items,
	}
}

func toItemDesc(items []models.Item) []*desc.Item {
	descItems := make([]*desc.Item, len(items))

	for i, item := range items {
		descItems[i] = &desc.Item{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	return descItems
}

func getStatusDesc(status models.OrderStatus) desc.OrderStatus {
	switch status {
	case models.New:
		return desc.OrderStatus_new
	case models.AwaitingPayment:
		return desc.OrderStatus_awaiting_payment
	case models.Failed:
		return desc.OrderStatus_failed
	case models.Payed:
		return desc.OrderStatus_payed
	case models.Cancelled:
		return desc.OrderStatus_cancelled
	default:
		return desc.OrderStatus_nil
	}
}
