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
	applogs.Ctx(ctx).Infof("SlowSlowSlowSlowSlow1")
	time.Sleep(time.Second * 5)
	applogs.Ctx(ctx).Infof("SlowSlowSlowSlowSlow2")

	var resp response.Data
	resp.Code = 200
	resp.Msg = "Slow"
	resp.Data = "Slow data"

	apphelper.EchoSuccess(ctx, resp)
}

func (l *LogsController) Hook(ctx *gin.Context) {
	logrus.WithContext(ctx).Infof("logrus 原生")
	applogs.Infof("applogs部分集成")
	applogs.Ctx(ctx).Debugf("ctx集成")
	applogs.Ctx(ctx).Infof("ctx集成")
	applogs.Ctx(ctx).Warnf("ctx集成")
	applogs.Ctx(ctx).Errorf("ctx集成")

	var resp response.Data
	resp.Code = 200
	resp.Msg = "hook"
	resp.Data = "hook data"

	apphelper.EchoSuccess(ctx, resp)
}
