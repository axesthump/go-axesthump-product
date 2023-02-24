package orderpayed

import (
	"context"
	"log"
	"route256/loms/internal/handlers"
)

type orderPayed interface {
	OrderPayed(ctx context.Context, orderID int64) error
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
	service orderPayed
}

func New(service orderPayed) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("orderPayed: %+v", req)
	err := h.service.OrderPayed(ctx, req.OrderID)
	return Response{}, err
}
