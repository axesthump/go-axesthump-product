package deletefromcart

import (
	"context"
	"log"
	"route256/checkout/internal/handlers"
)

type deleteFromCart interface {
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
}

type Handler struct {
	service deleteFromCart
}

type (
	Request struct {
		User  int64  `json:"user"`
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	}
	Response struct {
	}
)

func New(service deleteFromCart) *Handler {
	return &Handler{service: service}
}

func (r Request) Validate() error {
	if r.User == 0 {
		return handlers.ErrEmptyUser
	}
	if r.Sku == 0 {
		return handlers.ErrEmptySku
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("deleteFromCart: %+v", req)

	var response Response

	err := h.service.DeleteFromCart(ctx, req.User, req.Sku, req.Count)
	if err != nil {
		return response, err
	}

	return response, nil
}
