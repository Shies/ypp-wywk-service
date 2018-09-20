package dao

import (
	"encoding/base64"
	"encoding/json"
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
	"log"
)

const (
	_apiHost = "https://dy-api.eryufm.cn"
	_privateKey = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func (d *Dao) YppApi(method string, request map[string]string) string {
	u := _apiHost + "/api/index"

	reqstr, _ := json.Marshal(request)
	params := getApiSignParams(method, string(reqstr))

	res, err := d.http.HttpDo("POST", u, params.Encode(), nil)
	if err != nil {
		log.Printf("%v", err)
		return ""
	}

	return res
}

func getApiSignParams(method string, request string) url.Values {
	var (
		vNum = "44"
		timespan = "1454406417"
		params = url.Values{}
		platform = map[string]string{
			"platform": "android",
			"sys_version": "4.4.4x19",
			"soft_version": vNum,
			"device_model": "Nexus 5",
			"screen": "1080x1776",
		}
	)
	request = base64.StdEncoding.EncodeToString([]byte(request))
	params.Set("method", method)
	params.Set("request", request)
	params.Set("timespan", timespan)
	for k, v := range platform {
		params.Set(k, v)
	}
	signature := generateSignature(params, vNum)
	params.Set("signature", signature)

	return params
}

func generateSignature(params url.Values, vNum string) string {
	var (
		tmpArr []string
	)
	if vNum == "" {
		return ""
	}
	tmpArr = append(tmpArr, _privateKey)
	for _, v := range params {
		tmpArr = append(tmpArr, v[0])
	}
	sort.Strings(tmpArr)
	tmpstr := strings.Join(tmpArr, "")
	r := sha1.Sum([]byte(tmpstr))
	tmpstr = hex.EncodeToString(r[:])

	return tmpstr
}