package model

import "time"

type Flow struct {
	ID         int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	FlowNum    string
	OrderNum   string
	ProductId  int64
	PaidAmount float64
	PaidMethod int
	BuyCounts  int64
	CreateTime time.Time
}
