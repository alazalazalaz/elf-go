package middleware

import (
	"elf-go/components/applogs"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//@todo 可以修改PrintReqAndResp为一个方法，把要打印的参数都传递进去。
func PrintReqAndResp(ctx *gin.Context) {
	startAt := time.Now()

	headerString := ""
	for headerKey, headerValue := range ctx.Request.Header {
		headerString += fmt.Sprintf("%s:%s; ", headerKey, strings.Join(headerValue, ","))
	}
	if err := ctx.Request.ParseForm(); err != nil {
		applogs.Ctx(ctx).Errorf("PrintReqAndResp=>ParseForm error:%v", err)
	}

	//bodyString := ""
	//bodyBytes, err := io.ReadAll(ctx.Request.Body)//注意这里读取了body之后，ctx.Request.Body就为空了哟。
	//if err != nil {
	//	logs.Errorf("Read Body Error err:%v", err)
	//} else {
	//	bodyString = string(bodyBytes)
	//}

	applogs.Ctx(ctx).Infof(`[begin]=>Remote Address:%s | Request Method:%s | Request URI:"%s" | Request Headers:%s | Form Data:%s`,
		ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.RequestURI, headerString, ctx.Request.PostForm.Encode())

	ctx.Next()

	applogs.Ctx(ctx).Infof("[end]=> duration:%s", time.Since(startAt))
}
