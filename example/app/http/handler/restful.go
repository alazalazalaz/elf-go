package handler

import (
	"elf-go/components/logs"
	"elf-go/example/app/enum"
	"elf-go/utils/helper"
	"github.com/gin-gonic/gin"
)

type RestfulController struct {
}

func (*RestfulController) Get(ctx *gin.Context) {
	method := ctx.Request.Method
	logs.Infof("method:%s", method)

	id := ctx.Query("id")
	name := ctx.Query("name")
	age := ctx.Query("age")
	logs.Infof("id:%v", id)
	logs.Infof("name:%v", name)
	logs.Infof("age:%v", age)

	helper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Post(ctx *gin.Context) {
	method := ctx.Request.Method
	logs.Infof("method:%s", method)

	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	logs.Infof("id:%v", id)
	logs.Infof("name:%v", name)

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

func (*RestfulController) File(ctx *gin.Context) {
	method := ctx.Request.Method
	logs.Infof("method:%s", method)

	file, err := ctx.FormFile("file")
	if err != nil {
		logs.Errorf("read file error:%v", err)
		helper.EchoFailed(ctx, enum.RespReadFileError, "")
		return
	}

	filename := file.Filename
	logs.Infof("%v, %v", filename, err)

	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		logs.Errorf("read file error:%v", err)
		helper.EchoFailed(ctx, enum.RespSystemError, err.Error())
		return
	}

	logs.Infof("read file success")
	helper.EchoSuccess(ctx, method+" request")
}
