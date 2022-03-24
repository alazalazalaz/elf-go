package handler

import (
	"elf-go/app"
	"elf-go/components/jwts"
	"elf-go/components/logs"
	"elf-go/example/app/dao/entity"
	"elf-go/example/app/enum"
	"elf-go/example/app/http/response"
	"elf-go/utils/helper"
	"github.com/gin-gonic/gin"
	"time"
)

func Info(ctx *gin.Context) {
	var resp response.Data
	resp.Code = 200
	resp.Msg = "info api"
	resp.Data = "info data"

	helper.EchoSuccess(ctx, resp)
}

func Select(ctx *gin.Context) {
	logs.Info("get user info")
	var u entity.User
	db := app.Mysql()
	db.Where("id = ?", 1).Find(&u)
	logs.Infof("result: %v", u)

	helper.EchoSuccess(ctx, u)
}

func SlowQuery(ctx *gin.Context) {
	var u entity.User
	db := app.Mysql()
	db.Where("id = ?", 5).Find(&u)
	time.Sleep(time.Second * 5)
	logs.Infof("result: %v", u)

	helper.EchoSuccess(ctx, u)

}

func Version(ctx *gin.Context) {
	logs.Infof("version")
	ver := struct {
		Version string
	}{Version: "0.0.1"}

	helper.EchoSuccess(ctx, ver)
}

func Create(ctx *gin.Context) {
	var obj entity.User
	obj.Username = "小张"
	obj.CreatedAt = time.Now().Unix()
	obj.UpdatedAt = obj.CreatedAt

	re := app.Mysql().Create(&obj)
	if re.Error != nil {
		helper.EchoFailed(ctx, enum.RespDbError, re.Error.Error())
		return
	}

	helper.EchoSuccess(ctx, nil)
}

func Update(ctx *gin.Context) {
	//update 只更新当前字段
	app.Mysql().Model(&entity.User{}).Where("id = ?", 1).Update("updated_at", time.Now().Unix())

	helper.EchoSuccess(ctx, nil)
}

func Save(ctx *gin.Context) {
	var obj entity.User
	app.Mysql().Where("id = ?", 1).Find(&obj)

	obj.UpdatedAt = time.Now().Unix()
	app.Mysql().Save(obj) //save会update所有字段

	helper.EchoSuccess(ctx, nil)
}

func Panic(ctx *gin.Context) {
	logs.Warningf("即将panic")
	panic("手动panic")
}

func Login(ctx *gin.Context) {
	token, err := jwts.CreateJwtToken(10)
	if err != nil {
		helper.EchoFailed(ctx, enum.RespGenJWTError, err.Error())
		return
	}

	helper.EchoSuccess(ctx, token)
}

func Sleep(ctx *gin.Context) {
	time.Sleep(9 * time.Second)

	helper.EchoSuccess(ctx, "")
}

func Loop(ctx *gin.Context) {
	go func() {
		for i := 10; i > 0; i-- {
			logs.Infof("loop data i :%v", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	helper.EchoSuccess(ctx, "")
}

func Auth(ctx *gin.Context) {
	logs.Infof("auth SUCCESS action")
	helper.EchoSuccess(ctx, nil)
}
