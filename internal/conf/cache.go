// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"

	"github.com/alimy/tryst/cache"
	"github.com/rocboss/paopao-ce/pkg/types"
)

const (
	_defaultKeyPoolSize = 128
)

// 以下包含一些在cache中会用到的key的前缀
const (
	InfixCommentDefault      = "default"
	InfixCommentHots         = "hots"
	InfixCommentNewest       = "newest"
	PrefixNewestTweets       = "paopao:newesttweets:"
	PrefixHotsTweets         = "paopao:hotstweets:"
	PrefixFollowingTweets    = "paopao:followingtweets:"
	PrefixUserTweets         = "paopao:usertweets:"
	PrefixUnreadmsg          = "paopao:unreadmsg:"
	PrefixOnlineUser         = "paopao:onlineuser:"
	PrefixIdxTweetsNewest    = "paopao:index:tweets:newest:"
	PrefixIdxTweetsHots      = "paopao:index:tweets:hots:"
	PrefixIdxTweetsFollowing = "paopao:index:tweets:following:"
	PrefixIdxTrends          = "paopao:index:trends:"
	PrefixMessages           = "paopao:messages:"
	PrefixUserInfo           = "paopao:user:info:"
	PrefixUserProfile        = "paopao:user:profile:"
	PrefixUserInfoById       = "paopao:user:info:id:"
	PrefixUserInfoByName     = "paopao:user:info:name:"
	prefixUserProfileByName  = "paopao:user:profile:name:"
	PrefixMyFriendIds        = "paopao:myfriendids:"
	PrefixMyFollowIds        = "paopao:myfollowids:"
	PrefixTweetComment       = "paopao:comment:"
	KeySiteStatus            = "paopao:sitestatus"
	KeyHistoryMaxOnline      = "history.max.online"
)

// 以下包含一些在cache中会用到的池化后的key
var (
	KeyNewestTweets      cache.KeyPool[int]
	KeyHotsTweets        cache.KeyPool[int]
	KeyFollowingTweets   cache.KeyPool[string]
	KeyUnreadMsg         cache.KeyPool[int64]
	KeyOnlineUser        cache.KeyPool[int64]
	KeyUserInfoById      cache.KeyPool[int64]
	KeyUserInfoByName    cache.KeyPool[string]
	KeyUserProfileByName cache.KeyPool[string]
	KeyMyFriendIds       cache.KeyPool[int64]
	KeyMyFollowIds       cache.KeyPool[int64]
)

func initCacheKeyPool() {
	poolSize := _defaultKeyPoolSize
	if poolSize < CacheSetting.KeyPoolSize {
		poolSize = CacheSetting.KeyPoolSize
	}
	KeyNewestTweets = intKeyPool[int](poolSize, PrefixNewestTweets)
	KeyHotsTweets = intKeyPool[int](poolSize, PrefixHotsTweets)
	KeyFollowingTweets = strKeyPool(poolSize, PrefixFollowingTweets)
	KeyUnreadMsg = intKeyPool[int64](poolSize, PrefixUnreadmsg)
	KeyOnlineUser = intKeyPool[int64](poolSize, PrefixOnlineUser)
	KeyUserInfoById = intKeyPool[int64](poolSize, PrefixUserInfoById)
	KeyUserInfoByName = strKeyPool(poolSize, PrefixUserInfoById)
	KeyUserProfileByName = strKeyPool(poolSize, prefixUserProfileByName)
	KeyMyFriendIds = intKeyPool[int64](poolSize, PrefixMyFriendIds)
	KeyMyFollowIds = intKeyPool[int64](poolSize, PrefixMyFollowIds)
}

func strKeyPool(size int, prefix string) cache.KeyPool[string] {
	return cache.MustKeyPool(size, func(key string) string {
		return fmt.Sprintf("%s%s", prefix, key)
	})
}

func intKeyPool[T types.Integer](size int, prefix string) cache.KeyPool[T] {
	return cache.MustKeyPool[T](size, intKey[T](prefix))
}

func intKey[T types.Integer](prefix string) func(T) string {
	return func(key T) string {
		return fmt.Sprintf("%s%d", prefix, key)
	}
}
