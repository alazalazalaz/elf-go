package main

import (
	"elf-go"
	"elf-go/app"
	"elf-go/components/applogs"
	"elf-go/example/app/dao/plugin"
	"elf-go/example/app/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	configFile := "./config/conf.yaml"
	//初始化框架
	if err := elf_go.Init(configFile); err != nil {
		applogs.Errorf("framework init failed:%v", err)
		panic("")
	}

	//初始化redis
	if err := app.Redis().Init(); err != nil {
		applogs.Error(err.Error(), nil)
	}

	//初始化mysql
	applogs.Info("mysql before init", applogs.Content{"mysql:": app.Mysql().DB})
	if err := app.Mysql().Init(); err != nil {
		applogs.Error(err.Error(), nil)
	}
	applogs.Info("mysql after init", applogs.Content{"mysql:": app.Mysql().DB})

	// 初始化mysql的Hook
	if err := app.Mysql().Use(&plugin.BeforeAfterPlugin{}); err != nil {
		applogs.Errorf("init gorm plugin failed, err:%v", err)
	}

	//启动服务
	engine := gin.New()

	route.InitRoute(engine)

	port := app.Config().GetSysConfig().ListenPort
	applogs.Info(fmt.Sprintf("HttpServer Listen At:%d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), engine); err != nil {
		applogs.Error(err.Error())
	}

}
