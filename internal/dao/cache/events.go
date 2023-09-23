// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/alimy/tryst/event"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/sirupsen/logrus"
)

type expireIndexTweetsEvent struct {
	event.UnimplementedEvent
	ac          core.AppCache
	keysPattern []string
}

type expireHotsTweetsEvent struct {
	event.UnimplementedEvent
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

type cacheMyFriendIdsEvent struct {
	event.UnimplementedEvent
	ac      core.AppCache
	urs     core.UserRelationService
	userIds []int64
	expire  int64
}

type cacheMyFollowIdsEvent struct {
	event.UnimplementedEvent
	ac     core.AppCache
	urs    core.UserRelationService
	userId int64
	key    string
	expire int64
}

func OnExpireIndexTweetEvent(userId int64) {
	// TODO: 这里暴躁的将所有 最新/热门/关注 的推文列表缓存都过期掉，后续需要更精细话处理
	events.OnEvent(&expireIndexTweetsEvent{
		ac: _appCache,
		keysPattern: []string{
			conf.PrefixIdxTweetsNewest + "*",
			conf.PrefixIdxTweetsHots + "*",
			conf.PrefixIdxTweetsFollowing + "*",
			fmt.Sprintf("%s%d:*", conf.PrefixUserTweets, userId),
		},
	})
}

func OnExpireHotsTweetEvent() {
	events.OnEvent(&expireHotsTweetsEvent{
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

func OnCacheMyFriendIdsEvent(urs core.UserRelationService, userIds ...int64) {
	if len(userIds) == 0 {
		return
	}
	events.OnEvent(&cacheMyFriendIdsEvent{
		userIds: userIds,
		urs:     urs,
		ac:      _appCache,
		expire:  conf.CacheSetting.UserRelationExpire,
	})
}

func OnCacheMyFollowIdsEvent(urs core.UserRelationService, userId int64, key ...string) {
	cacheKey := ""
	if len(key) > 0 {
		cacheKey = key[0]
	} else {
		cacheKey = conf.KeyMyFollowIds.Get(userId)
	}
	events.OnEvent(&cacheMyFollowIdsEvent{
		userId: userId,
		urs:    urs,
		key:    cacheKey,
		ac:     _appCache,
		expire: conf.CacheSetting.UserRelationExpire,
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

func (e *cacheMyFriendIdsEvent) Name() string {
	return "cacheMyFriendIdsEvent"
}

func (e *cacheMyFriendIdsEvent) Action() error {
	logrus.Debug("cacheMyFriendIdsEvent action runnging")
	for _, userId := range e.userIds {
		myFriendIds, err := e.urs.MyFriendIds(userId)
		if err != nil {
			return err
		}
		bitmap := roaring64.New()
		for _, friendId := range myFriendIds {
			bitmap.Add(uint64(friendId))
		}
		data, err := bitmap.MarshalBinary()
		if err != nil {
			return err
		}
		e.ac.Set(conf.KeyMyFriendIds.Get(userId), data, e.expire)
	}
	return nil
}

func (e *cacheMyFollowIdsEvent) Name() string {
	return "cacheMyFollowIdsEvent"
}

func (e *cacheMyFollowIdsEvent) Action() (err error) {
	logrus.Debug("cacheMyFollowIdsEvent action runnging")
	myFollowIds, err := e.urs.MyFollowIds(e.userId)
	if err != nil {
		return err
	}
	bitmap := roaring64.New()
	for _, followId := range myFollowIds {
		bitmap.Add(uint64(followId))
	}
	data, err := bitmap.MarshalBinary()
	if err != nil {
		return err
	}
	e.ac.Set(e.key, data, e.expire)
	return nil
}
