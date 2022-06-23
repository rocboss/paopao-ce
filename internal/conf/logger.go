package conf

import (
	"io"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newFileLogger() io.Writer {
	return &lumberjack.Logger{
		Filename:  loggerFileSetting.SavePath + "/" + loggerFileSetting.FileName + loggerFileSetting.FileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}
}

func setupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(LoggerSetting.logLevel())

	if CfgIf("LoggerFile") {
		out := newFileLogger()
		logrus.SetOutput(out)
	} else if CfgIf("LoggerZinc") {
		hook := newZincLogHook()
		logrus.SetOutput(io.Discard)
		logrus.AddHook(hook)
	} else if CfgIf("LoggerMeili") {
		hook := newMeiliLogHook()
		logrus.SetOutput(io.Discard)
		logrus.AddHook(hook)
	}
}
