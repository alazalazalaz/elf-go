package elf_go

import (
	"elf-go/app"
	"elf-go/components/appconfig"
	"elf-go/components/appfile"
	"elf-go/components/applogs"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Framework struct {
	Conf *appconfig.ConfSys
}

//初始化
func Init(configFilePath string) error {
	//初始化配置文件
	app.Config().SetConfigFilePath(configFilePath)
	if err := app.Config().Init(); err != nil {
		applogs.Error(err.Error(), nil)
		return err
	}

	//初始化系统配置
	if err := initSys(); err != nil {
		applogs.Error(err.Error(), nil)
		return err
	}

	//初始化log配置
	if err := initLog(); err != nil {
		applogs.Error(err.Error(), nil)
		return err
	}

	return nil
}

func initSys() error {
	//sysConfig := app.Config().GetSysConfig()
	//if sysConfig.Debug == true {
	//	logs.SetLevel(logrus.DebugLevel)
	//} else {
	//	logs.SetLevel(logrus.InfoLevel)
	//}

	return nil
}

func initLog() error {
	// 设置Log等级
	applogs.SetLevel(logrus.InfoLevel)
	logConfig := app.Config().GetLogrusConfig()
	for _, l := range logrus.AllLevels {
		if l == logConfig.Level {
			applogs.SetLevel(l)
			break
		}
	}

	//设置输出文件路径
	if logConfig.WriteToFilePath != "" {
		if !appfile.IsFile(logConfig.WriteToFilePath) {
			return errors.New("unknown file path from config key=WriteToFilePath,value=" + logConfig.WriteToFilePath)
		}

		fileWrite, err := os.OpenFile(logConfig.WriteToFilePath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		mw := io.MultiWriter(os.Stdout, fileWrite)
		logrus.SetOutput(mw)
		applogs.Infof("log will be write to file :%v", logConfig.WriteToFilePath)
	}

	return nil
}
