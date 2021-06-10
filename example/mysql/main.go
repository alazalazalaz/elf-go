package main

import (
	"elf-go/components/config"
	"elf-go/components/logs"
	"elf-go/components/mysql"
)

func main(){
	path := "conf.yaml"
	if err := config.NewConfig(path); err != nil{
		logs.Error("init config failed:", logs.Content{"err": err})
	}

	mysqlConfig := config.GetMysqlConfig()

	if _, err := mysql.NewMysql(mysqlConfig); err != nil{
		logs.Error("init mysql failed:", logs.Content{"err": err})
	}

}
