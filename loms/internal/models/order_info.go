package models

type OrderInfo struct {
	Status string `json:"status"`
	User   int64  `json:"user"`
	Items  []Item `json:"items"`
}
