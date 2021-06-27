package main

import (
	"elf-go/app"
	"elf-go/components/logs"
)

func main(){
	path := "conf.yml"
	app.Config().SetConfigFilePath(path)
	if err := app.Config().Init(); err != nil {
		logs.Error("init config failed", logs.Content{"con": err})
	}

	logs.Info("config:", logs.Content{"config:": app.Config().C})
	if app.Config().GetSysConfig().Debug == true{
		logs.Info("debug==true")
	}else{
		logs.Info("debug==false")
	}
}
