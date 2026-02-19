//go:build constraint

package cache

import (
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ tweetsCache = (*bigCacheTweetsCache)(nil)

	_ core.RedisCache = (*redisCache)(nil)
	_ tweetsCache     = (*redisCacheTweetsCache)(nil)

	_ core.CacheIndexService = (*noneCacheIndexServant)(nil)
	_ core.VersionInfo       = (*noneCacheIndexServant)(nil)

	_ core.CacheIndexService = (*simpleCacheIndexServant)(nil)
	_ core.VersionInfo       = (*simpleCacheIndexServant)(nil)
)
