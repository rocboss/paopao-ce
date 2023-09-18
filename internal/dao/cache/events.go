// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/alimy/tryst/event"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/events"
)

type expireIndexTweetsEvent struct {
	event.UnimplementedEvent
	tweet       *ms.Post
	ac          core.AppCache
	keysPattern []string
}

type expireHotsTweetsEvent struct {
	event.UnimplementedEvent
	tweet      *ms.Post
	ac         core.AppCache
	keyPattern string
}

type expireFollowTweetsEvent struct {
	event.UnimplementedEvent
	tweet      *ms.Post
	ac         core.AppCache
	keyPattern string
}

type cacheUserInfoEvent struct {
	event.UnimplementedEvent
	tweet  *ms.Post
	ac     core.AppCache
	key    string
	data   *ms.User
	expire int64
}

func onExpireIndexTweetEvent(tweet *ms.Post) {
	events.OnEvent(&expireIndexTweetsEvent{
		tweet: tweet,
		ac:    _appCache,
		keysPattern: []string{
			conf.PrefixIdxTweetsNewest + "*",
			conf.PrefixIdxTweetsHots + "*",
			fmt.Sprintf("%s%d:*", conf.PrefixUserTweets, tweet.UserID),
		},
	})
}

func onExpireHotsTweetEvent(tweet *ms.Post) {
	events.OnEvent(&expireHotsTweetsEvent{
		tweet:      tweet,
		ac:         _appCache,
		keyPattern: conf.PrefixHotsTweets + "*",
	})
}

func onExpireFollowTweetEvent(tweet *ms.Post) {
	events.OnEvent(&expireFollowTweetsEvent{
		tweet:      tweet,
		ac:         _appCache,
		keyPattern: conf.PrefixFollowingTweets + "*",
	})
}

func onCacheUserInfoEvent(key string, data *ms.User) {
	events.OnEvent(&cacheUserInfoEvent{
		key:    key,
		data:   data,
		ac:     _appCache,
		expire: conf.CacheSetting.UserInfoExpire,
	})
}

func (e *expireIndexTweetsEvent) Name() string {
	return "expireIndexTweetsEvent"
}

func (e *expireIndexTweetsEvent) Action() (err error) {
	// logrus.Debug("expireIndexTweetsEvent action running")
	for _, pattern := range e.keysPattern {
		e.ac.DelAny(pattern)
	}
	return
}

func (e *expireHotsTweetsEvent) Name() string {
	return "expireHotsTweetsEvent"
}

func (e *expireHotsTweetsEvent) Action() (err error) {
	// logrus.Debug("expireHotsTweetsEvent action running")
	e.ac.DelAny(e.keyPattern)
	return
}

func (e *expireFollowTweetsEvent) Name() string {
	return "expireFollowTweetsEvent"
}

func (e *expireFollowTweetsEvent) Action() (err error) {
	// logrus.Debug("expireFollowTweetsEvent action running")
	e.ac.DelAny(e.keyPattern)
	return
}

func (e *cacheUserInfoEvent) Name() string {
	return "cacheUserInfoEvent"
}

func (e *cacheUserInfoEvent) Action() (err error) {
	buffer := &bytes.Buffer{}
	ge := gob.NewEncoder(buffer)
	if err = ge.Encode(e.data); err == nil {
		e.ac.Set(e.key, buffer.Bytes(), e.expire)
	}
	return
}
