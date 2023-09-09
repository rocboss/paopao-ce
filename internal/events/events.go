// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package events

import (
	"sync"

	"github.com/alimy/tryst/event"
	"github.com/alimy/tryst/pool"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

var (
	_defaultEventManager event.EventManager
	_onceInitial         sync.Once
)

// OnEvent push event to gorotine pool then handled automatic.
func OnEvent(event event.Event) {
	_defaultEventManager.OnEvent(event)
}

func Initial() {
	_onceInitial.Do(func() {
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
		_defaultEventManager = event.NewEventManager(func(req event.Event, err error) {
			if err != nil {
				logrus.Errorf("handle event[%s] occurs error: %s", req.Name(), err)
			}
		}, opts...)
	})
}

func Restart() {
	_defaultEventManager.Stop()
	_defaultEventManager.Start()
}

func Done() {
	_defaultEventManager.Stop()
}
