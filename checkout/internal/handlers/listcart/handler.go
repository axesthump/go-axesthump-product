package listcart

import (
	"context"
	"log"
	"route256/checkout/internal/handlers"
	"route256/checkout/internal/models"
)

type listCart interface {
	ListCart(ctx context.Context, user int64) (*models.CartInfo, error)
}

type Handler struct {
	service listCart
}

type (
	Request struct {
		User int64 `json:"user"`
	}
	Response struct {
		models.CartInfo
	}
)

func New(service listCart) *Handler {
	return &Handler{service: service}
}

func (r Request) Validate() error {
	if r.User == 0 {
		return handlers.ErrEmptyUser
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("listCart: %+v", req)

	var response Response
	cartInfo, err := h.service.ListCart(ctx, req.User)
	if err != nil {
		return response, err
	}

	response.CartInfo = *cartInfo
	return response, nil
}
