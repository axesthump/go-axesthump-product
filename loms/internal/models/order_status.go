package models

type OrderStatus int

const (
	Nil = iota
	New
	AwaitingPayment
	Failed
	Payed
	Cancelled
)
