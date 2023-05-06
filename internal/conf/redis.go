// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"log"
	"sync"

	"github.com/rueian/rueidis"
)

var (
	_redisClient rueidis.Client
	_onceRedis   sync.Once
)

func MustRedisClient() rueidis.Client {
	_onceRedis.Do(func() {
		client, err := rueidis.NewClient(rueidis.ClientOption{
			InitAddress:      redisSetting.InitAddress,
			Username:         redisSetting.Username,
			Password:         redisSetting.Password,
			SelectDB:         redisSetting.SelectDB,
			ConnWriteTimeout: redisSetting.ConnWriteTimeout,
		})
		if err != nil {
			log.Fatalf("create a redis client failed: %s", err)
		}
		_redisClient = client
	})
	return _redisClient
}
