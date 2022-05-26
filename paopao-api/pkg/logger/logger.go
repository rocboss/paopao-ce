package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/rocboss/paopao-api/global"
	"github.com/rocboss/paopao-api/pkg/setting"
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
			"_index": global.LoggerSetting.LogZincIndex,
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
		SetBasicAuth(global.LoggerSetting.LogZincUser, global.LoggerSetting.LogZincPassword).
		SetBody(logStr).
		Post(global.LoggerSetting.LogZincHost); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func (hook *ZincLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func New(s *setting.LoggerSettingS) (*logrus.Logger, error) {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}

	switch s.LogType {
	case setting.LogFileType:
		log.Out = &lumberjack.Logger{
			Filename:  s.LogFileSavePath + "/" + s.LogFileName + s.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}
	case setting.LogZincType:
		log.Out = io.Discard
		log.AddHook(&ZincLogHook{})
	}

	return log, nil
}
