// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"log"
	"net/http"
	"time"

	hx "github.com/rocboss/paopao-ce/pkg/http"
	"github.com/rocboss/paopao-ce/pkg/json"
	"github.com/rocboss/paopao-ce/pkg/obx"
	"github.com/sirupsen/logrus"
)

type observeLogData struct {
	Time    time.Time     `json:"time"`
	Level   logrus.Level  `json:"level"`
	Message string        `json:"message"`
	Data    logrus.Fields `json:"data"`
}

type observeLogHook struct {
	client obx.OpenObserveClient
}

func (h *observeLogHook) Fire(entry *logrus.Entry) error {
	info := []observeLogData{{
		Time:    entry.Time,
		Level:   entry.Level,
		Message: entry.Message,
		Data:    entry.Data,
	}}
	data, _ := json.Marshal(info)
	h.client.LogJson(data)
	return nil
}

func (h *observeLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func newObserveLogHook() *observeLogHook {
	s := loggerOpenObserveSetting
	obc := &obx.Config{
		Host:         s.Host,
		User:         s.User,
		Password:     s.Password,
		Organization: s.Organization,
		Stream:       s.Stream,
		Secure:       s.Secure,
	}
	acc := &hx.AsyncClientConf{
		MinWorker:         s.MinWorker,
		MaxRequestBuf:     s.MaxLogBuffer,
		MaxRequestTempBuf: 100,
		MaxIdleTime:       60 * time.Second,
	}
	return &observeLogHook{
		client: obx.NewClient(obc, acc, func(req *http.Request, resp *http.Response, err error) {
			if err == nil && resp != nil && resp.Body != nil {
				resp.Body.Close()
			} else if err != nil {
				log.Printf("logrus use observe do LogJson error: %s", err)
			}
		}),
	}
}
