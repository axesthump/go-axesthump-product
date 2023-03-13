package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	lomsV1 "route256/loms/internal/api/loms_v1"
	"route256/loms/internal/config"
	"route256/loms/internal/domain/loms"
	"route256/loms/internal/repository/postgres"
	desc "route256/loms/pkg/loms_v1"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8081"

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("%v", port))
	if err != nil {
		log.Fatal(err)
	}
	err = config.Init()
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	psqlConn := config.ConfigData.Postgres.Url
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	repository := postgres.New(pool)

	service := loms.New(repository)
	desc.RegisterLomsV1Server(s, lomsV1.New(service))
	log.Printf("server listening at %v", l.Addr())

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
