package applogs

import (
	"context"
	"elf-go/components/appconsts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type logWrapper struct {
	Ctx   Context
	entry *logrus.Entry
}

type Context struct {
	TraceId string
	SpanIds []string
	context.Context
}

// GenCtxFromGin 用于生成一个有gin.Context的Context
func GenCtxFromGin(ctx *gin.Context) Context {
	traceId := ctx.Request.Header.Get(appconsts.HeaderTraceId)
	var spanIds []string
	return Context{
		TraceId: traceId,
		SpanIds: spanIds,
	}
}

// GenCtxFromNoneGin 用于生成一个没有gin.Context的Context，比如定时任务
func GenCtxFromNoneGin() Context {
	return Context{
		TraceId: uuid.New().String(),
		SpanIds: []string{},
	}
}

// SpanCtx 生成一个新的Context，用于生成一个新的Span
func SpanCtx(ctx Context) Context {
	newCtx := Context{
		TraceId: ctx.TraceId,
		SpanIds: ctx.SpanIds,
	}
	spanId := uuid.New().String()[24:]
	newCtx.SpanIds = append(newCtx.SpanIds, spanId)
	return newCtx
}

func Ctx(ctx Context) *logWrapper {
	return &logWrapper{
		Ctx:   ctx,
		entry: defaultEntity,
	}
}

func (l *logWrapper) Debugf(format string, args ...interface{}) {
	traceId, spanIds := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Debugf(fmt.Sprintf("[%s:%d][%s][%s] %s", callerFile, callerLineNum, traceId, spanIds, format), args...)
}

func (l *logWrapper) Infof(format string, args ...interface{}) {
	traceId, spanIds := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Infof(fmt.Sprintf("[%s:%d][%s][%s] %s", callerFile, callerLineNum, traceId, spanIds, format), args...)
}

func (l *logWrapper) Warnf(format string, args ...interface{}) {
	traceId, spanIds := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Warnf(fmt.Sprintf("[%s:%d][%s][%s] %s", callerFile, callerLineNum, traceId, spanIds, format), args...)
}

func (l *logWrapper) Errorf(format string, args ...interface{}) {
	traceId, spanIds := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Errorf(fmt.Sprintf("[%s:%d][%s][%s] %s", callerFile, callerLineNum, traceId, spanIds, format), args...)
}

func (l *logWrapper) getTraceInfoFromContext() (string, string) {
	traceId, spanIds := l.Ctx.TraceId, ""
	spanIds = strings.Join(l.Ctx.SpanIds, "-")

	return traceId, spanIds
}

func (l *logWrapper) getCallerInfo(skip int) (string, int) {
	_, callerFile, callerLineNum, ok := runtime.Caller(skip)
	if ok {
		fileSplit := strings.Split(callerFile, "/")
		fileSplitLen := len(fileSplit)
		if fileSplitLen > 1 {
			callerFile = "..." + "/" + fileSplit[fileSplitLen-2] + "/" + fileSplit[fileSplitLen-1]
		}
	}

	return callerFile, callerLineNum
}
