package checkout_v1

import (
	"context"
	"fmt"
	desc "route256/checkout/pkg/checkout_v1"
)

func (h *Handler) DeleteFromCart(ctx context.Context, r *desc.DeleteFromCartRequest) (*desc.EmptyMessage, error) {
	productInfo := r.GetProductInfo()
	if productInfo == nil {
		return nil, fmt.Errorf("no info about order")
	}
	err := h.service.DeleteFromCart(ctx, productInfo.GetUser(), productInfo.GetSku(), productInfo.GetCount())
	if err != nil {
		return nil, err
	}
	return &desc.EmptyMessage{}, nil
}
