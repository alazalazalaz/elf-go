package config

import (
	"github.com/spf13/viper"
)

var C Conf

func InitConfig(configPath string) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&C); err != nil{
		return err
	}

	return nil
}

func GetRedisConfig() ConfRedis{
	return C.Redis
}

func GetMysqlConfig() ConfMysql{
	return C.Mysql
}
