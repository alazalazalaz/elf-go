package handler

import (
	"elf-go/components/logs"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context){
	logs.Info("get user info")

}

func Version(ctx *gin.Context){
	logs.Info("version")
}

func Update(ctx *gin.Context){
	logs.Warning("update")
}