package handler

import (
	"elf-go/components/appconsts"
	"elf-go/components/apphelper"
	"elf-go/components/applogs"
	"elf-go/example/app/http/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogsController struct{}

func (l *LogsController) Hook(ctx *gin.Context) {
	applogs.Infof("path: logs/hook")

	logCtx := logrus.WithContext(ctx)
	con, ok := logCtx.Context.(*gin.Context)
	traceId := ""
	if ok {
		traceId = con.Request.Header.Get(appconsts.HeaderTraceId)
	}
	logCtx.Infof("[%s] xxx", traceId)

	var resp response.Data
	resp.Code = 200
	resp.Msg = "hook"
	resp.Data = "hook data"

	apphelper.EchoSuccess(ctx, resp)
}
