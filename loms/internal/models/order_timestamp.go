package models

import "time"

type OrderTimestamp struct {
	ID       int64
	CreateAt time.Time
}
