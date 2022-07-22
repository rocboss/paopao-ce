package dao

import (
	"sync"

	"github.com/rocboss/paopao-ce/internal/conf"
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
		if conf.CfgIf("Gorm") {
			ds, v = jinzhu.NewDataService()
		} else if conf.CfgIf("Sqlx") && conf.CfgIf("MySQL") {
			ds, v = sakila.NewDataService()
		} else if conf.CfgIf("Sqlx") && (conf.CfgIf("Postgres") || conf.CfgIf("PostgreSQL")) {
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
		if conf.CfgIf("AliOSS") {
			oss, v = storage.NewAliossService()
		} else if conf.CfgIf("MinIO") {
			oss, v = storage.NewMinioService()
		} else if conf.CfgIf("S3") {
			oss, v = storage.NewS3Service()
			logrus.Infof("use S3 as object storage by version %s", v.Version())
			return
		} else if conf.CfgIf("LocalOSS") {
			oss, v = storage.NewLocalossService()
		} else {
			// default use AliOSS as object storage service
			oss, v = storage.NewAliossService()
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
		if conf.CfgIf("Zinc") {
			ts, v = search.NewZincTweetSearchService(ams)
		} else if conf.CfgIf("Meili") {
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
	if conf.CfgIf("Gorm") {
		s = jinzhu.NewAuthorizationManageService()
	} else if conf.CfgIf("Sqlx") && conf.CfgIf("MySQL") {
		s = sakila.NewAuthorizationManageService()
	} else if conf.CfgIf("Sqlx") && (conf.CfgIf("Postgres") || conf.CfgIf("PostgreSQL")) {
		s = slonik.NewAuthorizationManageService()
	} else {
		s = jinzhu.NewAuthorizationManageService()
	}
	return
}
