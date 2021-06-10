package main

import (
	"elf-go/app"
	"elf-go/components/logs"
	"os"
)

func main(){
	app.Config().SetConfigFilePath("conf.yaml")
	if err:= app.Config().Init(); err != nil{
		logs.Error(err.Error(), nil)
	}

	if err := app.Redis().Init(); err != nil{
		logs.Error(err.Error(), nil)
	}

	os.Exit(0)
}

