package checkout_v1

import (
	"context"
	"fmt"
	desc "route256/checkout/pkg/checkout_v1"
)

func (h *Handler) AddToCart(ctx context.Context, r *desc.AddToCartRequest) (*desc.EmptyMessage, error) {
	productInfo := r.GetProductInfo()
	if productInfo == nil {
		return nil, fmt.Errorf("no info about order")
	}
	err := h.service.AddToCart(ctx, productInfo.GetUser(), productInfo.GetSku(), productInfo.GetCount())
	if err != nil {
		return nil, err
	}
	return &desc.EmptyMessage{}, nil
}
