package dao

import (
	"go_pay/model"
	"go_pay/util"
	_ "go_pay/util"
)

func SaveOrder(order *model.Order) bool {
	//util.DB.Model(order).Related()
	util.DB.Table("orders").Create(order)
	return true
}

func GetOrderByOrderNum(orderNum string) (*model.Order, error) {
	var order model.Order
	util.DB.Table("orders").Where("order_num=?", orderNum).Find(&order)
	var product model.Product
	util.DB.Table("products").Where("id= ?", order.ProductId).Find(&product)
	order.Product = &product
	return &order, nil
}

func UpdateOrder(order *model.Order) {
	util.DB.Table("orders").Save(order)
}
