// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"io"
	"time"

	"github.com/alimy/cfg"
	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
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

func setupSentryLogrus(opts sentry.ClientOptions) {
	// Send only ERROR and higher level logs to Sentry
	sentryLevels := []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	sentryHook, err := sentrylogrus.New(sentryLevels, opts)
	if err != nil {
		panic(err)
	}
	logrus.AddHook(sentryHook)
	// Flushes before calling os.Exit(1) when using logger.Fatal
	// (else all defers are not called, and Sentry does not have time to send the event)
	logrus.RegisterExitHandler(func() { sentryHook.Flush(5 * time.Second) })
}
