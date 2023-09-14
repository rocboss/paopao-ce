// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package base

import (
	"fmt"

	"github.com/alimy/tryst/event"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/pkg/json"
)

type CacheRespEvent struct {
	event.UnimplementedEvent
	ac     core.AppCache
	key    string
	data   any
	expire int64
}

type ExpireRespEvent struct {
	event.UnimplementedEvent
	ac   core.AppCache
	keys []string
}

type ExpireAnyRespEvent struct {
	event.UnimplementedEvent
	ac      core.AppCache
	pattern string
}

type pushPostToSearchEvent struct {
	event.UnimplementedEvent
	fn   func(*ms.Post)
	post *ms.Post
}

type pushAllPostToSearchEvent struct {
	event.UnimplementedEvent
	fn func() error
}

func OnCacheRespEvent(ac core.AppCache, key string, data any, expire int64) {
	events.OnEvent(&CacheRespEvent{
		ac:     ac,
		key:    key,
		data:   data,
		expire: expire,
	})
}

func OnExpireRespEvent(ac core.AppCache, keys ...string) {
	if len(keys) != 0 {
		events.OnEvent(&ExpireRespEvent{
			ac:   ac,
			keys: keys,
		})
	}
}

func OnExpireAnyRespEvent(ac core.AppCache, pattern string) {
	events.OnEvent(&ExpireAnyRespEvent{
		ac:      ac,
		pattern: pattern,
	})
}

func (p *CacheRespEvent) Name() string {
	return "servants.base.CacheRespEvent"
}

func (p *CacheRespEvent) Action() error {
	if p.ac.Exist(p.key) {
		// do nothing
		return nil
	}
	resp := &joint.JsonResp{
		Code: 0,
		Msg:  "success",
		Data: p.data,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("CacheRespEvent action marshal resp occurs error: %w", err)
	}
	if err = p.ac.Set(p.key, data, p.expire); err != nil {
		return fmt.Errorf("CacheRespEvent action put resp data to redis cache occurs error: %w", err)
	}
	return nil
}

func (p *ExpireRespEvent) Name() string {
	return "servants.base.ExpireRespEvent"
}

func (p *ExpireRespEvent) Action() error {
	return p.ac.Delete(p.keys...)
}

func (p *ExpireAnyRespEvent) Name() string {
	return "servants.base.ExpireAnyRespEvent"
}

func (p *ExpireAnyRespEvent) Action() error {
	return p.ac.DelAny(p.pattern)
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
