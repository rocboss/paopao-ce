package dao

import (
	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
)

func newNoneCacheIndexServant(getIndexPosts indexPostsFunc) *noneCacheIndexServant {
	return &noneCacheIndexServant{
		getIndexPosts: getIndexPosts,
	}
}

func (s *noneCacheIndexServant) IndexPosts(user *model.User, offset int, limit int) ([]*model.PostFormated, error) {
	return s.getIndexPosts(user, offset, limit)
}

func (s *noneCacheIndexServant) SendAction(act core.IndexActionT) {
	// empty
}

func (s *noneCacheIndexServant) Name() string {
	return "NoneCacheIndex"
}

func (s *noneCacheIndexServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
