package applogs

import (
	"context"
	"elf-go/components/appconsts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type logsWithCtx struct {
	entry *logrus.Entry
}

func Ctx(ctx context.Context) *logsWithCtx {
	en := logrus.WithContext(ctx)
	return &logsWithCtx{
		entry: en,
	}
}

func (l *logsWithCtx) Debugf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Debugf(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logsWithCtx) Infof(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Infof(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logsWithCtx) Warnf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Warnf(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logsWithCtx) Errorf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	callerFile, callerLineNum := l.getCallerInfo(2)

	l.entry.Errorf(fmt.Sprintf("[%s:%d][%s] %s", callerFile, callerLineNum, traceId, format), args...)
}

func (l *logsWithCtx) getTraceInfoFromContext() (string, string) {
	traceId, spanId := "", ""
	con, ok := l.entry.Context.(*gin.Context)
	if ok {
		traceId = con.Request.Header.Get(appconsts.HeaderTraceId)
	}

	return traceId, spanId
}

func (l *logsWithCtx) getCallerInfo(skip int) (string, int) {
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
