package purchase

import (
	"context"
	"log"
	"route256/checkout/internal/handlers"
)

type purchase interface {
	Purchase(ctx context.Context, user int64) error
}

type Handler struct {
	service purchase
}

type (
	Request struct {
		User int64 `json:"user"`
	}
	Response struct {
	}
)

func New(service purchase) *Handler {
	return &Handler{service: service}
}

func (r Request) Validate() error {
	if r.User == 0 {
		return handlers.ErrEmptyUser
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("purchase: %+v", req)
	err := h.service.Purchase(ctx, req.User)
	return Response{}, err
}
