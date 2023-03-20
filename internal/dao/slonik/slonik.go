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
	var (
		v   core.VersionInfo
		cis core.CacheIndexService
		ips core.IndexPostsService
	)
	db := pgxDB()
	pvs := security.NewPhoneVerifyService()
	ams := NewAuthorizationManageService()
	ths := newTweetHelpService(db)

	// initialize core.IndexPostsService
	if cfg.If("Friendship") {
		ips = newFriendIndexService(db, ams, ths)
	} else if cfg.If("Followship") {
		ips = newFollowIndexService(db, ths)
	} else if cfg.If("Lightship") {
		ips = newLightIndexService(db, ths)
	} else {
		// default use lightship post index service
		ips = newLightIndexService(db, ths)
	}

	// initialize core.CacheIndexService
	if cfg.If("SimpleCacheIndex") {
		// simpleCache use special post index service
		ips = newSimpleIndexPostsService(db, ths)
		cis, v = cache.NewSimpleCacheIndexService(ips)
	} else if cfg.If("BigCacheIndex") {
		// TODO: make cache index post in different scence like friendship/followship/lightship
		cis, v = cache.NewBigCacheIndexService(ips, ams)
	} else {
		cis, v = cache.NewNoneCacheIndexService(ips)
	}
	logrus.Infof("use %s as cache index service by version: %s", v.Name(), v.Version())

	ds := &dataServant{
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
		SecurityService:        newSecurityService(db, pvs),
		AttachmentCheckService: security.NewAttachmentCheckService(),
	}
	return ds, ds
}

func NewDataServantA() (core.DataServantA, core.VersionInfo) {
	logrus.Fatal("not support now")
	return nil, nil
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
