package createorder

import (
	"context"
	"log"
	"route256/loms/internal/handlers"
	"route256/loms/internal/models"
)

type createOrder interface {
	CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error)
}

type (
	Request struct {
		User  int64         `json:"user"`
		Items []models.Item `json:"items"`
	}

	Response struct {
		OrderID int64 `json:"orderID"`
	}
)

func (r Request) Validate() error {
	if r.User == 0 {
		return handlers.ErrEmptyUser
	}
	return nil
}

type Handler struct {
	service createOrder
}

func New(service createOrder) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("createOrder: %+v", req)

	var response Response

	orderID, err := h.service.CreateOrder(ctx, req.User, req.Items)
	if err != nil {
		return response, err
	}
	response.OrderID = orderID
	return response, nil
}
