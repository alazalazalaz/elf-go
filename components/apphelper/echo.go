package apphelper

import (
	"elf-go/components/appconsts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorCode int

type responseData struct {
	Code    ErrorCode   `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id"`
}

const (
	StatusUnauthorized ErrorCode = 401
)

func EchoSuccess(ctx *gin.Context, data interface{}) {
	resp := responseData{
		Code:    http.StatusOK,
		Msg:     "success",
		Data:    data,
		TraceId: ctx.Request.Header.Get(appconsts.HeaderTraceId),
	}
	ctx.JSON(http.StatusOK, resp)
}

func EchoFailed(ctx *gin.Context, errorCode ErrorCode, errorMsg string) {
	resp := responseData{
		Code:    errorCode,
		Msg:     errorMsg,
		Data:    nil,
		TraceId: ctx.Request.Header.Get(appconsts.HeaderTraceId),
	}
	ctx.JSON(http.StatusBadRequest, resp)
}

func EchoData(ctx *gin.Context, httpCode int, errorCode ErrorCode, errorMsg string, data interface{}) {
	resp := responseData{
		Code:    errorCode,
		Msg:     errorMsg,
		Data:    data,
		TraceId: ctx.Request.Header.Get(appconsts.HeaderTraceId),
	}
	ctx.JSON(httpCode, resp)
}
