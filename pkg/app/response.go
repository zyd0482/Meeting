package app

import (
	"github.com/gin-gonic/gin"

	"meeting/pkg/e"
)

type Response struct {
	C *gin.Context
}

type ResponseInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Response) Send(httpCode, errCode int, data interface{}) {
	r.C.JSON(httpCode, ResponseInfo{
		Code: errCode,
		Msg:  e.Msg(errCode),
		Data: data,
	})
	return
}

func (r *Response) SendSucc(data interface{}) {
	r.C.JSON(200, ResponseInfo{
		Code: 1,
		Msg:  "success",
		Data: data,
	})
	return
}

func (r *Response) SendErr(code int, msg string) {
	if msg == "" {
		msg = e.Msg(code)
	} 
	r.C.JSON(200, ResponseInfo{
		Code: 0,
		Msg:  msg,
		Data: nil,
	})
	return
}