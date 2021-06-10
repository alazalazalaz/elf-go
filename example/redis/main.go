package main

import (
	"elf-go/components/config"
	"elf-go/components/redis"
	goredis "github.com/go-redis/redis"
	"go.uber.org/dig"
	"time"
)

//没有dig的版本
//func main(){
//	path := "conf.yaml"
//	if err := config.NewConfig(path); err != nil{
//		logs.Error("init config failed:", logs.Content{"err": err})
//	}
//
//	redisConfig := config.GetRedisConfig()
//
//	if _, err := redis.NewRedis(redisConfig); err != nil{
//		logs.Error("init redis failed:", logs.Content{"err": err})
//	}
//
//}

//新增dig版本
func main(){
	container := dig.New()

	path := "conf.yaml"
	container.Provide(config.NewConfig(path))
	container.Provide(config.GetRedisConfig)
	container.Provide(redis.NewRedis)

	container.Invoke(func(redisClient *goredis.Client) {
		redisClient.Set("test", 1233333, time.Second * 50)
	})
}
