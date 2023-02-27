package models

type CartInfo struct {
	Items      []Item `json:"items"`
	TotalPrice uint32 `json:"totalPrice"`
}
