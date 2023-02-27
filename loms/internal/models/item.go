package models

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}
