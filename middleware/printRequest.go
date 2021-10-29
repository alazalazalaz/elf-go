package middleware

import (
	"elf-go/components/logs"
	"elf-go/utils/traceid"
	"github.com/gin-gonic/gin"
)

//@todo 可以修改PrintReqAndResp为一个方法，把要打印的参数都传递进去。
func PrintReqAndResp(ctx *gin.Context) {
	traceid.GenTraceId()
	logs.Infof(`[begin]:%s | %s | "%s" |`, ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.RequestURI)

	ctx.Next()

	logs.Infof("[end]")
}
