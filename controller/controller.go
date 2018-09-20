package controller

import (
	"log"
	"strconv"
	"net/http"
	"html/template"
	"context"
	"encoding/json"
	"fmt"

	"ypp-wywk-service/service"
	"ypp-wywk-service/conf"
	"ypp-wywk-service/ecode"
	"time"
)

var (
	srv *service.Service
	cstZone = time.FixedZone("CST", 8*3600)       // 东八
	_ = time.Now().In(cstZone).Format("2006-01-02 15:04:05")
)

type Context interface {
	context.Context
    Request() *http.Request
    Response() http.ResponseWriter
    SetPath(string)
    SetData(interface{})
    GetPath() string
    GetData() interface{}
	JSON(interface{}, error)
}

type implContext struct {
	context.Context
    req  *http.Request
    res  http.ResponseWriter
    path string
    data interface{}
	err error
}

type Response struct {
	Code int64		 `json:"code"`
	Message string	 `json:"message"`
	Data interface{} `json:"data"`
}

func (ic *implContext) Request() *http.Request {
    return ic.req
}

func (ic *implContext) Response() http.ResponseWriter {
    return ic.res
}

func (ic *implContext) SetPath(path string) {
    ic.path = path
}

func (ic *implContext) SetData(data interface{}) {
    ic.data = data
}

func (ic *implContext) GetPath() string {
	return ic.path
}

func (ic *implContext) GetData() interface{} {
	return ic.data
}

func (ic *implContext) JSON(data interface{}, err error) {
	code, msg := ecode.GetException(err)
	bytes, e := json.Marshal(&Response{Code: code, Message: msg, Data:data})
	ic.err = e
	fmt.Fprintf(ic.Response(), string(bytes))
}

func ParseInt(value string) int64 {
	intval, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		intval = 0
	}

	return intval
}

func Atoi(value string) int {
	intval, err := strconv.Atoi(value)
	if err != nil {
		intval = 0
	}

	return intval
}

func views(c Context) error {
	t, err := template.ParseFiles("./" + c.GetPath())
	if err != nil {
		panic(err)
	}

	return t.Execute(c.Response(), c.GetData())
}

func initService() {
	srv = service.New(&conf.Conf)
}

func Register() bool {
	conf.ParseConfig()
	conf.Logger(conf.Conf.Log.Dir)
	initService()
	setHttpHandle()

	// 设置监听端口
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	return true
}

func setHttpHandle() error {
    var (
        err error
        router = make(map[string]interface{})
    )

    router["register"] = register
    for route, function := range router {
        f := function.(func(Context))
        http.HandleFunc("/" + route, func(w http.ResponseWriter, req *http.Request) {
        	f(&implContext{req: req, res: w})
        })
    }

	return err
}
