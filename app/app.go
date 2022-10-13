package app

import (
	"elf-go/components/appconfig"
	"elf-go/components/applogs"
	"elf-go/components/appmysql"
	"elf-go/components/appredis"
	"go.uber.org/dig"
)

var container = dig.New()
var elfConf *appconfig.Config
var elfRedis *appredis.Redis
var elfMysql *appmysql.Mysql

func init() {

	if err := container.Provide(appconfig.New); err != nil {
		applogs.Error("初始化配置文件失败：", applogs.Content{"err": err})
	}

	if err := container.Provide(appredis.New); err != nil {
		applogs.Error("初始化redis失败：", applogs.Content{"err": err})
	}

	if err := container.Provide(appmysql.New); err != nil {
		applogs.Error("初始化mysql 失败: ", applogs.Content{"err": err})
	}
}

func Config() *appconfig.Config {
	if elfConf == nil {
		_ = container.Invoke(func(conf *appconfig.Config) {
			elfConf = conf
			return
		})
	}
	return elfConf
}

func Redis() *appredis.Redis {
	if elfRedis == nil {
		_ = container.Invoke(func(redisClient *appredis.Redis) {
			elfRedis = redisClient
			return
		})
	}
	return elfRedis
}

func Mysql() *appmysql.Mysql {
	if elfMysql == nil {
		_ = container.Invoke(func(mysqlClient *appmysql.Mysql) {
			elfMysql = mysqlClient
		})
	}
	return elfMysql
}
