package mysql

import (
	"elf-go/components/config"
	"elf-go/components/logs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct{
	Conf config.ConfMysql
	*gorm.DB
}

func New(c *config.Config)*Mysql{
	return &Mysql{
		Conf: c.GetMysqlConfig(),
	}
}

func(m *Mysql) Init()error{
	host := m.Conf.Ip
	port := m.Conf.Port
	user := m.Conf.User
	pw := m.Conf.Password
	dbName := m.Conf.Db

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, port, dbName)
	logs.Info("mysql info :" + dsn)
	var err error
	m.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}


