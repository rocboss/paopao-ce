// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dao

import (
	"sync"

	"github.com/alimy/cfg"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu"
	"github.com/rocboss/paopao-ce/internal/dao/sakila"
	"github.com/rocboss/paopao-ce/internal/dao/search"
	"github.com/rocboss/paopao-ce/internal/dao/slonik"
	"github.com/rocboss/paopao-ce/internal/dao/storage"
	"github.com/sirupsen/logrus"
)

var (
	ts  core.TweetSearchService
	ds  core.DataService
	oss core.ObjectStorageService

	onceTs, onceDs, onceOss sync.Once
)

func DataService() core.DataService {
	onceDs.Do(func() {
		var v core.VersionInfo
		if cfg.If("Gorm") {
			ds, v = jinzhu.NewDataService()
		} else if cfg.If("Sqlx") && cfg.If("MySQL") {
			ds, v = sakila.NewDataService()
		} else if cfg.If("Sqlx") && (cfg.If("Postgres") || cfg.If("PostgreSQL")) {
			ds, v = slonik.NewDataService()
		} else {
			// default use gorm as orm for sql database
			ds, v = jinzhu.NewDataService()
		}
		logrus.Infof("use %s as data service with version %s", v.Name(), v.Version())
	})
	return ds
}

func ObjectStorageService() core.ObjectStorageService {
	onceOss.Do(func() {
		var v core.VersionInfo
		if cfg.If("AliOSS") {
			oss, v = storage.MustAliossService()
		} else if cfg.If("COS") {
			oss, v = storage.NewCosService()
		} else if cfg.If("HuaweiOBS") {
			oss, v = storage.MustHuaweiobsService()
		} else if cfg.If("MinIO") {
			oss, v = storage.MustMinioService()
		} else if cfg.If("S3") {
			oss, v = storage.MustS3Service()
			logrus.Infof("use S3 as object storage by version %s", v.Version())
			return
		} else if cfg.If("LocalOSS") {
			oss, v = storage.MustLocalossService()
		} else {
			// default use AliOSS as object storage service
			oss, v = storage.MustAliossService()
			logrus.Infof("use default AliOSS as object storage by version %s", v.Version())
			return
		}
		logrus.Infof("use %s as object storage by version %s", v.Name(), v.Version())
	})
	return oss
}

func TweetSearchService() core.TweetSearchService {
	onceTs.Do(func() {
		var v core.VersionInfo
		ams := newAuthorizationManageService()
		if cfg.If("Zinc") {
			ts, v = search.NewZincTweetSearchService(ams)
		} else if cfg.If("Meili") {
			ts, v = search.NewMeiliTweetSearchService(ams)
		} else {
			// default use Zinc as tweet search service
			ts, v = search.NewZincTweetSearchService(ams)
		}
		logrus.Infof("use %s as tweet search serice by version %s", v.Name(), v.Version())

		ts = search.NewBridgeTweetSearchService(ts)
	})
	return ts
}

func newAuthorizationManageService() (s core.AuthorizationManageService) {
	if cfg.If("Gorm") {
		s = jinzhu.NewAuthorizationManageService()
	} else if cfg.If("Sqlx") && cfg.If("MySQL") {
		s = sakila.NewAuthorizationManageService()
	} else if cfg.If("Sqlx") && (cfg.If("Postgres") || cfg.If("PostgreSQL")) {
		s = slonik.NewAuthorizationManageService()
	} else {
		s = jinzhu.NewAuthorizationManageService()
	}
	return
}
