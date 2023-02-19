package listorder

import (
	"context"
	"log"
	"route256/loms/internal/handlers"
	"route256/loms/internal/models"
)

type listOrder interface {
	ListOrder(ctx context.Context, orderID int64) (models.OrderInfo, error)
}

type (
	Request struct {
		OrderID int64 `json:"orderID"`
	}

	Response struct {
		models.OrderInfo
	}
)

func (r Request) Validate() error {
	if r.OrderID == 0 {
		return handlers.ErrEmptyOrderID
	}
	return nil
}

type Handler struct {
	service listOrder
}

func New(service listOrder) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("listOrder: %+v", req)

	orderInfo, err := h.service.ListOrder(ctx, req.OrderID)
	if err != nil {
		return Response{}, err
	}
	return Response{orderInfo}, nil
}
