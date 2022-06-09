package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/resty.v1"
)

type zincLogIndex struct {
	Index map[string]string `json:"index"`
}

type zincLogData struct {
	Time    time.Time     `json:"time"`
	Level   logrus.Level  `json:"level"`
	Message string        `json:"message"`
	Data    logrus.Fields `json:"data"`
}

type zincLogHook struct {
	host     string
	index    string
	user     string
	password string
}

func (h *zincLogHook) Fire(entry *logrus.Entry) error {
	index := &zincLogIndex{
		Index: map[string]string{
			"_index": h.index,
		},
	}
	indexBytes, _ := json.Marshal(index)

	data := &zincLogData{
		Time:    entry.Time,
		Level:   entry.Level,
		Message: entry.Message,
		Data:    entry.Data,
	}
	dataBytes, _ := json.Marshal(data)

	logStr := string(indexBytes) + "\n" + string(dataBytes) + "\n"
	client := resty.New()

	if _, err := client.SetDisableWarn(true).R().
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(h.user, h.password).
		SetBody(logStr).
		Post(h.host); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func (h *zincLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func setupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if CfgIf("LoggerFile") {
		out := &lumberjack.Logger{
			Filename:  loggerFileSetting.SavePath + "/" + loggerFileSetting.FileName + loggerFileSetting.FileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}
		logrus.SetOutput(out)
	} else if CfgIf("LoggerZinc") {
		hook := &zincLogHook{
			host:     loggerZincSetting.Host,
			index:    loggerZincSetting.Index,
			user:     loggerZincSetting.User,
			password: loggerZincSetting.Password,
		}
		logrus.SetOutput(io.Discard)
		logrus.AddHook(hook)
	}
}
