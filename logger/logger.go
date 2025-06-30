package logger

import (
	"KYC/iternals/configs"
	"fmt"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

func Init() error {
	logParams := configs.AppSettings.LogParams
	if _, err := os.Stat(logParams.Directory); os.IsNotExist(err) {
		err = os.Mkdir(logParams.Directory, 0755)
		if err != nil {
			return err
		}
	}

	lumberLogInfo := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.Directory, logParams.Info),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.Directory, logParams.Error),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogWarn := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.Directory, logParams.Warn),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.Directory, logParams.Debug),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	Info = log.New(lumberLogInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lumberLogError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(lumberLogWarn, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(lumberLogDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	return nil
}
