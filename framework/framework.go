package framework

import (
	"elf-go/app"
	"elf-go/components/config"
	"elf-go/components/logs"
	"github.com/sirupsen/logrus"
)

type Framework struct {
	Conf *config.ConfSys
}

//初始化
func Init(configFilePath string) error {
	//初始化配置文件
	app.Config().SetConfigFilePath(configFilePath)
	if err := app.Config().Init(); err != nil {
		logs.Error(err.Error(), nil)
		return err
	}

	//初始化系统配置
	if err := initSys(); err != nil {
		logs.Error(err.Error(), nil)
		return err
	}

	return nil
}

func initSys() error {
	sysConfig := app.Config().GetSysConfig()
	if sysConfig.Debug == true {
		logs.SetLevel(logrus.DebugLevel)
	} else {
		logs.SetLevel(logrus.InfoLevel)
	}

	return nil
}
