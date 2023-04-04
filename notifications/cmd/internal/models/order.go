package models

type Order struct {
	OrderID int64         `json:"order_id"`
	Status  OrderStatusID `json:"status"`
}
