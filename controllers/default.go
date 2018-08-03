package controllers

import (
	"github.com/astaxie/beego"
	"github.com/smartwalle/alipay"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	_"weizhi/util"
	"weizhi/util"
	"weizhi/def"
)

type MainController struct {
	beego.Controller
}
var client = alipay.New(def.ZHIFUBAOAPPID, "2088802940812132", def.ZHIFUBAO_KEY, def.ZHIFUBAOprivateKey, false)

func (c *MainController) Get() {
	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = "http://192.144.176.213:8070/return"
	p.ReturnURL = "http://192.144.176.213:8070/apliy"
	p.Subject = "这是测试"
	p.OutTradeNo = "2342341233121w3q2eq131w2"
	p.TotalAmount = "10.00"
	p.ProductCode = "商品编码"

	var html, _ = client.TradeWapPay(p)

	fmt.Print(html)
	c.Data["html"] = html
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func(c *MainController) GetUrl() string{
	req1 :=c.Ctx.Request
	fmt.Print(req1,"--------------req1-------------")
	fmt.Print(req1.Form,"------------req1.Form----------")
	ok, err := client.VerifySign(req1.Form)
	fmt.Println(ok, err)
	return "success"

}
func GetGzpt(){
	userMap:=&util.StringMap{}
	(*userMap)["appid"] = def.WEIXINAPPID
	(*userMap)["mch_id"] = def.WEIXINMCH_ID
	(*userMap)["nonce_str"] = util.GetRandomString()
	(*userMap)["body"] = "erewrwe"
	(*userMap)["out_trade_no"] = "123456"
	(*userMap)["total_fee"] = "12"
	(*userMap)["spbill_create_ip"] = "123.12.12.123"
	(*userMap)["trade_type"] = "APP"
	(*userMap)["notify_url"]="http://www.weixin.qq.com/wxpay/pay.php"
	(*userMap)["sign_type"]="MD5"
	xml:=util.MapToxml(userMap)
	response, _ := http.Post("https://api.mch.weixin.qq.com/sandbox/pay/unifiedorder", "application/xml;charset=utf-8", strings.NewReader(xml))
	defer response.Body.Close()
	token_body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(token_body))
}
func GetH5(){
	userMap:=&util.StringMap{}
	(*userMap)["appId"] = def.WEIXINAPPID
	(*userMap)["timeStamp"] = "21312"
	(*userMap)["nonceStr"] = util.GetRandomString()
	(*userMap)["package"] = "erewrwe"
	(*userMap)["sign_type"]="MD5"
	(*userMap)["paySign"]="MD5"
	xml:=util.MapToxml(userMap)
	response, _ := http.Post("https://api.mch.weixin.qq.com/sandbox/pay/unifiedorder", "application/xml;charset=utf-8", strings.NewReader(xml))
	defer response.Body.Close()
	token_body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(token_body))
}

func (c *MainController)ZHIFUBAO(){
	var client = alipay.New(def.ZHIFUBAOAPPID, "132123", def.ZHIFUBAO_KEY, def.ZHIFUBAOprivateKey, false)
	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = "http://www.baidu.com"
	p.Subject = "这是测试"
	p.OutTradeNo = "23423423121wqeqw"
	p.TotalAmount = "10.00"
	p.ProductCode = "商品编码"

	var html, _ = client.TradeWapPay(p)
	fmt.Print(html)
	c.Data["html"] = html
	// 将html输出到浏览器

}