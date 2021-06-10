package config

type Conf struct{
	Redis ConfRedis
	Mysql ConfMysql
}

type ConfRedis struct{
	Ip string
	Port int
}

type ConfMysql struct{
	Ip string
	Port int
	User string
	Password string
	Db string
}