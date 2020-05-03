package controller

import (
	"fmt"
	"go_pay/service"
	"go_pay/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	//获取请求中的创建订单所需信息
	productId := c.PostForm("productId")
	buyNum := c.PostForm("buyCounts")

	fmt.Printf("productId:%v  buyNum:%v\n", productId, buyNum)

	productIdInt, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		util.ErrorLog.Errorf("parse string productId :[%s] to int err :%v", productId, err)
	}
	buyNumInt, err := strconv.ParseInt(buyNum, 10, 64)
	if err != nil {
		util.ErrorLog.Errorf("parse string buyNum :[%s] to int err :%v", buyNumInt, err)
	}

	//调用service创建订单
	order, err2 := service.CreateOrder(productIdInt, buyNumInt)
	if err2 != nil || order == nil {
		msg := "订单创建失败"
		ErrorMsg(c, msg)
	}

	fmt.Println("order:", order)

	//将创建的订单信息返回到页面
	c.HTML(http.StatusOK, "order.html", gin.H{
		"order": *order,
	})
}
