package cancelorder

import (
	"context"
	"log"
	"route256/loms/internal/handlers"
)

type cancelOrder interface {
	CancelOrder(ctx context.Context, orderID int64) error
}

type (
	Request struct {
		OrderID int64 `json:"orderID"`
	}

	Response struct{}
)

func (r Request) Validate() error {
	if r.OrderID == 0 {
		return handlers.ErrEmptyOrderID
	}
	return nil
}

type Handler struct {
	service cancelOrder
}

func New(service cancelOrder) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("cancelOrder: %+v", req)
	err := h.service.CancelOrder(ctx, req.OrderID)
	return Response{}, err
}
