// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/redis/rueidis"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/utils"
)

var (
	_webCache core.WebCache = (*redisWebCache)(nil)
)

type redisWebCache struct {
	cscExpire       time.Duration
	unreadMsgExpire int64
	c               rueidis.Client
}

func (s *redisWebCache) Name() string {
	return "RedisWebCache"
}

func (s *redisWebCache) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *redisWebCache) GetUnreadMsgCountResp(uid int64) ([]byte, error) {
	key := conf.KeyUnreadMsg.Get(uid)
	res, err := rueidis.MGetCache(s.c, context.Background(), s.cscExpire, []string{key})
	if err != nil {
		return nil, err
	}
	message := res[key]
	return message.AsBytes()
}

func (s *redisWebCache) PutUnreadMsgCountResp(uid int64, data []byte) error {
	return s.c.Do(context.Background(), s.c.B().Set().
		Key(conf.KeyUnreadMsg.Get(uid)).
		Value(utils.String(data)).
		ExSeconds(s.unreadMsgExpire).
		Build()).
		Error()
}

func (s *redisWebCache) DelUnreadMsgCountResp(uid int64) error {
	return s.c.Do(context.Background(), s.c.B().Del().Key(conf.KeyUnreadMsg.Get(uid)).Build()).Error()
}

func newWebCache() *redisWebCache {
	s := conf.CacheSetting
	return &redisWebCache{
		cscExpire:       s.CientSideCacheExpire,
		unreadMsgExpire: s.UnreadMsgExpire,
		c:               conf.MustRedisClient(),
	}
}
