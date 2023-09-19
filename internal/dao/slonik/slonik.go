// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/Masterminds/semver/v3"
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
	core.FollowingManageService
	core.SecurityService
	core.AttachmentCheckService
	core.TweetMetricServantA
}

type webDataSrvA struct {
	core.TopicServantA
	core.TweetServantA
	core.TweetManageServantA
	core.TweetHelpServantA
}

func NewDataService() (core.DataService, core.VersionInfo) {
	db := pgxDB()
	pvs := security.NewPhoneVerifyService()
	tms := NewTweetMetricServentA(db)
	cis := cache.NewEventCacheIndexSrv(tms)
	ds := &dataSrv{
		TweetMetricServantA:    tms,
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
	return cache.NewCacheDataService(ds), ds
}

func NewWebDataServantA() (core.WebDataServantA, core.VersionInfo) {
	logrus.Fatal("not support now")
	return nil, nil
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	return newAuthorizationManageService(pgxDB())
}

func (s *dataSrv) Name() string {
	return "sqlc/pgx"
}

func (s *dataSrv) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *webDataSrvA) Name() string {
	return "sqlc/pgx"
}

func (s *webDataSrvA) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
