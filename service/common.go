package service

import (
	"regexp"
	"strings"
	"math"
	"time"
	"crypto/md5"
	"fmt"

	"github.com/satori/go.uuid"
)

const (
	_baseFormat = "2006-01-02 15:04:05"
	_regular = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\\\d{8}$"
)

func UUid() string {
	v := uuid.Must(uuid.NewV4()).String()
	has := md5.Sum([]byte(v))
	md5val := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5val
}

func ConvertTime(str string) time.Time {
	timer, _ := time.Parse(_baseFormat, str)
	return timer
}

func IsMobile(mobileNum string) bool {
	reg := regexp.MustCompile(_regular)
	return reg.MatchString(mobileNum)
}

func InArray(code string, commonCode []string) bool {
	var isHmwk = false
	for _, v := range commonCode {
		if v == code {
			isHmwk = true
			break
		}
	}

	return isHmwk
}

func FitlerStr(custName string) string {
	custName = strings.Replace(custName, "(改)", "", -1)
	custName = strings.Replace(custName, "(读)", "", -1)
	custName = strings.Replace(custName, "(扫)", "", -1)
	custName = strings.Replace(custName, "（读）", "", -1)
	custName = strings.Replace(custName, "（扫改）", "", -1)
	custName = strings.Replace(custName, "（读改）", "", -1)
	custName = strings.Replace(custName, "(扫)(改)", "", -1)
	custName = strings.Replace(custName, "（读）（读改）", "", -1)
	custName = strings.Replace(custName, "(手)", "", -1)

	return custName
}

func ValidateIsTimeout(timef int64) bool {
	timef = int64(math.Round(float64(timef) / 1000))
	if (time.Now().Unix() - timef) > 60 {
		return true
	}

	return false
}
