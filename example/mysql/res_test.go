package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
)

func main() {
	dsn0 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", 3306, "test")
	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", 3306, "test1")
	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", 3306, "test2")

	db, err := gorm.Open(mysql.Open(dsn0), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err1 := db.Use(dbresolver.Register(dbresolver.Config{
		// `db2` 作为 sources，`db3`、`db4` 作为 replicas
		Sources:  []gorm.Dialector{mysql.Open(dsn0)},
		Replicas: []gorm.Dialector{mysql.Open(dsn1), mysql.Open(dsn2)},
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		var user User
		db.Where("id=?", 10).Find(&user)
		fmt.Println(user)
	}

	db2, _ := gorm.Open(mysql.Open(dsn0), &gorm.Config{})
	for i := 0; i < 10; i++ {
		var user User
		db2.Where("id=?", 10).Find(&user)
		fmt.Println(user)
	}
}
