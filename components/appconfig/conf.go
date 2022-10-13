package appconfig

import "github.com/sirupsen/logrus"

type Conf struct {
	Sys    ConfSys
	Redis  ConfRedis
	Mysql  ConfMysql
	Logrus ConfLogrus
}

//系统配置
type ConfSys struct {
	Debug      bool
	ListenPort int
}

//redis配置
type ConfRedis struct {
	Ip   string
	Port int
}

//mysql配置
type ConfMysql map[string]ConfMysqlDb

//logrus配置
type ConfLogrus struct {
	Level           logrus.Level
	WriteToFilePath string
}

type ConfMysqlDb struct {
	Master []ConfMysqlItem
	Slaver []ConfMysqlItem
}

type ConfMysqlItem struct {
	Ip       string
	Port     int
	User     string
	Password string
	Db       string
}
