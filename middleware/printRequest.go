package middleware

import (
	"bytes"
	"elf-go/components/applogs"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type responseBody struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBody) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

//@todo 可以修改PrintReqAndResp为一个方法，把要打印的参数都传递进去。
func PrintReqAndResp(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	startAt := time.Now()

	headerString := ""
	for headerKey, headerValue := range ctx.Request.Header {
		headerString += fmt.Sprintf("%s:%s; ", headerKey, strings.Join(headerValue, ","))
	}
	if err := ctx.Request.ParseForm(); err != nil {
		applogs.Ctx(logCtx).Errorf("PrintReqAndResp=>ParseForm error:%v", err)
	}

	//bodyString := ""
	//bodyBytes, err := io.ReadAll(ctx.Request.Body)//注意这里读取了body之后，ctx.Request.Body就为空了哟。
	//if err != nil {
	//	logs.Errorf("Read Body Error err:%v", err)
	//} else {
	//	bodyString = string(bodyBytes)
	//}

	applogs.Ctx(logCtx).Infof(`[begin]=>Remote Address:%s | Request Method:%s | Request URI:"%s" | Request Headers:%s | Form Data:%s`,
		ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.RequestURI, headerString, ctx.Request.PostForm.Encode())

	// 放置记录response的writer
	bodyWriter := &responseBody{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = bodyWriter

	ctx.Next()

	applogs.Ctx(logCtx).Infof("[end]=> duration:%s, statusCode:%v, response:%v", time.Since(startAt), ctx.Writer.Status(), bodyWriter.body.String())
}
