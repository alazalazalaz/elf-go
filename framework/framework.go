package framework

import (
	"elf-go/app"
	"elf-go/components/config"
	"elf-go/components/logs"
	"elf-go/utils/file"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
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

	//初始化log配置
	if err := initLog(); err != nil {
		logs.Error(err.Error(), nil)
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
	logs.SetLevel(logrus.InfoLevel)
	logConfig := app.Config().GetLogrusConfig()
	for _, l := range logrus.AllLevels {
		if l == logConfig.Level {
			logs.SetLevel(l)
			break
		}
	}

	//设置输出文件路径
	if logConfig.WriteToFilePath != "" {
		if !file.IsFile(logConfig.WriteToFilePath) {
			return errors.New("unknown file path from config key=WriteToFilePath,value=" + logConfig.WriteToFilePath)
		}

		fileWrite, err := os.OpenFile(logConfig.WriteToFilePath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		mw := io.MultiWriter(os.Stdout, fileWrite)
		logrus.SetOutput(mw)
		logs.Infof("log will be write to file :%v", logConfig.WriteToFilePath)
	}

	return nil
}
