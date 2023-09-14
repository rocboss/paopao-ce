// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package metrics

import (
	"github.com/alimy/tryst/event"
	"github.com/alimy/tryst/pool"
)

type simpleMetricManager struct {
	mm event.EventManager
}

func (s *simpleMetricManager) Start() {
	s.mm.Start()
}

func (s *simpleMetricManager) Stop() {
	s.mm.Stop()
}

func (s *simpleMetricManager) OnMeasure(metric Metric) {
	s.mm.OnEvent(metric)
}

func NewMetricManager(fn pool.RespFn[Metric], opts ...pool.Option) MetricManager {
	return &simpleMetricManager{
		mm: event.NewEventManager(fn, opts...),
	}
}
