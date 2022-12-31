// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"
	"time"

	"github.com/rocboss/paopao-ce/pkg/json"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

type zincLogData struct {
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

func newZincLogHook() *zincLogHook {
	return &zincLogHook{
		host:     loggerZincSetting.Endpoint() + "/es/_bulk",
		index:    loggerZincSetting.Index,
		user:     loggerZincSetting.User,
		password: loggerZincSetting.Password,
	}
}
