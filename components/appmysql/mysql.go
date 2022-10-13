package appmysql

import (
	"elf-go/components/appconfig"
	"elf-go/components/applogs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Conf appconfig.ConfMysql
	*gorm.DB
}

func New(c *appconfig.Config) *Mysql {
	return &Mysql{
		Conf: c.GetMysqlConfig(),
	}
}

func (m *Mysql) TestConnection() error {
	dsn := m.formatDsn(m.Conf["default"])
	var err error
	if m.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}

	return nil
}

//初始化多db和主从
func (m *Mysql) InitClauses() error {
	//var res dbresolver.DBResolver
	//for dbName := range m.Conf {
	//	var masterDBs []gorm.Dialector
	//	for _, configMysqlItem := range m.Conf[dbName].Master {
	//		masterDBs = append(masterDBs, mysql.Open(m.formatDsn(configMysqlItem)))
	//	}
	//	var slaverDBs []gorm.Dialector
	//	for _, configMysqlItem := range m.Conf[dbName].Slaver {
	//		slaverDBs = append(slaverDBs, mysql.Open(m.formatDsn(configMysqlItem)))
	//	}
	//
	//	con := dbresolver.Config{
	//		Sources:  masterDBs,
	//		Replicas: slaverDBs,
	//		Policy:   dbresolver.RandomPolicy{},
	//	}
	//	if dbName == "default" {
	//		res.Register(con)
	//	} else {
	//		res.Register(con, dbName)
	//	}
	//}
	//
	//if err := m.Use(&res); err != nil {
	//	return err
	//}

	return nil
}

func (m *Mysql) formatDsn(item appconfig.ConfMysqlItem) string {
	host := item.Ip
	port := item.Port
	user := item.User
	pw := item.Password
	dbName := item.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, port, dbName)
	applogs.Debug("mysql info :" + dsn)
	return dsn
}
