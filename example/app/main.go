package main

import (
	"elf-go/app"
	"elf-go/components/logs"
	"elf-go/example/app/dao/plugin"
	"elf-go/example/app/route"
	"elf-go/framework"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	configFile := "./config/conf.yml"
	//初始化框架
	if err := framework.Init(configFile); err != nil {
		logs.Errorf("framework init failed:%v", err)
		panic("")
	}

	//初始化redis
	if err := app.Redis().Init(); err != nil {
		logs.Error(err.Error(), nil)
	}

	//初始化mysql
	logs.Info("mysql before init", logs.Content{"mysql:": app.Mysql().DB})
	if err := app.Mysql().Init(); err != nil {
		logs.Error(err.Error(), nil)
	}
	logs.Info("mysql after init", logs.Content{"mysql:": app.Mysql().DB})

	// 初始化mysql的Hook
	if err := app.Mysql().Use(&plugin.BeforeAfterPlugin{}); err != nil {
		logs.Errorf("init gorm plugin failed, err:%v", err)
	}

	//启动服务
	engine := gin.New()

	route.InitRoute(engine)

	port := app.Config().GetSysConfig().ListenPort
	logs.Info(fmt.Sprintf("HttpServer Listen At:%d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), engine); err != nil {
		logs.Error(err.Error())
	}

}
