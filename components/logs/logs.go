package logs

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式
	//logrus.SetFormatter(&logrus.TextFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为InfoLevel以上
	logrus.SetLevel(logrus.InfoLevel)

}

type Content map[string]interface{}

func Info(s string, contentSlice ...Content){
	content := Content{}
	if len(contentSlice) > 0{
		content = contentSlice[0]
	}
	logrus.WithFields(logrus.Fields(content)).Info(s)
}

func Warning(s string, content Content){
	logrus.WithFields(logrus.Fields(content)).Warning(s)
}

func Error(s string, content Content){
	logrus.WithFields(logrus.Fields(content)).Error(s)
}