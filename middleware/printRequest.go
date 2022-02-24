package middleware

import (
	"elf-go/components/logs"
	"elf-go/utils/traceid"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"time"
)

//@todo 可以修改PrintReqAndResp为一个方法，把要打印的参数都传递进去。
func PrintReqAndResp(ctx *gin.Context) {
	startAt := time.Now()
	traceid.GenTraceId()

	headerString := ""
	for headerKey, headerValue := range ctx.Request.Header {
		headerString += fmt.Sprintf("%s:%s;", headerKey, strings.Join(headerValue, ","))
	}
	if err := ctx.Request.ParseForm(); err != nil {
		logs.Errorf("PrintReqAndResp=>ParseForm error:%v", err)
	}

	bodyString := ""
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logs.Errorf("Read Body Error err:%v", err)
	} else {
		bodyString = string(bodyBytes)
	}

	logs.Infof(`[begin]=>Remote Address:%s | Request Method:%s | Request URI:"%s" | Request Headers:%s | Form Data:%s | Body:%s`,
		ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.RequestURI, headerString, ctx.Request.PostForm.Encode(), bodyString)

	ctx.Next()

	logs.Infof("[end]=> duration:%s", time.Since(startAt))
}
