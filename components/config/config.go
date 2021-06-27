package config

import "C"
import (
	"github.com/spf13/viper"
)


type Config struct{
	//配置文件路径
	confPath string

	//总的配置文件
	C Conf
}

func New() *Config{
	return &Config{
		confPath: "",
		C: Conf{},
	}
}


func(c *Config) Init() error {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigFile(c.confPath)
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&c.C); err != nil{
		return err
	}

	return nil
}

func(c *Config) SetConfigFilePath(p string){
	c.confPath = p
}

func(c *Config) GetRedisConfig() ConfRedis{
	return c.C.Redis
}

func(c *Config) GetMysqlConfig() ConfMysql{
	return c.C.Mysql
}
