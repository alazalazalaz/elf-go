package handler

import (
	"elf-go/components/apphelper"
	"elf-go/components/applogs"
	"elf-go/example/app/http/response"
	"github.com/gin-gonic/gin"
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
	applogs.Infof("path: logs/hook")

	applogs.Ctx(ctx).Infof("hahahha")

	var resp response.Data
	resp.Code = 200
	resp.Msg = "hook"
	resp.Data = "hook data"

	apphelper.EchoSuccess(ctx, resp)
}
