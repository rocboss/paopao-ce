// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"io"

	"github.com/rocboss/paopao-ce/pkg/cfg"
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
	logrus.SetLevel(loggerSetting.logLevel())

	cfg.In(cfg.Actions{
		"LoggerFile": func() {
			out := newFileLogger()
			logrus.SetOutput(out)
		},
		"LoggerZinc": func() {
			hook := newZincLogHook()
			logrus.SetOutput(io.Discard)
			logrus.AddHook(hook)
		},
		"LoggerMeili": func() {
			hook := newMeiliLogHook()
			logrus.SetOutput(io.Discard)
			logrus.AddHook(hook)
		},
	})
}
