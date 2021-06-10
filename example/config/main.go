package main

import (
	"elf-go/components/config"
	"fmt"
)

func main(){
	path := "conf.yaml"
	conf := config.InitConfig(path)
	fmt.Println("读取配置文件：", conf.Redis.Ip, conf.Redis.Port)
	fmt.Println(config.GetRedisConfig())
}
