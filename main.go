package main

import (
	"fmt"
	"io"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron"
)

var (
	Logger *log.Logger
	Json   jsoniter.API
)

func init() {
	userInfo, err := os.OpenFile("fuckwzxy.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("loading userinfo.json error ", err)
		os.Exit(-1)
	}
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
	Logger = log.New(io.MultiWriter(os.Stdout, userInfo), "", log.Lshortfile|log.Ltime)
}

func main() {

	c := cron.New()

	LoadUserInfo()
	for k, v := range userinfos {
		fmt.Println(k, v)

		spec := fmt.Sprintf(SpecTpl, v.CheckinTime)
		c.AddFunc(spec, func() {
			token := Login(userinfos[k].LoginContent)
			if token != "" {
				Checkin(v.CheckinContent, token)
			} else {
				PushMsg(k + "登陆失败")
			}

		})
	}
	c.Start()
	select {}
}
