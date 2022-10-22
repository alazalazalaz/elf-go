package middleware

import (
	"elf-go/components/appconsts"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceId(ctx *gin.Context) {
	traceId := ctx.Request.Header.Get(appconsts.HeaderTraceId)
	if traceId == "" {
		traceId = uuid.New().String()
		ctx.Request.Header.Set(appconsts.HeaderTraceId, traceId)
	}

	//这个logrus的hook不行，有并发问题，还是需要把ctx挂到logger下面，或者使用logger()方法来打印，实现一个中间件把ctx塞到logger()里面。
	//hook := &applogs.TraceIdHook{
	//	Ctx: ctx,
	//}
	//logrus.AddHook(hook)

	ctx.Next()
}
