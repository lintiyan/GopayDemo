package service

import (
	"fmt"
	"go_pay/consts"
	"go_pay/dao"
	"go_pay/model"
	"go_pay/util"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay"
)

var (
	client *alipay.Client
	err    error
)

func init() {
	initAlipay()
}

func initAlipay() {
	client, err = alipay.New(consts.AppID, consts.PrivateKey, false)

	client.LoadAppPublicCertFromFile(consts.AppCertPublicKey)       // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile(consts.AlipayRootCert)        // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile(consts.AlipayCertPublicKey) // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		util.ErrorLog.Errorf("init alipay failed,err:%v ", err)
	}
}

func Return(c *gin.Context) (*model.Order, bool) {
	//验签
	c.Request.ParseForm()
	ok, err := client.VerifySign(c.Request.Form)
	if err != nil {
		util.ErrorLog.Errorf("verify alipay return failed,err:%v ", err)
		return nil, false
	}
	//验签通过
	if ok {
		fmt.Println("verify success")

		//获取订单号
		orderNum := c.Query("out_trade_no")
		//获取订单金额
		totalAmount := c.Query("total_amount")
		//获取卖家id
		sellerId := c.Query("seller_id")
		//获取app_id
		appId := c.Query("app_id")
		//验证内容数据
		order, err := dao.GetOrderByOrderNum(orderNum)
		if err != nil {
			util.ErrorLog.Errorf("get order err", err)
		}

		ok, err := verifyContent(order, orderNum, totalAmount, sellerId, appId)
		if ok {
			//更新订单信息
			order.OrderStatus = "已支付"
			order.PaidAmount = order.OrderAmount
			order.PaidTime = time.Now()
			dao.UpdateOrder(order)

			//生成流水
			flow := &model.Flow{
				FlowNum:    util.GenerateId(),
				OrderNum:   order.OrderNum,
				ProductId:  order.ProductId,
				PaidAmount: order.PaidAmount,
				PaidMethod: 1,
				BuyCounts:  order.BuyCounts,
				CreateTime: time.Now(),
			}
			dao.SaveFlow(flow)
			return order, true
		} else {
			util.ErrorLog.Errorf("verify order content failed,err:%v ", err)
			return nil, false
		}

	} else {
		util.ErrorLog.Errorf("verify alipay sign failed,err:%v ", err)
		return nil, false
	}
}

func WebPageAlipay(orderNum string) {

	//根据订单编号查询订单信息
	order, _ := dao.GetOrderByOrderNum(orderNum)

	//fmt.Println("order :", order)

	//fmt.Println("==================")
	pay := alipay.TradePagePay{}
	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://localhost:8081/notify"
	// 支付成功之后，浏览器将会重定向到该 URL
	pay.ReturnURL = "http://localhost:8081/return"
	//支付标题
	pay.Subject = fmt.Sprintf("购买商品%s %d件", order.Product.Name, order.BuyCounts)
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = order.OrderNum
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = consts.AliProductCode
	//金额  保留小数点后两位
	pay.TotalAmount = fmt.Sprintf("%.2f", order.OrderAmount)

	url, err := client.TradePagePay(pay)
	if err != nil {
		util.ErrorLog.Errorf("alipay  failed,err:%v ", err)
	}
	payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)

	payURL = strings.Replace(payURL, "&", "^&", -1)
	exec.Command("cmd", "/c", "start", payURL).Start()
}


//验证通知中商户信息是否跟商户服务器中的信息一致
func verifyContent(order *model.Order, orderNum string, totalAmount string, sellerId string, appId string) (b bool, err error) {

	taotalAmountFlo, err := strconv.ParseFloat(totalAmount, 64)
	if err != nil {
		util.ErrorLog.Errorf("string to float err :%v ", err)
	}

	if order.OrderAmount == taotalAmountFlo && consts.SellerId == sellerId && consts.AppID == appId {
		b = true
		return
	}
	return
}
