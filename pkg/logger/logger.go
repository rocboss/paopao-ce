package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/rocboss/paopao-ce/global"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/resty.v1"
)

type ZincLogIndex struct {
	Index map[string]string `json:"index"`
}

type ZincLogData struct {
	Time    time.Time     `json:"time"`
	Level   logrus.Level  `json:"level"`
	Message string        `json:"message"`
	Data    logrus.Fields `json:"data"`
}

type ZincLogHook struct {
	Fired bool
}

func (hook *ZincLogHook) Fire(entry *logrus.Entry) error {
	index := &ZincLogIndex{
		Index: map[string]string{
			"_index": global.LoggerZincSetting.Index,
		},
	}
	indexBytes, _ := json.Marshal(index)

	data := &ZincLogData{
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
		SetBasicAuth(global.LoggerZincSetting.User, global.LoggerZincSetting.Password).
		SetBody(logStr).
		Post(global.LoggerZincSetting.Host); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func (hook *ZincLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func New() (*logrus.Logger, error) {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}

	if global.CfgIf("LoggerFile") {
		s := global.LoggerFileSetting
		log.Out = &lumberjack.Logger{
			Filename:  s.SavePath + "/" + s.FileName + s.FileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}
	} else if global.CfgIf("LoggerZinc") {
		log.Out = io.Discard
		log.AddHook(&ZincLogHook{})
	}

	return log, nil
}
