package product_service

import (
	productServiceAPI "route256/checkout/pkg/product_service_v1"

	"google.golang.org/grpc"
)

type Client struct {
	productServiceClient productServiceAPI.ProductServiceClient
}

func New(cc *grpc.ClientConn) *Client {
	return &Client{
		productServiceClient: productServiceAPI.NewProductServiceClient(cc),
	}
}
