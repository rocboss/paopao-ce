// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package metrics

import (
	"sync"

	"github.com/alimy/tryst/event"
	"github.com/alimy/tryst/pool"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

var (
	_defaultMetricManager event.EventManager
	_onceInitial          sync.Once
)

type Metric = event.Event

type BaseMetric = event.UnimplementedEvent

type MetricManager interface {
	Start()
	Stop()
	OnMeasure(metric Metric)
}

func StartMetricManager() {
	_defaultMetricManager.Start()
}

func StopMetricManager() {
	_defaultMetricManager.Stop()
}

// OnMeasure push Metric to gorotine pool then handled automatic.
func OnMeasure(metric Metric) {
	_defaultMetricManager.OnEvent(metric)
}

func Initial() {
	_onceInitial.Do(func() {
		initMetricManager()
	})
}

func initMetricManager() {
	var opts []pool.Option
	s := conf.EventManagerSetting
	if s.MinWorker > 5 {
		opts = append(opts, pool.MinWorkerOpt(s.MinWorker))
	} else {
		opts = append(opts, pool.MinWorkerOpt(5))
	}
	if s.MaxEventBuf > 10 {
		opts = append(opts, pool.MaxRequestBufOpt(s.MaxEventBuf))
	} else {
		opts = append(opts, pool.MaxRequestBufOpt(10))
	}
	if s.MaxTempEventBuf > 10 {
		opts = append(opts, pool.MaxRequestTempBufOpt(s.MaxTempEventBuf))
	} else {
		opts = append(opts, pool.MaxRequestTempBufOpt(10))
	}
	opts = append(opts, pool.MaxTickCountOpt(s.MaxTickCount), pool.TickWaitTimeOpt(s.TickWaitTime))
	_defaultMetricManager = event.NewEventManager(func(req Metric, err error) {
		if err != nil {
			logrus.Errorf("handle event[%s] occurs error: %s", req.Name(), err)
		}
	}, opts...)
}
