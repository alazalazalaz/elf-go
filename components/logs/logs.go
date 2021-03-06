package logs

import (
	"elf-go/utils/traceid"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为InfoLevel以上，由初始化的时候去设置
	//logrus.SetLevel(logrus.InfoLevel)
}

type Content map[string]interface{}

func SetLevel(l logrus.Level) {
	logrus.SetLevel(l)
}

func Debug(s string, contentSlice ...Content) {
	content := Content{}
	if len(contentSlice) > 0 {
		content = contentSlice[0]
	}
	logrus.WithFields(logrus.Fields(content)).Debug(s)
}

func Info(s string, contentSlice ...Content) {
	content := Content{}
	if len(contentSlice) > 0 {
		content = contentSlice[0]
	}
	logrus.WithFields(logrus.Fields(content)).Info(s)
}

func Warning(s string, contentSlice ...Content) {
	content := Content{}
	if len(contentSlice) > 0 {
		content = contentSlice[0]
	}
	logrus.WithFields(logrus.Fields(content)).Warning(s)
}

func Error(s string, contentSlice ...Content) {
	content := Content{}
	if len(contentSlice) > 0 {
		content = contentSlice[0]
	}
	logrus.WithFields(logrus.Fields(content)).Error(s)
}

func Debugf(format string, args ...interface{}) {
	format = fmt.Sprintf("[%s]", traceid.TraceId) + format
	logrus.Infof(format, args...)
}

func Infof(format string, args ...interface{}) {
	format = fmt.Sprintf("[%s]", traceid.TraceId) + format
	logrus.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	format = fmt.Sprintf("[%s]", traceid.TraceId) + format
	logrus.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	format = fmt.Sprintf("[%s]", traceid.TraceId) + format
	logrus.Infof(format, args...)
}
