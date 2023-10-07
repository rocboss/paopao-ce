// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/alimy/tryst/event"
	"github.com/alimy/tryst/pool"
)

type Event = event.Event

type EventManager interface {
	Start()
	Stop()
	OnEvent(event Event)
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

func NewEventManager(fn pool.RespFn[Event], opts ...pool.Option) EventManager {
	return &simpleEventManager{
		em: event.NewEventManager(fn, opts...),
	}
}
