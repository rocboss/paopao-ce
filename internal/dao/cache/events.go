// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
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
