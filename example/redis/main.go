package main

import (
	"elf-go/components/config"
	"elf-go/components/logs"
	"elf-go/components/redis"
)

func main(){
	path := "conf.yaml"
	if err := config.InitConfig(path); err != nil{
		logs.Error("init config failed:", logs.Content{"err": err})
	}

	redisConfig := config.GetRedisConfig()

	if _, err := redis.NewRedis(redisConfig); err != nil{
		logs.Error("init redis failed:", logs.Content{"err": err})
	}

}
