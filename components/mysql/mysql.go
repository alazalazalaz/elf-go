package mysql

import (
	"elf-go/components/config"
	"elf-go/components/logs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(conf config.ConfMysql) (*gorm.DB, error){
	host := conf.Ip
	port := conf.Port
	user := conf.User
	pw := conf.Password
	dbName := conf.Db

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, port, dbName)
	logs.Info("mysql info :" + dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

