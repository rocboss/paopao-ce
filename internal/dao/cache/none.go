// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ core.CacheIndexService = (*noneCacheIndexServant)(nil)
	_ core.VersionInfo       = (*noneCacheIndexServant)(nil)
)

type noneCacheIndexServant struct {
	ips core.IndexPostsService
}

func (s *noneCacheIndexServant) IndexPosts(user *core.User, offset int, limit int) (*core.IndexTweetList, error) {
	return s.ips.IndexPosts(user, offset, limit)
}

func (s *noneCacheIndexServant) SendAction(_act core.IdxAct, _post *core.Post) {
	// empty
}

func (s *noneCacheIndexServant) Name() string {
	return "NoneCacheIndex"
}

func (s *noneCacheIndexServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
