package client

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

var req *resty.Request

func initR(isJson ...bool) {
	var headers = make(map[string]string)
	if len(isJson) > 0 && !isJson[0] {

	} else {
		headers["Accept"] = "application/json"
	}

	//headers["Authorization"] = "Bearer iUfi6qzGU1r8u1bfC6MukcM1qZxtiDyrRQy2HCROa0vpcM4z1KYmaGg9cMB6ceNqgyODB7nXkuyLPTzWSjHMCZnnk-gxK3diF0b-oq8oS4x1hHYQf42DJaHjhIDAwYXwnEoIaI2L0sBxpzai10sl8HsPQ1mQQH-we0ejLik7PdidCAnKMxOfLQle6hAa-eSfRtsCQARQeLSFd4m8-sJoErFGoGoaF7yINlIYSv74TYQ4_YxWZgkzNqs1U77J0Lh-kPN792ZRLsJ-5M1NGQTATiVHixVNfSVtlMc0GjJfFDjDfnl8HigfETYoR90d-AW4JQy8QO7EWXFbxKYBywKrYw"
	headers["TenantGuid"] = viper.GetString("tenantGuid")
	headers["DomainGuid"] = viper.GetString("domainGuid")
	req = resty.R().SetHeaders(headers).SetAuthToken(viper.GetString("token"))
}

func Get(url string) *resty.Response {
	initR()
	url = fmt.Sprintf("http://%s/ServicesAPI/%s", viper.GetString("apiServer"), url)
	res, error := req.Get(url)
	if error != nil {
		panic("rest connection failed")
	}
	return res
}

func GetPlainTxt(url string) *resty.Response {
	initR(false)
	url = fmt.Sprintf("http://%s/ServicesAPI/%s", viper.GetString("apiServer"), url)
	res, error := req.Get(url)
	if error != nil {
		panic("rest connection failed")
	}
	return res
}

func Post(url string, body interface{}) *resty.Response {
	initR()
	url = fmt.Sprintf("http://%s/ServicesAPI/%s", viper.GetString("apiServer"), url)
	res, error := req.SetBody(body).Post(url)
	if error != nil {
		panic("rest connection failed")
	}
	return res
}
