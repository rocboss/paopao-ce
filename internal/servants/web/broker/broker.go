// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package broker

import (
	"github.com/alimy/cfg"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/sirupsen/logrus"
)

var (
	ds                 core.DataService
	ts                 core.TweetSearchService
	oss                core.ObjectStorageService
	DisablePhoneVerify bool

	_redis core.RedisCache
)

func Initialize() {
	ds = dao.DataService()
	ts = dao.TweetSearchService()
	oss = dao.ObjectStorageService()
	DisablePhoneVerify = !cfg.If("Sms")
	_redis = cache.NewRedisCache()
}

// persistMediaContents 获取媒体内容并持久化
func persistMediaContents(contents []*PostContentItem) (items []string, err error) {
	items = make([]string, 0, len(contents))
	for _, item := range contents {
		switch item.Type {
		case core.ContentTypeImage,
			core.ContentTypeVideo,
			core.ContentTypeAudio,
			core.ContentTypeAttachment,
			core.ContentTypeChargeAttachment:
			items = append(items, item.Content)
			if err != nil {
				continue
			}
			if err = oss.PersistObject(oss.ObjectKey(item.Content)); err != nil {
				logrus.Errorf("service.persistMediaContents failed: %s", err)
			}
		}
	}
	return
}
