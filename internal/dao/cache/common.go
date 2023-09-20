// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"encoding/gob"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

type cacheDataService struct {
	core.DataService
	ac core.AppCache
}

func NewCacheDataService(ds core.DataService) core.DataService {
	lazyInitial()
	return &cacheDataService{
		DataService: ds,
		ac:          _appCache,
	}
}

func (s *cacheDataService) GetUserByID(id int64) (res *ms.User, err error) {
	// 先从缓存获取， 不处理错误
	key := conf.KeyUserInfoById.Get(id)
	if data, xerr := s.ac.Get(key); xerr == nil {
		buf := bytes.NewBuffer(data)
		res = &ms.User{}
		err = gob.NewDecoder(buf).Decode(res)
		return
	}
	// 最后查库
	if res, err = s.DataService.GetUserByID(id); err == nil {
		// 更新缓存
		onCacheUserInfoEvent(key, res)
	}
	return
}

func (s *cacheDataService) GetUserByUsername(username string) (res *ms.User, err error) {
	// 先从缓存获取， 不处理错误
	key := conf.KeyUserInfoByName.Get(username)
	if data, xerr := s.ac.Get(key); xerr == nil {
		buf := bytes.NewBuffer(data)
		res = &ms.User{}
		err = gob.NewDecoder(buf).Decode(res)
		return
	}
	// 最后查库
	if res, err = s.DataService.GetUserByUsername(username); err == nil {
		// 更新缓存
		onCacheUserInfoEvent(key, res)
	}
	return
}

func (s *cacheDataService) IsMyFriend(userId int64, friendIds []int64) (res map[int64]bool, err error) {
	size := len(friendIds)
	res = make(map[int64]bool, size)
	if size == 0 {
		return
	}
	// 从缓存中获取
	key := conf.KeyMyFriendIds.Get(userId)
	if data, xerr := s.ac.Get(key); xerr == nil {
		bitmap := roaring64.New()
		if err = bitmap.UnmarshalBinary(data); err == nil {
			for _, friendId := range friendIds {
				res[friendId] = bitmap.Contains(uint64(friendId))
			}
			return
		}
	}
	// 直接查库并触发缓存更新事件
	OnCacheMyFriendIdsEvent(s.DataService, userId)
	return s.DataService.IsMyFriend(userId, friendIds)
}

func (s *cacheDataService) IsMyFollow(userId int64, followIds []int64) (res map[int64]bool, err error) {
	size := len(followIds)
	res = make(map[int64]bool, size)
	if size == 0 {
		return
	}
	// 从缓存中获取
	key := conf.KeyMyFollowIds.Get(userId)
	if data, xerr := s.ac.Get(key); xerr == nil {
		bitmap := roaring64.New()
		if err = bitmap.UnmarshalBinary(data); err == nil {
			for _, followId := range followIds {
				res[followId] = bitmap.Contains(uint64(followId))
			}
			return
		}
	}
	// 直接查库并触发缓存更新事件
	OnCacheMyFollowIdsEvent(s.DataService, userId, key)
	return s.DataService.IsMyFollow(userId, followIds)
}
