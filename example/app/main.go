package main

import (
	"elf-go/app"
	"elf-go/components/logs"
	"elf-go/framework"
	"os"
)

func main(){
	//初始化框架
	framework.Init("conf.yml")

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

