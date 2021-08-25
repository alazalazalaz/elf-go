package handler

import (
	"elf-go/app"
	"elf-go/components/logs"
	"elf-go/example/app/dao/entity"
	"github.com/gin-gonic/gin"
)


func GetUserInfo(ctx *gin.Context){
	logs.Info("get user info")
	var u entity.User
	app.Mysql().Preload("Article").Find(&u)
	logs.Info("result: ", logs.Content{"article:": u})
	ctx.JSON(200, u)
}

func Version(ctx *gin.Context){
	logs.Info("version")
}

func Update(ctx *gin.Context){
	logs.Warning("update")
}