package main

import (
	"elf-go/app"
	"elf-go/components/logs"
	"elf-go/example/app/route"
	"elf-go/framework"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	//初始化框架
	framework.Init("example/app/config/conf.yml")

	//初始化redis
	if err := app.Redis().Init(); err != nil{
		logs.Error(err.Error(), nil)
	}

	//初始化mysql
	logs.Info("mysql before init", logs.Content{"mysql:": app.Mysql().DB})
	if err := app.Mysql().Init(); err != nil{
		logs.Error(err.Error(), nil)
	}
	logs.Info("mysql after init", logs.Content{"mysql:": app.Mysql().DB})

	//启动服务
	engine := gin.New()

	route.InitRoute(engine)

	if err := http.ListenAndServe(":7070", engine); err != nil{
		logs.Error(err.Error())
	}
}

