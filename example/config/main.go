package main

import (
	"elf-go/app"
)

func main() {
	path := "conf.yml"
	app.Config().SetConfigFilePath(path)
	if err := app.Config().Init(); err != nil {
		applogs.Error("init config failed", applogs.Content{"con": err})
	}

	applogs.Info("config:", applogs.Content{"config:": app.Config().C})
	if app.Config().GetSysConfig().Debug == true {
		applogs.Info("debug==true")
	} else {
		applogs.Info("debug==false")
	}
}
