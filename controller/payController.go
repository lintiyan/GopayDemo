package controller

import (
	"fmt"
	"go_pay/service"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func Aplipay(c *gin.Context) {
	//获取订单编号
	orderNum := c.Query("orderNum")

	service.WebPageAlipay(orderNum)
}

func ApliPayReturn(c *gin.Context) {
	fmt.Println("return ....................")
	order, ok := service.Return(c)
	if ok {
		//跳转到订单页面
		c.HTML(http.StatusOK, "pay_success.html", gin.H{
			"order": order,
		})
	} else {
		msg := "支付失败"
		ErrorMsg(c, msg)
	}
	//获取支付信息

	//paramsRes := ""

	// urlReturn := c.Request.URL.String()
	// fmt.Println("res:", urlReturn)
	// strs := strings.Split(urlReturn, "?")
	// params, err := url.ParseQuery(strs[1])
	// if err != nil {
	// 	util.ErrorLog.Errorf("Parse param err : %v", err)
	// 	ErrorMsg(c, "页面访问错误")
	// }
	// fmt.Println("parmas:", params)
	// maps := make(map[string]interface{}, 0)

	// for k, v := range params {
	// 	if k == "sign" || k == "sign_type" {
	// 		continue
	// 	}
	// 	maps[k] = v[0]
	// }

	// //获取签名
	// sign := params["sign"][0]

	// fmt.Println("maps:", maps)
	// //得到去除sign和sign_type的参数字符串
	// signStr := getSignStr(maps)
	// //验签
	// success := verifySign(signStr, sign)

	//跳转页面

}

func ApliPayNotify(c *gin.Context) {
	fmt.Println("notify ....................")
	fmt.Println("================================================================")
	//获取支付信息

	//验签

	//更新订单信息

	//返回success

}

//将maps里的键值对按键排序 后拼接成string返回
func getSignStr(maps map[string]interface{}) string {

	keys := make([]string, 0)
	for k, _ := range maps {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	resStr := ""

	index := 0
	for _, v := range keys {
		resStr += fmt.Sprintf("%v", maps[v])
		if index < len(keys)-1 {
			resStr += "&"
		}
		index++
	}

	return resStr
}
