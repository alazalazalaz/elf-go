package app

import (
	"elf-go/components/config"
	"elf-go/components/logs"
	elfredis "elf-go/components/redis"
	"go.uber.org/dig"
)

var container = dig.New()
var elfConf *config.Config
var elfRedis *elfredis.Redis

func init(){

	if err := container.Provide(config.New); err != nil {
		logs.Error("初始化配置文件失败：", logs.Content{"err": err})
	}

	if err := container.Provide(elfredis.New); err != nil {
		logs.Error("初始化redis失败：", logs.Content{"err": err})
	}
}

func Config() *config.Config{
	if elfConf == nil{
		_ = container.Invoke(func(conf *config.Config) {
			elfConf = conf
			return
		})
	}
	return elfConf
}

func Redis() *elfredis.Redis{
	if elfRedis == nil{
		_ = container.Invoke(func(redisClient *elfredis.Redis){
			elfRedis = redisClient
			return
		})
	}
	return elfRedis
}
