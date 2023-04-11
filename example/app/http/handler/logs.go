package handler

import (
	"elf-go/components/apphelper"
	"elf-go/components/applogs"
	"elf-go/example/app/http/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type LogsController struct{}

func (l *LogsController) Slow(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	applogs.Ctx(logCtx).Infof("SlowSlowSlowSlowSlow1")
	time.Sleep(time.Second * 5)
	applogs.Ctx(logCtx).Infof("SlowSlowSlowSlowSlow2")

	var resp response.Data
	resp.Code = 200
	resp.Msg = "Slow"
	resp.Data = "Slow data"

	apphelper.EchoSuccess(ctx, resp)
}

func (l *LogsController) Hook(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	logrus.WithContext(ctx).Infof("logrus 原生")
	applogs.Infof("applogs部分集成")
	applogs.Ctx(logCtx).Debugf("ctx集成")
	applogs.Ctx(logCtx).Infof("ctx集成")
	applogs.Ctx(logCtx).Warnf("ctx集成")
	applogs.Ctx(logCtx).Errorf("ctx集成")

	var resp response.Data
	resp.Code = 200
	resp.Msg = "hook"
	resp.Data = "hook data"

	apphelper.EchoSuccess(ctx, resp)
}
