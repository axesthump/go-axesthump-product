package models

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}
