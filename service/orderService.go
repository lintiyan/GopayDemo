package service

import (
	"go_pay/dao"
	"go_pay/model"
	"go_pay/util"
	"time"
)

func CreateOrder(productId int64, buyCount int64) (*model.Order, error) {
	//创建订单
	order := &model.Order{
		OrderNum:    util.GenerateId(),
		OrderStatus: "待付款",
		BuyCounts:   buyCount,
		CreateTime:  time.Now(),
		ProductId:   productId,
	}
	//查询商品信息
	product, err := dao.GetProductById(productId)
	if product == nil || err != nil {
		util.ErrorLog.Errorf("查询商品信息失败，商品id：%v", productId)
		return nil, err
	}

	//填充订单
	order.Product = product

	orderAmount := product.Price * float64(buyCount)
	order.OrderAmount = orderAmount

	//存储订单
	dao.SaveOrder(order)

	return order, nil
}
