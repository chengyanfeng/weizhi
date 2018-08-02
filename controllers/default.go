package controllers

import (
	"github.com/astaxie/beego"
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

func (c *MainController) Get() {

	GetGzpt()
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


func GetGzpt(){
	userMap:=&util.StringMap{}
	(*userMap)["appid"] = def.APPID
	(*userMap)["mch_id"] = def.MCH_ID
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
	(*userMap)["appId"] = def.APPID
	(*userMap)["timeStamp"] = def.MCH_ID
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
