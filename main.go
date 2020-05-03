package main

import (
	"go_pay/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//r.LoadHTMLFiles("views/pages/error.html")
	r.LoadHTMLGlob("views/pages/*")

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	//获取产品列表
	r.GET("/getProducts", controller.GetProducts)

	//选中商品购买
	r.GET("/buy", controller.Buy)

	//创建订单
	r.POST("/createOrder", controller.CreateOrder)

	//支付宝支付
	r.GET("/alipay", controller.Aplipay)

	r.GET("/return", controller.ApliPayReturn)

	r.POST("/notify", controller.ApliPayNotify)

	//全局的异常处理器
	r.GET("/error", controller.Error)
	r.Run(":8081")
}
