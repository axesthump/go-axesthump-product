package main

import (
	"fmt"
	"log"
	"net"
	"route256/checkout/internal/api/checkout_v1"
	lomsService "route256/checkout/internal/client/grpc/loms_service"
	productService "route256/checkout/internal/client/grpc/product_service"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain/checkout"
	desc "route256/checkout/pkg/checkout_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const port = ":8080"

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}
	l, err := net.Listen("tcp", fmt.Sprintf("%v", port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	lomsConn, err := grpc.Dial(
		config.ConfigData.Services.Loms,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer lomsConn.Close()
	lomsClient := lomsService.New(lomsConn)

	productServiceConn, err := grpc.Dial(
		config.ConfigData.Services.ProductService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer productServiceConn.Close()
	productServiceClient := productService.New(productServiceConn)

	service := checkout.New(lomsClient, productServiceClient, lomsClient)
	desc.RegisterCheckoutV1Server(s, checkout_v1.New(service))
	log.Printf("server listening at %v", l.Addr())

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
