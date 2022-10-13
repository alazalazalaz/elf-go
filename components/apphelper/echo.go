package apphelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorCode int

type responseData struct {
	Code ErrorCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	StatusUnauthorized ErrorCode = 401
)

func EchoSuccess(c *gin.Context, data interface{}) {
	resp := responseData{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func EchoFailed(c *gin.Context, errorCode ErrorCode, errorMsg string) {
	resp := responseData{
		Code: errorCode,
		Msg:  errorMsg,
		Data: nil,
	}
	c.JSON(http.StatusBadRequest, resp)
}

func EchoData(c *gin.Context, httpCode int, errorCode ErrorCode, errorMsg string, data interface{}) {
	resp := responseData{
		Code: errorCode,
		Msg:  errorMsg,
		Data: data,
	}
	c.JSON(httpCode, resp)
}
