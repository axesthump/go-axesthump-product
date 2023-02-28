package loms_service

import (
	lomsAPI "route256/loms/pkg/loms_v1"

	"google.golang.org/grpc"
)

type Client struct {
	lomsClient lomsAPI.LomsV1Client
}

func New(cc *grpc.ClientConn) *Client {
	return &Client{
		lomsClient: lomsAPI.NewLomsV1Client(cc),
	}
}
