// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"encoding/gob"

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
