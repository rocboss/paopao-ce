// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"encoding/json"
	"fmt"

	"github.com/alimy/tryst/event"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/sirupsen/logrus"
)

type cacheUnreadMsgEvent struct {
	event.UnimplementedEvent
	ds  core.DataService
	wc  core.WebCache
	uid int64
}

type createMessageEvent struct {
	event.UnimplementedEvent
	ds      core.DataService
	wc      core.WebCache
	message *ms.Message
}

type createTweetEvent struct {
	event.UnimplementedEvent
	tweet *ms.Post
	ac    core.AppCache
}

func onCacheUnreadMsgEvent(uid int64) {
	events.OnEvent(&cacheUnreadMsgEvent{
		ds:  _ds,
		wc:  _wc,
		uid: uid,
	})
}

func onCreateMessageEvent(data *ms.Message) {
	events.OnEvent(&createMessageEvent{
		ds:      _ds,
		wc:      _wc,
		message: data,
	})
}

func onCreateTweetEvent(tweet *ms.Post) {
	events.OnEvent(&createTweetEvent{
		ac:    _ac,
		tweet: tweet,
	})
}

func (e *cacheUnreadMsgEvent) Name() string {
	return "cacheUnreadMsgEvent"
}

func (e *cacheUnreadMsgEvent) Action() error {
	if e.wc.ExistUnreadMsgCountResp(e.uid) {
		// do nothing
		return nil
	}
	count, err := e.ds.GetUnreadCount(e.uid)
	if err != nil {
		return fmt.Errorf("cacheUnreadMsgEvent action occurs error: %w", err)
	}
	resp := &joint.JsonResp{
		Code: 0,
		Msg:  "success",
		Data: &web.GetUnreadMsgCountResp{
			Count: count,
		},
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cacheUnreadMsgEvent action marshal resp occurs error: %w", err)
	}
	if err = e.wc.PutUnreadMsgCountResp(e.uid, data); err != nil {
		return fmt.Errorf("cacheUnreadMsgEvent action put resp data to redis cache occurs error: %w", err)
	}
	return nil
}

func (e *createMessageEvent) Name() string {
	return "createMessageEvent"
}

func (e *createMessageEvent) Action() (err error) {
	if _, err = e.ds.CreateMessage(e.message); err == nil {
		err = e.wc.DelUnreadMsgCountResp(e.message.ReceiverUserID)
	}
	return
}

func (e *createTweetEvent) Name() string {
	return "createTweetEvent"
}

func (e *createTweetEvent) Action() (err error) {
	// TODO: 过期缓存，重新计算rank等
	logrus.Debug("createTweetEvent post action running")
	return
}
