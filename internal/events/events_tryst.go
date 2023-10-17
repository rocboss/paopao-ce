// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/alimy/tryst/event"
	"github.com/alimy/tryst/pool"
	"github.com/rocboss/paopao-ce/internal/core"
)

type Event = event.Event

type EventManager interface {
	Start()
	Stop()
	OnEvent(event Event)
}

type simpleWorkerHook struct {
	name  string
	cache core.MetricCache
}

type simpleEventManager struct {
	em event.EventManager
}

func (s *simpleEventManager) Start() {
	s.em.Start()
}

func (s *simpleEventManager) Stop() {
	s.em.Stop()
}

func (s *simpleEventManager) OnEvent(event Event) {
	s.em.OnEvent(event)
}

func (h *simpleWorkerHook) OnJoin(count int) {
	h.cache.SetEventTempWorkerCount(h.name, count)
}

func (h *simpleWorkerHook) OnLeave(count int) {
	h.cache.SetEventTempWorkerCount(h.name, count)
}

func NewEventManager(fn pool.RespFn[Event], opts ...pool.Option) EventManager {
	return &simpleEventManager{
		em: event.NewEventManager(fn, opts...),
	}
}

func NewEventWorkerHook(name string, mc core.MetricCache) pool.WorkerHook {
	return &simpleWorkerHook{
		name:  name,
		cache: mc,
	}
}
