package conf

import (
	"fmt"
	"io"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rocboss/paopao-ce/pkg/json"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/resty.v1"
)

type logData struct {
	Time    time.Time     `json:"time"`
	Level   logrus.Level  `json:"level"`
	Message string        `json:"message"`
	Data    logrus.Fields `json:"data"`
}

type zincLogIndex struct {
	Index map[string]string `json:"index"`
}

type zincLogHook struct {
	host     string
	index    string
	user     string
	password string
}

type meiliLogHook struct {
	index *meilisearch.Index
}

func (h *zincLogHook) Fire(entry *logrus.Entry) error {
	index := &zincLogIndex{
		Index: map[string]string{
			"_index": h.index,
		},
	}
	indexBytes, _ := json.Marshal(index)

	data := &logData{
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

func (h *meiliLogHook) Fire(entry *logrus.Entry) error {
	data := &logData{
		Time:    entry.Time,
		Level:   entry.Level,
		Message: entry.Message,
		Data:    entry.Data,
	}
	if _, err := h.index.AddDocuments(data); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (h *meiliLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func newZincLogHook() *zincLogHook {
	return &zincLogHook{
		host:     loggerZincSetting.Endpoint() + "/es/_bulk",
		index:    loggerZincSetting.Index,
		user:     loggerZincSetting.User,
		password: loggerZincSetting.Password,
	}
}

func newMeiliLogHook() *meiliLogHook {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   loggerMeiliSetting.Endpoint(),
		APIKey: loggerMeiliSetting.ApiKey,
	})

	index := client.Index(loggerMeiliSetting.Index)
	if _, err := index.FetchInfo(); err != nil {
		logrus.Debugf("newMeiliLogHook create index because fetch index info error: %v", err)
		client.CreateIndex(&meilisearch.IndexConfig{
			Uid: loggerMeiliSetting.Index,
		})
	}

	return &meiliLogHook{
		index: index,
	}
}

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
