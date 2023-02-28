package checkout_v1

import (
	"context"
	desc "route256/checkout/pkg/checkout_v1"
)

func (h *Handler) Purchase(ctx context.Context, r *desc.PurchaseRequest) (*desc.PurchaseResponse, error) {
	orderID, err := h.service.Purchase(ctx, r.GetUser())
	if err != nil {
		return nil, err
	}
	return &desc.PurchaseResponse{
		OrderID: orderID,
	}, nil
}
