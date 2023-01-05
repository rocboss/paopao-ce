// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/Masterminds/semver/v3"
	"github.com/alimy/cfg"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/dao/security"
	"github.com/sirupsen/logrus"
)

var (
	_ core.DataService = (*dataServant)(nil)
	_ core.VersionInfo = (*dataServant)(nil)
)

type dataServant struct {
	core.IndexPostsService
	core.WalletService
	core.MessageService
	core.TopicService
	core.TweetService
	core.TweetManageService
	core.TweetHelpService
	core.CommentService
	core.CommentManageService
	core.UserManageService
	core.ContactManageService
	core.SecurityService
	core.AttachmentCheckService
}

func NewDataService() (core.DataService, core.VersionInfo) {
	// initialize CacheIndex if needed
	var (
		c core.CacheIndexService
		v core.VersionInfo
	)
	db := pgxDB()
	pvs := security.NewPhoneVerifyService()

	i := newIndexPostsService(db)
	if cfg.If("SimpleCacheIndex") {
		i = newSimpleIndexPostsService(db)
		c, v = cache.NewSimpleCacheIndexService(i)
	} else if cfg.If("BigCacheIndex") {
		a := newAuthorizationManageService(db)
		c, v = cache.NewBigCacheIndexService(i, a)
	} else {
		c, v = cache.NewNoneCacheIndexService(i)
	}
	logrus.Infof("use %s as cache index service by version: %s", v.Name(), v.Version())

	ds := &dataServant{
		IndexPostsService:      c,
		WalletService:          newWalletService(db),
		MessageService:         newMessageService(db),
		TopicService:           newTopicService(db),
		TweetService:           newTweetService(db),
		TweetManageService:     newTweetManageService(db, c),
		TweetHelpService:       newTweetHelpService(db),
		CommentService:         newCommentService(db),
		CommentManageService:   newCommentManageService(db),
		UserManageService:      newUserManageService(db),
		ContactManageService:   newContactManageService(db),
		SecurityService:        newSecurityService(db, pvs),
		AttachmentCheckService: security.NewAttachmentCheckService(),
	}
	return ds, ds
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	return newAuthorizationManageService(pgxDB())
}

func (s *dataServant) Name() string {
	return "sqlc/pgx"
}

func (s *dataServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
