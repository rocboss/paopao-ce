// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package statistics

import (
	"sync"

	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_metricCache     core.MetricCache
	_onceMetricCache sync.Once
)

type metricCache struct {
	eventTempWorkerCount map[string]int32
}

func (m *metricCache) SetEventTempWorkerCount(name string, count int32) {
	// 直接赋值，不需要加锁，因为这仅仅是一个统计信息
	m.eventTempWorkerCount[name] = count
}

func (m *metricCache) GetEventTempWorkerCount(name string) int32 {
	return m.eventTempWorkerCount[name]
}

func NewMetricCache() core.MetricCache {
	_onceMetricCache.Do(func() {
		_metricCache = &metricCache{
			eventTempWorkerCount: make(map[string]int32),
		}
	})
	return _metricCache
}
