// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/dao/security"
)

var (
	_ core.DataService = (*dataSrv)(nil)
	_ core.VersionInfo = (*dataSrv)(nil)

	_ core.WebDataServantA = (*webDataSrvA)(nil)
	_ core.VersionInfo     = (*webDataSrvA)(nil)

	_onceInitial sync.Once
)

type dataSrv struct {
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
	core.UserRelationService
	core.FollowingManageService
	core.SecurityService
	core.AttachmentCheckService
	core.TweetMetricServantA
	core.CommentMetricServantA
}

type webDataSrvA struct {
	core.TopicServantA
	core.TweetServantA
	core.TweetManageServantA
	core.TweetHelpServantA
}

func NewDataService() (core.DataService, core.VersionInfo) {
	lazyInitial()
	pvs := security.NewPhoneVerifyService()
	cms := newCommentMetricServentA(_db)
	tms := newTweetMetricServentA(_db)
	cis := cache.NewEventCacheIndexSrv(tms)
	ds := &dataSrv{
		TweetMetricServantA:    tms,
		CommentMetricServantA:  cms,
		WalletService:          newWalletService(_db),
		MessageService:         newMessageService(_db),
		TopicService:           newTopicService(_db),
		TweetService:           newTweetService(_db),
		TweetManageService:     newTweetManageService(_db, cis),
		TweetHelpService:       newTweetHelpService(_db),
		CommentService:         newCommentService(_db),
		CommentManageService:   newCommentManageService(_db),
		UserManageService:      newUserManageService(_db),
		ContactManageService:   newContactManageService(_db),
		FollowingManageService: newFollowingManageService(_db),
		SecurityService:        newSecurityService(_db, pvs),
		AttachmentCheckService: security.NewAttachmentCheckService(),
	}
	return cache.NewCacheDataService(ds), ds
}

func NewWebDataServantA() (core.WebDataServantA, core.VersionInfo) {
	lazyInitial()
	tms := newTweetMetricServentA(_db)
	cis := cache.NewEventCacheIndexSrv(tms)
	ds := &webDataSrvA{
		TopicServantA:       newTopicServantA(_db),
		TweetServantA:       newTweetServantA(_db),
		TweetManageServantA: newTweetManageServantA(_db, cis),
		TweetHelpServantA:   newTweetHelpServantA(_db),
	}
	return ds, ds
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	lazyInitial()
	return newAuthorizationManageService(_db)
}

func (s *dataSrv) Name() string {
	return "Sqlx"
}

func (s *dataSrv) Version() *semver.Version {
	return semver.MustParse("v0.3.0")
}

func (s *webDataSrvA) Name() string {
	return "Sqlx"
}

func (s *webDataSrvA) Version() *semver.Version {
	return semver.MustParse("v0.0.1")
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		initSqlxDB()
	})
}
