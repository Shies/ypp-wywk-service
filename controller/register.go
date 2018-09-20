package controller

import (
	"fmt"
	"net/url"
	"time"
	"crypto/md5"
	"strconv"

	"ypp-wywk-service/model"
	"ypp-wywk-service/service"
	"ypp-wywk-service/ecode"
)

const (
	_forbiddenNum = "9991"
	_key = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	commonCode = []string{
		"0498",
		"0514",
		"0536",
		"0540",
		"0561",
		"0562",
		"0565",
		"0569",
		"0571",
		"0574",
		"0587",
		"0585",
		"0592",
		"0603",
		"0648",
		"0654",
		"0644",
		"0645",
		"0655",
		"0652",
		"0661",
		"0666",
		"0672",
		"0673",
		"0683",
		"0420",
		"0346",
		"0693",
		"0699",
		"0700",
		"0706",
	}
	types = map[string]string{
		"0": "开户",
		"99": "新建会员",
		"999": "下机",
	}
)

func register(c Context) {
	var (
		createTime time.Time
		alert error
		data url.Values
		status int8
		err error
	)
	alert, data = validateGetParams(c)
	if alert != nil {
		c.JSON(nil, alert)
		return
	}
	mobilephone := data.Get("mobilephone")
	custNo := data.Get("custNo")
	custName := data.Get("custName")
	cardCreateTime := service.ConvertTime(data.Get("createTime"))
	if mobilephone != "" {
		user, _ := srv.UserExists(mobilephone)
		if user != nil {
			createTime = user.CreateTime
		}
	}
	activeTime := createTime
	if createTime.IsZero() {
		status = 0
	} else {
		status = 1
	}

	switch data.Get("type") {
	case "0":
		srv.CheckIsYpp(status, mobilephone)
		srv.CheckReserve(custNo)
	case "999":
		srv.CheckReserve(custNo)
	}

	custName = service.FitlerStr(custName)
	register := &model.TWyRegister{
		TypeID: Atoi(data.Get("type")),
		IDType: data.Get("certifiType"),
		IDCard: data.Get("certifiNo"),
		Truename: custName,
		CreateTime: time.Now(),
		Mobile: mobilephone,
		CardNO: custNo,
		Status: status,
		ActiveTime: activeTime,
		CardCreateTime: cardCreateTime,
		CommonCode: data.Get("commonCode"),
	}
	isHmwy := service.InArray(data.Get("commonCode"), commonCode)
	if isHmwy {
		err = srv.SaveHMRegister(register)
	} else {
		err = srv.SaveWYRegister(register)
	}
	
	var redis = make(map[string]string)
	if err == nil && status == 0 {
		if service.IsMobile(mobilephone) {
			redis["id_card"] = data.Get("certifiNo")
			redis["truename"] = data.Get("custName")
			redis["card_no"] = data.Get("custNo")
			redis["type_id"] = data.Get("type")
			redis["is_hmwk"] = strconv.FormatBool(isHmwy)
			srv.SetUserCache(mobilephone, redis)

			var (
				action string
				source string
			)
			switch data.Get("type") {
			case "99":
				action = "anyRegisterNew"
				source = "wywk_new"
			case "0":
				action = "anyRegisterOld"
				source = "wywk_old"
			}
			if data.Get("commonCode") != _forbiddenNum {
				if (service.IsMobile(mobilephone)) {
					if res := srv.SendSms(action, mobilephone); res {
						srv.SendCount(source)
					}
				}
			}
		}
		c.JSON(nil, ecode.Success)
	} else {
		if status > 0 {
			c.JSON(nil, ecode.Success)
		} else {
			c.JSON(nil, ecode.DataRecordFail)
		}
	}
	return
}


func validateGetParams(c Context) (alert error, p url.Values) {
	var (
		data = url.Values{}
	)
	c.Request().ParseForm()
	for k, v := range c.Request().Form {
		v, _ := url.QueryUnescape(v[0])
		data.Set(k, v)
	}

	if _, ok := types[data.Get("type")]; !ok {
		alert = ecode.InvalidState
		return
	}

	sign := data.Get("sign")
	timestamp := data.Get("timestamp")
	mobilephone := data.Get("mobilephone")
	if data.Get("type") == "999" {
		alert = ecode.InvalidState
		return
	}

	srv.SaveRegister(&model.TRegister{Data: data.Encode(), CreateTime: time.Now()})
	if !service.IsMobile(mobilephone) {
		alert = ecode.LeakRequiredField
		return
	}
	if timestamp == "" {
		alert = ecode.LeakRequiredField
		return
	}
	if service.ValidateIsTimeout(ParseInt(timestamp)) {
		alert = ecode.RequestTimeout
		return
	}

	data.Del("sign")
	signstr := data.Encode()+"&key="+_key

	bytes := []byte(signstr)
	has := md5.Sum(bytes)
	md5sign := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if sign != md5sign {
		alert = ecode.SignVerifyFail
		return
	}

	return nil, data
}
