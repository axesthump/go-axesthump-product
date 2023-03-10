package loms_v1

import (
	"context"
	"log"
	desc "route256/loms/pkg/loms_v1"
)

func (h *Handler) OrderPayed(ctx context.Context, r *desc.OrderPayedRequest) (*desc.EmptyMessage, error) {
	err := h.service.OrderPayed(ctx, r.GetOrderID())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &desc.EmptyMessage{}, nil
}
