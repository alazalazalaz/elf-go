package handler

import (
	"elf-go/app"
	"elf-go/components/logs"
	"elf-go/example/app/dao/entity"
	"elf-go/example/app/http/response"
	"github.com/gin-gonic/gin"
	"time"
)

func Info(ctx *gin.Context) {
	logs.Infof("im info1")
	time.Sleep(time.Second * 5)
	logs.Infof("im info2")
	ctx.JSON(200, "info")
}

func GetUserInfo(ctx *gin.Context) {
	logs.Info("get user info")
	var u entity.User
	app.Mysql().Preload("Article").Find(&u)
	logs.Info("result: ", logs.Content{"article:": u})
	ctx.JSON(200, u)
}

func Version(ctx *gin.Context) {
	logs.Infof("version")
	ver := struct {
		Version string
	}{Version: "0.0.1"}
	ctx.JSON(200, ver)
}

func Update(ctx *gin.Context) {
	logs.Warningf("update")

	success := response.Success{
		Code: 200,
		Msg:  "success",
	}
	ctx.JSON(200, success)
}

func Panic(ctx *gin.Context) {
	logs.Warningf("即将panic")
	panic("手动panic")
}
