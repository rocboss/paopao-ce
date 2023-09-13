// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"

	"github.com/alimy/tryst/cache"
)

const (
	_defaultKeyPoolSize = 128
)

// 以下包含一些在cache中会用到的key的前缀
const (
	PrefixUserTweets = "paopao:usertweets:"
)

// 以下包含一些在cache中会用到的池化后的key
var (
	KeyUnreadMsg cache.KeyPool[int64]
)

func initCacheKeyPool() {
	poolSize := _defaultKeyPoolSize
	if poolSize < CacheSetting.KeyPoolSize {
		poolSize = CacheSetting.KeyPoolSize
	}
	KeyUnreadMsg = cache.MustKeyPool[int64](poolSize, func(key int64) string {
		return fmt.Sprintf("paopao:unreadmsg:%d", key)
	})
}
