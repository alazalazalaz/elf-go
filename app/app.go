package app

import (
	"elf-go/components/config"
	"elf-go/components/logs"
	"elf-go/components/mysql"
	"elf-go/components/redis"
	"go.uber.org/dig"
)

var container = dig.New()
var elfConf *config.Config
var elfRedis *redis.Redis
var elfMysql *mysql.Mysql

func init(){

	if err := container.Provide(config.New); err != nil {
		logs.Error("初始化配置文件失败：", logs.Content{"err": err})
	}

	if err := container.Provide(redis.New); err != nil {
		logs.Error("初始化redis失败：", logs.Content{"err": err})
	}

	if err := container.Provide(mysql.New); err != nil{
		logs.Error("初始化mysql 失败: ", logs.Content{"err": err})
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

func Redis() *redis.Redis{
	if elfRedis == nil{
		_ = container.Invoke(func(redisClient *redis.Redis){
			elfRedis = redisClient
			return
		})
	}
	return elfRedis
}

func Mysql() *mysql.Mysql{
	if elfMysql == nil{
		_ = container.Invoke(func(mysqlClient *mysql.Mysql) {
			elfMysql = mysqlClient
		})
	}
	return elfMysql
}
