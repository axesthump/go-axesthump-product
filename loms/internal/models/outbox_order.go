package models

type OutboxOrder struct {
	ID      int64         `json:"id"`
	OrderID int64         `json:"order_id"`
	Status  OrderStatusID `json:"status"`
}
