package main

import (
	"elf-go/app"
	"elf-go/components/config"
	"elf-go/components/logs"
	"fmt"
)

type User struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Phone     int64
	Username  string
	Password  string
	ClassId   int
	ClassName string
	IdCard    string
	Level     int
	Status    int
	SchoolNum int
	CreatedAt int
}

func (User) TableName() string {
	return "t_user"
}

func main() {
	//初始化配置文件
	app.Config().SetConfigFilePath("conf.yml")
	if err := app.Config().Init(); err != nil {
		logs.Error(err.Error(), nil)
	}

	//初始化redis
	if err := app.Redis().Init(); err != nil {
		logs.Error(err.Error(), nil)
	}

	//初始化mysql
	//读操作默认使用从库
	//写操作默认使用主库
	db := app.Mysql()
	if err := db.Init(); err != nil {
		logs.Error(err.Error(), nil)
	}
	for i := 0; i < 10; i++ {
		var user User
		db.Where("id=?", 10).Find(&user)
		fmt.Println(user)
	}

	//手动切换，强制使用主库来查询
	//db.Clauses(dbresolver.Write).Find(&user)

	//使用其他db
	//db.Clauses(dbresolver.Use("backend")).Find(&user)

}

func _format(item config.ConfMysqlItem) string {
	host := item.Ip
	port := item.Port
	user := item.User
	pw := item.Password
	dbName := item.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, port, dbName)
	logs.Info("mysql info :" + dsn)
	return dsn
}
