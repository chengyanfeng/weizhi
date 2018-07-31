package controllers

import (
	"github.com/astaxie/beego"
	"encoding/xml"
	"strings"
	"fmt"
	"sort"
	"crypto/md5"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	userMap := make(map[string]string)
	userMap["appid"] = "wx53d52d70ccd6439f"
	userMap["mch_id"] = "1484906012"
	userMap["nonce_str"] = "1484324dsfesfefe906012"
	userMap["body"] = "测试"
	userMap["out_trade_no"] = "123456"
	userMap["total_fee"] = "12"
	userMap["trade_type"] = "JSAPI"
	userMap["spbill_create_ip"] = "127.0.0.1"
	userMap["notify_url"] = "127.0.0.1"

	sign := c.GetSign(userMap)
	userMap["sign"] = sign
	buf, _ := xml.Marshal(StringMap(userMap))
	xml := string(buf)
	xml = strings.Replace(xml, "StringMap", "xml", -1)
	response, _ := http.Post("https://api.mch.weixin.qq.com/pay/unifiedorder", "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(xml))
	defer response.Body.Close()
	token_body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(token_body))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetSign(p StringMap) string {
	md := md5.New()
	var sign = ""
	strs := []string{"appid", "mch_id", "nonce_str", "body", "out_trade_no", "total_fee", "trade_type","spbill_create_ip","notify_url"}
	sort.Strings(strs)
	for _, v := range strs {
		sign = sign + v + "=" + p[v] + "&"
	}
	sign = sign + "key=dfb513840c45e387cd869af3887e69cb"
	fmt.Print(sign)
	md.Write([]byte(sign))
	sign = fmt.Sprintf("%x", md5.Sum([]byte(sign)))
	return strings.ToUpper(sign)

}
func JsonEncode(v interface{}) (r string) {
	b, err := json.Marshal(v)
	if err != nil {
		Error(err)
	}
	r = string(b)
	return
}

// 记录err信息
func Error(v ...interface{}) {
	beego.Error(v)
}
func JsonDecode(b []byte) (p *map[string]interface{}) {
	p = &map[string]interface{}{}
	err := json.Unmarshal(b, p)
	if err != nil {
		Error("JsonDecode", string(b), err)
	}
	return
}
