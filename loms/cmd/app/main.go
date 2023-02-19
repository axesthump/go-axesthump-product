package main

import (
	"log"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/internal/domain"
	"route256/loms/internal/handlers/cancelorder"
	"route256/loms/internal/handlers/createorder"
	"route256/loms/internal/handlers/listorder"
	"route256/loms/internal/handlers/orderpayed"
	"route256/loms/internal/handlers/stocks"
)

const port = ":8081"

func main() {
	service := domain.New()
	createOrderHandler := createorder.New(service)
	listOrderHandler := listorder.New(service)
	orderPayedHandler := orderpayed.New(service)
	cancelOrderHandler := cancelorder.New(service)
	stocksHandler := stocks.New(service)
	http.Handle("/createOrder", srvwrapper.New(createOrderHandler.Handle))
	http.Handle("/listOrder", srvwrapper.New(listOrderHandler.Handle))
	http.Handle("/orderPayed", srvwrapper.New(orderPayedHandler.Handle))
	http.Handle("/cancelOrder", srvwrapper.New(cancelOrderHandler.Handle))
	http.Handle("/stocks", srvwrapper.New(stocksHandler.Handle))

	log.Println("listening http at", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
