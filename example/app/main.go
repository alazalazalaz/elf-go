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
	//初始化框架，必须步骤
	if err := elf_go.Init(configFile); err != nil {
		applogs.Errorf("framework init failed:%v", err)
		panic("")
	}

	//初始化redis，可选步骤，不影响redis使用
	if err := app.Redis().TestConnection(); err != nil {
		applogs.Errorf("redis TestConnection failed, err:%v", err)
	}

	//初始化mysql，可选步骤，不影响mysql使用
	if err := app.Mysql().TestConnection(); err != nil {
		applogs.Errorf("mysql TestConnection failed, err:%v", err)
	}

	// 初始化mysql的Hook，可选步骤，可加载0-多个插件
	if err := app.Mysql().Use(&plugin.BeforeAfterPlugin{}); err != nil {
		applogs.Errorf("init gorm plugin failed, err:%v", err)
	}

	//启动gin服务，可选步骤，如果是cronjob可不使用
	engine := gin.New()

	//注册路由，可选步骤，如果是cronjob可不使用
	route.InitRoute(engine)

	//启动web服务并加载gin，可选步骤，如果是cronjob可不使用
	port := app.Config().GetSysConfig().ListenPort
	applogs.Infof("HttpServer Listen At:%v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), engine); err != nil {
		applogs.Errorf("http.ListenAndServe failed, port:%v, err:%v", port, err)
		panic("panic: start gin server error")
	}

}
