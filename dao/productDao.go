package dao

import (
	"fmt"
	"go_pay/model"
	"go_pay/util"
)

func GetProductById(productId int64) (*model.Product, error) {

	err := processError()
	var product model.Product
	util.DB.Table("products").Where("id=?", productId).Find(&product)
	return &product, err
}

func GetProducts() (products []*model.Product, err error) {
	err = processError()

	util.DB.Table("products").Find(&products)
	return products, err
}

func processError() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("db err :%v\n", p)
			util.ErrorLog.Errorf("DB ERR :%v\n", err)
		}
	}()
	return nil
}
