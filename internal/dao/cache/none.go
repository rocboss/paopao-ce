// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

type noneCacheIndexServant struct {
	ips core.IndexPostsService
}

func (s *noneCacheIndexServant) IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error) {
	return s.ips.IndexPosts(user, offset, limit)
}

func (s *noneCacheIndexServant) TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *noneCacheIndexServant) SendAction(_act core.IdxAct, _post *ms.Post) {
	// empty
}

func (s *noneCacheIndexServant) Name() string {
	return "NoneCacheIndex"
}

func (s *noneCacheIndexServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
