package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"
)

func (h *Handler) OrderPayed(ctx context.Context, r *desc.OrderPayedRequest) (*desc.EmptyMessage, error) {
	err := h.service.OrderPayed(ctx, r.GetOrderID())
	if err != nil {
		return nil, err
	}
	return &desc.EmptyMessage{}, nil
}
