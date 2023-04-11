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
	logCtx := applogs.GenCtxFromGin(ctx)
	method := ctx.Request.Method
	applogs.Ctx(logCtx).Infof("method:%s", method)

	id := ctx.Query("id")
	name := ctx.Query("name")
	age := ctx.Query("age")
	applogs.Ctx(logCtx).Infof("id:%v", id)
	applogs.Ctx(logCtx).Infof("name:%v", name)
	applogs.Ctx(logCtx).Infof("age:%v", age)

	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Post(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	method := ctx.Request.Method
	applogs.Ctx(logCtx).Infof("method:%s", method)

	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	applogs.Ctx(logCtx).Infof("id:%v", id)
	applogs.Ctx(logCtx).Infof("name:%v", name)

	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Put(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	method := ctx.Request.Method
	applogs.Ctx(logCtx).Infof("method:%s", method)
	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) Delete(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	method := ctx.Request.Method
	applogs.Ctx(logCtx).Infof("method:%s", method)
	apphelper.EchoSuccess(ctx, method+" request")
}

func (*RestfulController) File(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	method := ctx.Request.Method
	applogs.Ctx(logCtx).Infof("method:%s", method)

	file, err := ctx.FormFile("file")
	if err != nil {
		applogs.Ctx(logCtx).Errorf("read file error:%v", err)
		apphelper.EchoFailed(ctx, enum.RespReadFileError, "")
		return
	}

	filename := file.Filename
	applogs.Ctx(logCtx).Infof("%v, %v", filename, err)

	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		applogs.Ctx(logCtx).Errorf("read file error:%v", err)
		apphelper.EchoFailed(ctx, enum.RespSystemError, err.Error())
		return
	}

	applogs.Ctx(logCtx).Infof("read file success")
	apphelper.EchoSuccess(ctx, method+" request")
}
