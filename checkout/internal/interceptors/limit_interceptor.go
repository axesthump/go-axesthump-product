package interceptors

import (
	"context"
	"route256/checkout/internal/limiter"
	"time"

	"google.golang.org/grpc"
)

func LimitInterceptor(limiter *limiter.Limiter) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()

		err := limiter.Wait(ctxTimeout)
		if err != nil {
			return err
		}
		return invoker(ctxTimeout, method, req, reply, cc, opts...)
	}
}
