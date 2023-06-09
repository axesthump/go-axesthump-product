package checkout_v1

import (
	"context"
	"log"
	desc "route256/checkout/pkg/checkout_v1"
)

func (h *Handler) Purchase(ctx context.Context, r *desc.PurchaseRequest) (*desc.PurchaseResponse, error) {
	orderID, err := h.service.Purchase(ctx, r.GetUser())
	if err != nil {
		log.Println(err.Error())
		if orderID != 0 {
			return &desc.PurchaseResponse{
				OrderID: orderID,
			}, nil
		}
		return nil, err
	}
	return &desc.PurchaseResponse{
		OrderID: orderID,
	}, nil
}
