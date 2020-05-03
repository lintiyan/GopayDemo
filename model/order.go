package model

import "time"

type Order struct {
	ID          int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	OrderNum    string
	OrderStatus string
	BuyCounts   int64
	OrderAmount float64
	PaidAmount  float64
	CreateTime  time.Time
	PaidTime    time.Time
	ProductId   int64
	Product     *Product `gorm:"FOREIGNKEY:Id;ASSOCIATION_FOREIGNKEY:ProductId"`
}
