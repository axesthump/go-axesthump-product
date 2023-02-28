package checkout_v1

import (
	"context"
	"route256/checkout/internal/models"
	desc "route256/checkout/pkg/checkout_v1"
)

func (h *Handler) ListCart(ctx context.Context, r *desc.ListCartRequest) (*desc.ListCartResponse, error) {
	cartInfo, err := h.service.ListCart(ctx, r.GetUser())
	if err != nil {
		return nil, err
	}

	items := toItemsDesc(cartInfo.Items)
	return &desc.ListCartResponse{
		Items:      items,
		TotalPrice: cartInfo.TotalPrice,
	}, nil
}

func toItemsDesc(items []models.Item) []*desc.CartItem {
	descItems := make([]*desc.CartItem, len(items))
	for i, item := range items {
		descItems[i] = &desc.CartItem{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  item.Name,
			Price: item.Price,
		}
	}
	return descItems
}
