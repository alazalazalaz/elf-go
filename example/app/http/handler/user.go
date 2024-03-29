package handler

import (
	"elf-go/app"
	"elf-go/components/apphelper"
	"elf-go/components/appjwts"
	"elf-go/components/applogs"
	"elf-go/example/app/dao/entity"
	"elf-go/example/app/enum"
	"elf-go/example/app/http/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Info(ctx *gin.Context) {
	var resp response.Data
	resp.Code = 200
	resp.Msg = "info api"
	resp.Data = "info data"

	var aaa *applogs.Context
	if aaa == nil {
		fmt.Println("aaa is nil")
	} else {
		fmt.Println(aaa.TraceId)
	}

	apphelper.EchoSuccess(ctx, resp)
}

func Select(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	applogs.Ctx(logCtx).Infof("get user info")
	var u entity.User
	db := app.Mysql()
	db.Where("id = ?", 1).Find(&u)
	applogs.Ctx(logCtx).Infof("before: %v", u)
	go func() {
		time.Sleep(time.Second * 2)
		newCtx := applogs.SpanCtx(logCtx)
		applogs.Ctx(newCtx).Infof("goroutine child1: %v", u)
		time.Sleep(time.Second * 2)

		newCtx2 := applogs.SpanCtx(newCtx)
		applogs.Ctx(newCtx2).Infof("goroutine child1-1: %v", u)
	}()
	go func() {

		newCtx3 := applogs.SpanCtx(logCtx)
		applogs.Ctx(newCtx3).Infof("goroutine child2: %v", u)
		time.Sleep(time.Second * 2)
		newCtx4 := applogs.SpanCtx(newCtx3)
		applogs.Ctx(newCtx4).Infof("goroutine child2-1: %v", u)
	}()

	applogs.Ctx(logCtx).Infof("result: %v", u)
	applogs.Ctx(logCtx).Infof("result: %v", u)
	applogs.Ctx(logCtx).Infof("result: %v", u)
	apphelper.EchoSuccess(ctx, u)
}

func Type(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	username := ctx.GetString("username")
	//password := ctx.GetString("password")

	err := checkUsername(username)
	if err != nil {
		applogs.Ctx(logCtx).Errorf("check username error: %v", err)
		apphelper.EchoFailed(ctx, 1001, err.Error())
		return
	}

}

func checkUsername(username string) error {
	usernameArr := strings.Split(username, "|")
	if len(usernameArr) != 2 {
		return errors.New("username error")
	}

	return nil
}

func SlowQuery(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	var u entity.User
	db := app.Mysql()
	db.Where("id = ?", 5).Find(&u)
	time.Sleep(time.Second * 5)
	applogs.Ctx(logCtx).Infof("result: %v", u)

	apphelper.EchoSuccess(ctx, u)
}

func Version(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	applogs.Ctx(logCtx).Infof("version")
	ver := struct {
		Version string
	}{Version: "0.0.1"}

	apphelper.EchoSuccess(ctx, ver)
}

func Create(ctx *gin.Context) {
	var obj entity.User
	obj.Username = "小张"
	obj.CreatedAt = time.Now().Unix()
	obj.UpdatedAt = obj.CreatedAt

	re := app.Mysql().Create(&obj)
	if re.Error != nil {
		apphelper.EchoFailed(ctx, enum.RespDbError, re.Error.Error())
		return
	}

	apphelper.EchoSuccess(ctx, nil)
}

func Update(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	//update 只更新当前字段
	re := app.Mysql().Model(&entity.User{}).Where("id = ?", 1).Update("updatexxd_at", time.Now().Unix())
	if re.Error != nil {
		applogs.Ctx(logCtx).Errorf("update error: %v", re.Error)
		apphelper.EchoFailed(ctx, enum.RespDbError, re.Error.Error())
		return
	}
	applogs.Ctx(logCtx).Infof("update success, row:%v", re.RowsAffected)

	apphelper.EchoSuccess(ctx, nil)
}

func Save(ctx *gin.Context) {
	var obj entity.User
	app.Mysql().Where("id = ?", 1).Find(&obj)

	obj.UpdatedAt = time.Now().Unix()
	app.Mysql().Save(obj) //save会update所有字段

	apphelper.EchoSuccess(ctx, nil)
}

func Panic(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	applogs.Ctx(logCtx).Warnf("即将panic")
	panic("手动panic")
}

func Login(ctx *gin.Context) {
	token, err := appjwts.CreateJwtToken(10)
	if err != nil {
		apphelper.EchoFailed(ctx, enum.RespGenJWTError, err.Error())
		return
	}

	apphelper.EchoSuccess(ctx, token)
}

func Sleep(ctx *gin.Context) {
	time.Sleep(9 * time.Second)

	apphelper.EchoSuccess(ctx, "")
}

func Loop(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	go func() {
		for i := 10; i > 0; i-- {
			applogs.Ctx(logCtx).Infof("loop data i :%v", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	apphelper.EchoSuccess(ctx, "")
}

func Auth(ctx *gin.Context) {
	logCtx := applogs.GenCtxFromGin(ctx)
	applogs.Ctx(logCtx).Infof("auth SUCCESS action")
	apphelper.EchoSuccess(ctx, nil)
}
