package checkout_v1

import (
	"context"
	"route256/checkout/internal/models"
	desc "route256/checkout/pkg/checkout_v1"
)

type Service interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error
	ListCart(ctx context.Context, user int64) (models.CartInfo, error)
	Purchase(ctx context.Context, user int64) (int64, error)
}

type Handler struct {
	desc.UnimplementedCheckoutV1Server
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}
