package models

type OrderStatusID int

const (
	UnknownOrderStatusID OrderStatusID = iota
	New
	AwaitingPayment
	Failed
	Payed
	Cancelled
)
