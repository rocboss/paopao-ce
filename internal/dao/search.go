package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"github.com/sirupsen/logrus"
)

var (
	_ core.TweetSearchService = (*zincTweetSearchServant)(nil)
	_ core.TweetSearchService = (*bridgeTweetSearchServant)(nil)
)

type documents struct {
	primaryKey  []string
	docItems    core.DocItems
	identifiers []string
}

type bridgeTweetSearchServant struct {
	ts           core.TweetSearchService
	updateDocsCh chan *documents
}

type zincTweetSearchServant struct {
	indexName string
	client    *zinc.ZincClient
}

func NewTweetSearchService() core.TweetSearchService {
	bts := &bridgeTweetSearchServant{}

	capacity := conf.TweetSearchSetting.MaxUpdateQPS
	if capacity < 10 {
		capacity = 10
	} else if capacity > 10000 {
		capacity = 10000
	}
	bts.updateDocsCh = make(chan *documents, capacity)

	if conf.CfgIf("Zinc") {
		bts.ts = newZincTweetSearchServant()
	} else {
		// default use Zinc as tweet search service
		bts.ts = newZincTweetSearchServant()
	}
	logrus.Infof("use %s as tweet search serice by version %s", bts.ts.Name(), bts.ts.Version())

	// 启动文档更新器
	go bts.startUpdateDocs()

	return bts
}
