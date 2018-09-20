package ecode

import (
	"fmt"
	"strconv"
)

type errorCode int

var (
	codes = map[string]map[string]string{
		"0": {"code": "0", "message": "成功"},
		"40001": {"code": "40001", "message": "签名验证失败"},
		"40002": {"code": "40002", "message": "请求超时"},
		"40003": {"code": "40003", "message": "缺少必要参数"},
		"40004": {"code": "40004", "message": "状态不处理"},
		"50001": {"code": "50001", "message": "数据记录失败"},
		"90001": {"code": "90001", "message": "余额充足"},
		"90002": {"code": "90002", "message": "鱼泡泡用户不发送"},
		"90003": {"code": "90003", "message": "已发送过短信"},
	}
)

const (
	Failed errorCode = 400
	Success errorCode = 0
	SignVerifyFail errorCode = 40001
	RequestTimeout errorCode = 40002
	LeakRequiredField errorCode = 40003
	InvalidState errorCode = 40004
	DataRecordFail errorCode = 50001
	BalanceEnough errorCode = 90001
	YppUserNotSend errorCode = 90002
	SmsSent errorCode  = 90003
)

func (v errorCode) Code() int64 {
	return int64(v)
}

func (v errorCode) Error() string {
	var format string
	intstr := strconv.Itoa(int(v))
	if vv, ok := codes[intstr]; ok {
		format = fmt.Sprintf("%s", vv["message"])
	} else {
		format = fmt.Sprintf("%s", "Request Error")
	}

	return format
}

func GetException(err error) (int64, string) {
	if err == nil {
		return int64(Success), string("")
	}

	if v, ok := err.(errorCode); ok {
		return v.Code(), v.Error()
	}
	return int64(Failed), err.Error()
}