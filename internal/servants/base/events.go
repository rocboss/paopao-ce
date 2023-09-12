// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/alimy/tryst/event"
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

type pushPostToSearchEvent struct {
	event.UnimplementedEvent
	fn   func(*ms.Post)
	post *ms.Post
}

type pushAllPostToSearchEvent struct {
	event.UnimplementedEvent
	fn func() error
}

func (p *pushPostToSearchEvent) Name() string {
	return "servants.base.pushPostToSearchEvent"
}

func (p *pushPostToSearchEvent) Action() (err error) {
	p.fn(p.post)
	return
}

func (p *pushAllPostToSearchEvent) Name() string {
	return "servants.base.pushAllPostToSearchEvent"
}

func (p *pushAllPostToSearchEvent) Action() error {
	return p.fn()
}
