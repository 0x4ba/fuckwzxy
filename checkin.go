package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"fuckwzxy/utils"
)

func Login(loginInfo LoginContent) string {
	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf(LoginUrlTpl,
		loginInfo.Username,
		loginInfo.Password,
		loginInfo.OpenId,
		loginInfo.UnionId,
		loginInfo.PhoneInfo),
		strings.NewReader("{ }"))
	if err != nil {
		Logger.Fatalln("create request error", err)
		os.Exit(-1)
	}

	req.Header = *LoginHeader
	resp, err := client.Do(req)
	if err != nil {
		Logger.Fatalln("requesting error", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()
	bodyStatus := &BodyStatus{}
	body, _ := ioutil.ReadAll(resp.Body)

	err = Json.Unmarshal(body, bodyStatus)
	if err != nil {
		Logger.Fatalln("unmarshal body error", err)
		return ""
	}
	if resp.StatusCode != http.StatusOK || bodyStatus.Code != 0 {
		Logger.Fatalln("Login failed", resp.StatusCode, "  ", bodyStatus.Code, "  ", bodyStatus.Message)
		return ""
	}
	Logger.Println("get Jwsession succeess:", loginInfo.Username, "   ", resp.Header["Jwsession"][0])

	PushMsg(loginInfo.Username + "  logined   " + "Jwsession:" + resp.Header["Jwsession"][0])
	return resp.Header["Jwsession"][0]
}

func Checkin(info CheckinContent, jwsession string) {
	//CheckinContentTpl = `answers=%s&latitude=%s&longitude=%s&country=%s&city=%s&district=%s&province=%s&township=%s&street=%s&areacode=%s`
	client := &http.Client{}

	u := url.Values{
		"answers":   {utils.RandomTemperature()},
		"latitude":  {info.Latitude},
		"longitude": {info.Longitude},
		"country":   {info.Country},
		"city":      {info.City},
		"district":  {info.District},
		"province":  {info.Province},
		"township":  {info.Township},
		"street":    {info.Street},
		"areacode":  {info.Areacode},
	}
	body := utils.UrlEncode(u)

	req, err := http.NewRequest("POST", CheckinUrl, strings.NewReader(body))
	if err != nil {
		Logger.Fatalln("create request error", err)
		os.Exit(-1)
	}
	req.Header = *CheckinHeader
	req.Header.Add("JWSESSION", jwsession)

	resp, err := client.Do(req)
	if err != nil {
		Logger.Fatalln("request error", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()
	bodyStatus := &BodyStatus{}
	RespBody, _ := ioutil.ReadAll(resp.Body)

	err = Json.Unmarshal(RespBody, bodyStatus)
	if err != nil {
		Logger.Fatalln("unmarshal body error", err)
		return
	}
	if resp.StatusCode != http.StatusOK || bodyStatus.Code != 0 {
		Logger.Fatalln("打卡失败", err)
		PushMsg(jwsession + ": 打卡失败")
		return
	}
	PushMsg(jwsession + ": 打卡成功")
}

func LoadUserInfo() {
	f, err := os.Open("./userinfo.json")
	if err != nil {
		Logger.Fatalln("open userinfo.json failed", err)
		os.Exit(-1)
	}
	fileContext, err := ioutil.ReadAll(f)
	if err != nil {
		Logger.Fatalln("read userinfo.json failed", err)
		os.Exit(-1)
	}

	err = Json.Unmarshal(fileContext, &userinfos)
	if err != nil {
		Logger.Fatalln("userinfo Unmarshal failed", err)
		os.Exit(-1)
	}
}

func PushMsg(msg string) {
	url := fmt.Sprintf(PushUrlTpl, PushKey)
	body := fmt.Sprintf(PushTpl, msg)

	client := &http.Client{}
	client.Post(url, "application/json", strings.NewReader(body))
}
