package applogs

import (
	"context"
	"elf-go/components/appconsts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	l.entry.Debugf("["+traceId+"] "+format, args...)
}

func (l *logsWithCtx) Infof(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	l.entry.Infof("["+traceId+"] "+format, args...)
}

func (l *logsWithCtx) Warnf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	l.entry.Warnf("["+traceId+"] "+format, args...)
}

func (l *logsWithCtx) Errorf(format string, args ...interface{}) {
	traceId, _ := l.getTraceInfoFromContext()
	l.entry.Errorf("["+traceId+"] "+format, args...)
}

func (l *logsWithCtx) getTraceInfoFromContext() (string, string) {
	traceId, spanId := "", ""
	con, ok := l.entry.Context.(*gin.Context)
	if ok {
		traceId = con.Request.Header.Get(appconsts.HeaderTraceId)
	}

	return traceId, spanId
}
