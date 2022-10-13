package handler

import (
	"elf-go/components/apphelper"
	"elf-go/components/applogs"
	"elf-go/example/app/enum"
	"github.com/gin-gonic/gin"
)

type RestfulController struct {
}

func (*RestfulController) Get(ctx *gin.Context) {
	method := ctx.Request.Method
	applogs.Infof("method:%s", method)

	id := ctx.Query("id")
	name := ctx.Query("name")
	age := ctx.Query("age")
	applogs.Infof("id:%v", id)
	applogs.Infof("name:%v", name)
	applogs.Infof("age:%v", age)

	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Post(ctx *gin.Context) {
	method := ctx.Request.Method
	applogs.Infof("method:%s", method)

	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	applogs.Infof("id:%v", id)
	applogs.Infof("name:%v", name)

	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Put(ctx *gin.Context) {
	method := ctx.Request.Method
	applogs.Infof("method:%s", method)
	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Delete(ctx *gin.Context) {
	method := ctx.Request.Method
	applogs.Infof("method:%s", method)
	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) File(ctx *gin.Context) {
	method := ctx.Request.Method
	applogs.Infof("method:%s", method)

	file, err := ctx.FormFile("file")
	if err != nil {
		applogs.Errorf("read file error:%v", err)
		apphelper.EchoFailed(ctx, enum.RespReadFileError, "")
		return
	}

	filename := file.Filename
	applogs.Infof("%v, %v", filename, err)

	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		applogs.Errorf("read file error:%v", err)
		apphelper.EchoFailed(ctx, enum.RespSystemError, err.Error())
		return
	}

	applogs.Infof("read file success")
	apphelper.EchoSuccess(ctx, method+" request")
}
