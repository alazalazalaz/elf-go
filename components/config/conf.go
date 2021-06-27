package config

type Conf struct{
	Sys ConfSys
	Redis ConfRedis
	Mysql ConfMysql
}

//系统配置
type ConfSys struct{
	Debug bool
}

//redis配置
type ConfRedis struct{
	Ip string
	Port int
}

//mysql配置
type ConfMysql map[string]ConfMysqlDb

type ConfMysqlDb struct{
	Master []ConfMysqlItem
	Slaver []ConfMysqlItem
}

type ConfMysqlItem struct{
	Ip string
	Port int
	User string
	Password string
	Db string
}