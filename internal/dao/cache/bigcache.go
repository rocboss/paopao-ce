// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/Masterminds/semver/v3"
	"github.com/allegro/bigcache/v3"
)

var (
	_ tweetsCache = (*bigCacheTweetsCache)(nil)
)

type bigCacheTweetsCache struct {
	name    string
	version *semver.Version
	bc      *bigcache.BigCache
}

func (s *bigCacheTweetsCache) getTweetsBytes(key string) ([]byte, error) {
	return s.bc.Get(key)
}

func (s *bigCacheTweetsCache) setTweetsBytes(key string, bs []byte) error {
	return s.bc.Set(key, bs)
}

func (s *bigCacheTweetsCache) delTweets(keys []string) error {
	for _, k := range keys {
		s.bc.Delete(k)
	}
	return nil
}

func (s *bigCacheTweetsCache) allKeys() ([]string, error) {
	var keys []string
	for it := s.bc.Iterator(); it.SetNext(); {
		entry, err := it.Value()
		if err != nil {
			return nil, err
		}
		keys = append(keys, entry.Key())
	}
	return keys, nil
}

func (s *bigCacheTweetsCache) Name() string {
	return "BigCacheIndex"
}

func (s *bigCacheTweetsCache) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}
