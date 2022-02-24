package handler

import (
	"elf-go/components/logs"
	"elf-go/utils/helper"
	"github.com/gin-gonic/gin"
)

type RestfulController struct {
}

func (*RestfulController) Get(ctx *gin.Context) {
	method := ctx.Request.Method
	logs.Infof("method:%s", method)
	helper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Post(ctx *gin.Context) {
	method := ctx.Request.Method
	logs.Infof("method:%s", method)
	helper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Put(ctx *gin.Context) {
	method := ctx.Request.Method
	logs.Infof("method:%s", method)
	helper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Delete(ctx *gin.Context) {
	method := ctx.Request.Method
	logs.Infof("method:%s", method)
	helper.EchoSuccess(ctx, method+" request")
}
