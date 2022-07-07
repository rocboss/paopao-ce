package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu"
	"github.com/rocboss/paopao-ce/internal/dao/sakila"
	"github.com/rocboss/paopao-ce/internal/dao/search"
	"github.com/rocboss/paopao-ce/internal/dao/slonik"
	"github.com/rocboss/paopao-ce/internal/dao/storage"
	"github.com/sirupsen/logrus"
)

func NewDataService() (s core.DataService) {
	var v core.VersionInfo
	if conf.CfgIf("Gorm") {
		s, v = jinzhu.NewDataService()
	} else if conf.CfgIf("Sqlx") && conf.CfgIf("MySQL") {
		s, v = sakila.NewDataService()
	} else if conf.CfgIf("Sqlx") && (conf.CfgIf("Postgres") || conf.CfgIf("PostgreSQL")) {
		s, v = slonik.NewDataService()
	} else {
		// default use gorm as orm for sql database
		s, v = jinzhu.NewDataService()
	}
	logrus.Infof("use %s as data service with version %s", v.Name(), v.Version())
	return
}

func NewObjectStorageService() (oss core.ObjectStorageService) {
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
	return
}

func NewTweetSearchService() core.TweetSearchService {
	var (
		ts core.TweetSearchService
		v  core.VersionInfo
	)
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

	return search.NewBridgeTweetSearchService(ts)
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
