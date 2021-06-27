package config

type Conf struct{
	Redis ConfRedis
	Mysql ConfMysql
}

type ConfRedis struct{
	Ip string
	Port int
}

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