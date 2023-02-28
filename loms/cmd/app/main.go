package main

import (
	"fmt"
	"log"
	"net"
	lomsV1 "route256/loms/internal/api/loms_v1"
	"route256/loms/internal/domain/loms"
	desc "route256/loms/pkg/loms_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8081"

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("%v", port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	service := loms.New()
	desc.RegisterLomsV1Server(s, lomsV1.New(service))
	log.Printf("server listening at %v", l.Addr())

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
