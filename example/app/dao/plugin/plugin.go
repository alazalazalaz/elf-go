package plugin

import (
	"elf-go/components/applogs"
	"elf-go/example/app/enum"
	"gorm.io/gorm"
)

type BeforeAfterPlugin struct {
}

func (*BeforeAfterPlugin) Name() string {
	return "before_after_plugin"
}

// Initialize 函数会在加载插件的时候被自动调用
func (*BeforeAfterPlugin) Initialize(db *gorm.DB) error {
	if err := db.Callback().Query().Before("gorm:query").Register("my_plugin:before_query", before); err != nil {
		return err
	}
	if err := db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", before); err != nil {
		return err
	}
	if err := db.Callback().Update().Before("gorm:update").Register("my_plugin:before_update", before); err != nil {
		return err
	}
	if err := db.Callback().Delete().Before("gorm:delete").Register("my_plugin:before_delete", before); err != nil {
		return err
	}

	if err := db.Callback().Query().After("gorm:query").Register("my_plugin:after_query", afterQuery); err != nil {
		return err
	}
	if err := db.Callback().Create().After("gorm:create").Register("my_plugin:after_create", afterCreate); err != nil {
		return err
	}
	//update包含update()和save()方法
	if err := db.Callback().Update().After("gorm:update").Register("my_plugin:after_update", afterUpdate); err != nil {
		return err
	}
	if err := db.Callback().Delete().After("gorm:delete").Register("my_plugin:after_delete", afterDelete); err != nil {
		return err
	}

	return nil
}

func before(db *gorm.DB) {
	//sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	//logs.Infof("before sql:%s", sql)
}

func afterQuery(db *gorm.DB) {
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	printLog(enum.SqlTypeSelect, sql)
}

func afterCreate(db *gorm.DB) {
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	printLog(enum.SqlTypeInsert, sql)
}

func afterUpdate(db *gorm.DB) {
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	printLog(enum.SqlTypeUpdate, sql)
}

func afterDelete(db *gorm.DB) {
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	printLog(enum.SqlTypeDelete, sql)
}

func printLog(t, sql string) {
	applogs.Infof("after %s, sql:%s", t, sql)
}
