package main

import (
	"elf-go/app"
	"elf-go/components/logs"
	"os"
)

//@todolist
//1.新增mysql组件
//2.新增redis集群
func main(){
	//初始化配置文件
	app.Config().SetConfigFilePath("conf.yml")
	if err:= app.Config().Init(); err != nil{
		logs.Error(err.Error(), nil)
	}

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

	os.Exit(0)
}

