package controller

import (
	"go_pay/service"
	"go_pay/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Buy 购买
func Buy(c *gin.Context) {
	//获取商品id
	productId := c.Query("productId")
	productIdInt, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		//fmt.Println("string to int err")
		util.ErrorLog.Errorf("string to int err:%v", err)
	}
	//fmt.Println("productIdInt",productIdInt)
	//根据商品id查询商品信息
	product, err2 := service.GetProductById(productIdInt)
	if err2 != nil {
		msg := "获取商品信息失败"
		ErrorMsg(c, msg)
	}

	if product == nil {
		msg2 := "不存在该商品"
		ErrorMsg(c, msg2)
	}

	//fmt.Println("product", product)

	//将商品信息返回给页面
	c.HTML(http.StatusOK, "product_buy.html", gin.H{
		"product": product,
	})
}

//GetProducts 获取商品列表
func GetProducts(c *gin.Context) {
	//调用service方法查询获得产品列表信息
	products, err := service.GetProducts()
	if err != nil {
		msg := "获取产品列表失败"
		ErrorMsg(c, msg)
	}

	//fmt.Println("products==========", products)

	c.HTML(http.StatusOK, "products.html", gin.H{
		"products": products,
	})

}
