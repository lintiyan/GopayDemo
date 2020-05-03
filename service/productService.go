package service

import (
	"go_pay/dao"
	"go_pay/model"
)

func GetProductById(productId int64) (*model.Product, error) {
	return dao.GetProductById(productId)
}

func GetProducts() ([]*model.Product, error) {

	return dao.GetProducts()
	//return nil, errors.New("err")
}
