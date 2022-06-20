package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"github.com/sirupsen/logrus"
)

var (
	_ core.TweetSearchService = (*zincTweetSearchServant)(nil)
)

type zincTweetSearchServant struct {
	indexName string
	client    *zinc.ZincClient
}

func NewTweetSearchService() (ts core.TweetSearchService) {
	if conf.CfgIf("Zinc") {
		ts = newZincTweetSearchServant()
	} else {
		// default use Zinc as tweet search service
		ts = newZincTweetSearchServant()
	}
	logrus.Infof("use %s as tweet search serice by version %s", ts.Name(), ts.Version())
	return
}
