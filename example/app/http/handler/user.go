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
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Info(ctx *gin.Context) {
	var resp response.Data
	resp.Code = 200
	resp.Msg = "info api"
	resp.Data = "info data"

	apphelper.EchoSuccess(ctx, resp)
}

func Select(ctx *gin.Context) {
	applogs.Ctx(ctx).Infof("get user info")
	var u entity.User
	db := app.Mysql()
	db.Where("id = ?", 1).Find(&u)
	applogs.Ctx(ctx).Infof("result: %v", u)

	apphelper.EchoSuccess(ctx, u)
}

func Type(ctx *gin.Context) {
	username := ctx.GetString("username")
	//password := ctx.GetString("password")

	err := checkUsername(username)
	if err != nil {
		applogs.Ctx(ctx).Errorf("check username error: %v", err)
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
	var u entity.User
	db := app.Mysql()
	db.Where("id = ?", 5).Find(&u)
	time.Sleep(time.Second * 5)
	applogs.Ctx(ctx).Infof("result: %v", u)

	apphelper.EchoSuccess(ctx, u)
}

func Version(ctx *gin.Context) {
	applogs.Ctx(ctx).Infof("version")
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
	//update 只更新当前字段
	app.Mysql().Model(&entity.User{}).Where("id = ?", 1).Update("updated_at", time.Now().Unix())

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
	applogs.Ctx(ctx).Warnf("即将panic")
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
	go func() {
		for i := 10; i > 0; i-- {
			applogs.Ctx(ctx).Infof("loop data i :%v", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	apphelper.EchoSuccess(ctx, "")
}

func Auth(ctx *gin.Context) {
	applogs.Ctx(ctx).Infof("auth SUCCESS action")
	apphelper.EchoSuccess(ctx, nil)
}
