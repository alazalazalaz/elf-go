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
	entry *logrus.Entry
}

type Context struct {
	TraceId string
	SpanIds []string
	context.Context
}

func GenCtxFromGin(ctx *gin.Context) Context {
	traceId := ctx.Request.Header.Get(appconsts.HeaderTraceId)
	var spanIds []string
	return Context{
		TraceId: traceId,
		SpanIds: spanIds,
	}
}

func GenCtxFromNoneGin() Context {
	return Context{
		TraceId: uuid.New().String(),
		SpanIds: []string{},
	}
}

func Ctx(ctx Context) *logWrapper {
	en := logrus.WithContext(ctx)
	return &logWrapper{
		entry: en,
	}
}

func (l *logWrapper) Debugf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Debugf(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logWrapper) Infof(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Infof(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logWrapper) Warnf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Warnf(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logWrapper) Errorf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Errorf(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logWrapper) getTraceInfoFromContext() (string, string) {
	traceId, spanId := "", ""
	con, ok := l.entry.Context.(*gin.Context)
	if ok {
		traceId = con.Request.Header.Get(appconsts.HeaderTraceId)
	}

	return traceId, spanId
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
