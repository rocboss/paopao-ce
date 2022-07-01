package dao

import (
	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
)

func newNoneCacheIndexServant(getIndexPosts indexPostsFunc) (*noneCacheIndexServant, versionInfo) {
	obj := &noneCacheIndexServant{
		getIndexPosts: getIndexPosts,
	}
	return obj, obj
}

func (s *noneCacheIndexServant) IndexPosts(user *model.User, offset int, limit int) ([]*model.PostFormated, error) {
	return s.getIndexPosts(user, offset, limit)
}

func (s *noneCacheIndexServant) SendAction(act core.IndexActionT) {
	// empty
}

func (s *noneCacheIndexServant) name() string {
	return "NoneCacheIndex"
}

func (s *noneCacheIndexServant) version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
