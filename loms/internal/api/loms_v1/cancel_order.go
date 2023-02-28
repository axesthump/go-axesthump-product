package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"
)

func (h *Handler) CancelOrder(ctx context.Context, r *desc.CancelOrderRequest) (*desc.EmptyMessage, error) {
	err := h.service.CancelOrder(ctx, r.GetOrderID())
	if err != nil {
		return nil, err
	}
	return &desc.EmptyMessage{}, nil
}
