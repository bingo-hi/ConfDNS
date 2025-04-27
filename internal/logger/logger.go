package logger

import (
	"confdns/internal/config"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log = logrus.New()

// InitLoggerWithConfig initializes the logging system (supports advanced configurations)
func InitLoggerWithConfig(cfg config.LogConfig) {
	// If no log path is set, the default will be used
	if cfg.FilePath == "" {
		cfg.FilePath = "dns.log"
	}

	// Default value set
	if cfg.MaxSizeMB == 0 {
		cfg.MaxSizeMB = 10 // Default is 10MB
	}
	if cfg.MaxAgeDays == 0 {
		cfg.MaxAgeDays = 7 // Default is 7 days
	}
	if cfg.MaxBackups == 0 {
		cfg.MaxBackups = 5 // Default keeps a maximum of 5 logs
	}

	// Set Lumberjack for log rotation
	log.SetOutput(&lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxSizeMB,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAgeDays,
		Compress:   cfg.Compress,
	})

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set log level
	parsedLevel, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		log.Warn("If log level parsing fails, the default is Info")
		parsedLevel = logrus.InfoLevel
	}
	log.SetLevel(parsedLevel)
}

// The following are functions for outputting logs at various levels
func LogDebugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func LogInfof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func LogWarnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func LogErrorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func LogFatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
