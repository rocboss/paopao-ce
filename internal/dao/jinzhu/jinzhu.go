// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// package jinzhu Core service implement base gorm+mysql/postgresql/sqlite3.
// Jinzhu is the primary developer of gorm so use his name as
// package name as a saluter.

package jinzhu

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/alimy/tryst/cfg"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/dao/security"
	"github.com/sirupsen/logrus"
)

var (
	_ core.DataService = (*dataSrv)(nil)
	_ core.VersionInfo = (*dataSrv)(nil)

	_ core.WebDataServantA = (*webDataSrvA)(nil)
	_ core.VersionInfo     = (*webDataSrvA)(nil)

	_onceInitial sync.Once
)

type dataSrv struct {
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
	core.FollowingManageService
	core.SecurityService
	core.AttachmentCheckService
}

type webDataSrvA struct {
	core.TopicServantA
	core.TweetServantA
	core.TweetManageServantA
	core.TweetHelpServantA
}

func NewDataService() (core.DataService, core.VersionInfo) {
	lazyInitial()

	var (
		v   core.VersionInfo
		cis core.CacheIndexService
	)
	db := conf.MustGormDB()
	pvs := security.NewPhoneVerifyService()
	ams := NewAuthorizationManageService()
	ths := newTweetHelpService(db)
	ips := newShipIndexService(db, ams, ths)

	// initialize core.CacheIndexService
	cfg.On(cfg.Actions{
		"SimpleCacheIndex": func() {
			// simpleCache use special post index service
			ips = newSimpleIndexPostsService(db, ths)
			cis, v = cache.NewSimpleCacheIndexService(ips)
		},
		"BigCacheIndex": func() {
			cis, v = cache.NewBigCacheIndexService(ips, ams)
		},
		"RedisCacheIndex": func() {
			cis, v = cache.NewRedisCacheIndexService(ips, ams)
		},
	}, func() {
		// defualt no cache
		cis, v = cache.NewNoneCacheIndexService(ips)
	})
	logrus.Infof("use %s as cache index service by version: %s", v.Name(), v.Version())

	ds := &dataSrv{
		IndexPostsService:      cis,
		WalletService:          newWalletService(db),
		MessageService:         newMessageService(db),
		TopicService:           newTopicService(db),
		TweetService:           newTweetService(db),
		TweetManageService:     newTweetManageService(db, cis),
		TweetHelpService:       newTweetHelpService(db),
		CommentService:         newCommentService(db),
		CommentManageService:   newCommentManageService(db),
		UserManageService:      newUserManageService(db),
		ContactManageService:   newContactManageService(db),
		FollowingManageService: newFollowingManageService(db),
		SecurityService:        newSecurityService(db, pvs),
		AttachmentCheckService: security.NewAttachmentCheckService(),
	}
	return ds, ds
}

func NewWebDataServantA() (core.WebDataServantA, core.VersionInfo) {
	lazyInitial()
	db := conf.MustGormDB()
	ds := &webDataSrvA{
		TopicServantA:       newTopicServantA(db),
		TweetServantA:       newTweetServantA(db),
		TweetManageServantA: newTweetManageServantA(db),
		TweetHelpServantA:   newTweetHelpServantA(db),
	}
	return ds, ds
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	return newAuthorizationManageService(conf.MustGormDB())
}

func (s *dataSrv) Name() string {
	return "Gorm"
}

func (s *dataSrv) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}

func (s *webDataSrvA) Name() string {
	return "Gorm"
}

func (s *webDataSrvA) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		initTableName()
	})
}
