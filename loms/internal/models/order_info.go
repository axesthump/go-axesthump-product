package models

type OrderInfo struct {
	Status OrderStatus
	User   int64
	Items  []Item
}
