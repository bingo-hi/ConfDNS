package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger(filepath string, level string) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		// 如果文件打开失败，仍输出到控制台
		log.Warn("无法打开日志文件，使用默认输出")
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		parsedLevel = logrus.InfoLevel
	}
	log.SetLevel(parsedLevel)
}

// LogDebugf 输出 Debug 级别日志
func LogDebugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// LogInfof 输出 Info 级别日志
func LogInfof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// LogWarnf 输出 Warning 级别日志
func LogWarnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// LogErrorf 输出 Error 级别日志
func LogErrorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// LogFatalf 输出 Fatal 级别日志并退出
func LogFatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
