package models

type OrderInfo struct {
	Status OrderStatusID
	User   int64
	Items  []Item
}
