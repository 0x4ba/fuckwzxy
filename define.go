package main

import (
	"net/http"
)

type LoginContent struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	OpenId    string `json:"openId"`
	UnionId   string `json:"unionId"`
	PhoneInfo string `json:"phoneInfo"`
}

type CheckinContent struct {
	//Answers   string `json:"answer"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Country   string `json:"country"`
	City      string `json:"city"`
	District  string `json:"district"`
	Province  string `json:"province"`
	Township  string `json:"township"`
	Street    string `json:"street"`
	Areacode  string `json:"areacode"`
}

type Userinfo struct {
	LoginContent
	CheckinContent
	CheckinTime string `json:"checkinTime"`
}

type BodyStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	LoginUrlTpl = "https://gw.wozaixiaoyuan.com/basicinfo/mobile/login/username?username=%s&password=%s&openId=%s&unionId=%s&phoneInfo=%s"
	CheckinUrl  = "https://student.wozaixiaoyuan.com/health/save.json"

	PushUrlTpl = "https://oapi.dingtalk.com/robot/send?access_token=%s"
	PushKey    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	PushTpl    = `{"msgtype": "text","text": {"content":"fuckwzxy \n%s"}}`

	SpecTpl = "0 %s * * ?"

	LoginHeader = &http.Header{
		"accept":          {"application/json, text/plain, */*"},
		"content-type":    {"application/json;charset=UTF-8"},
		"origin":          {"https://gw.wozaixiaoyuan.com"},
		"accept-language": {"en-us"},
		"user-agent":      {"Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.9(0x1800092d) NetType/WIFI Language/en miniProgram"},
		"referer":         {"https://gw.wozaixiaoyuan.com/h5/mobile/basicinfo/index/login/index"},
		"accept-encoding": {"gzip, deflate, br"},
	}
	CheckinHeader = &http.Header{
		"Host":            {"student.wozaixiaoyuan.com"},
		"content-type":    {"application/x-www-form-urlencoded"},
		"Accept-Encoding": {"gzip,compress,br,deflate"},
		"User-Agent":      {"Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.9(0x1800092d) NetType/WIFI Language/en"},
		"Referer":         {"https://servicewechat.com/wxce6d08f781975d91/176/page-"},
	}

	userinfos = make(map[string]Userinfo)
)
