// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	_redisClient *redis.Client
	_onceRedis   sync.Once
)

func MustRedis() *redis.Client {
	_onceRedis.Do(func() {
		_redisClient = redis.NewClient(&redis.Options{
			Addr:     redisSetting.Host,
			Password: redisSetting.Password,
			DB:       redisSetting.DB,
		})
	})
	return _redisClient
}
