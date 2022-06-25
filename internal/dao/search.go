package dao

import (
	"github.com/meilisearch/meilisearch-go"
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

type tweetSearchFilter struct {
	ams core.AuthorizationManageService
}

type zincTweetSearchServant struct {
	tweetSearchFilter

	indexName     string
	client        *zinc.ZincClient
	publicFilter  string
	privateFilter string
	friendFilter  string
}

type meiliTweetSearchServant struct {
	tweetSearchFilter

	client        *meilisearch.Client
	index         *meilisearch.Index
	publicFilter  string
	privateFilter string
	friendFilter  string
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
	} else if conf.CfgIf("Meili") {
		bts.ts = newMeiliTweetSearchServant()
	} else {
		// default use Zinc as tweet search service
		bts.ts = newZincTweetSearchServant()
	}
	logrus.Infof("use %s as tweet search serice by version %s", bts.ts.Name(), bts.ts.Version())

	numWorker := conf.TweetSearchSetting.MinWorker
	if numWorker < 5 {
		numWorker = 5
	} else if numWorker > 1000 {
		numWorker = 1000
	}
	logrus.Debugf("use %d backend worker to update documents to search engine", numWorker)
	// 启动文档更新器
	for ; numWorker > 0; numWorker-- {
		go bts.startUpdateDocs()
	}

	return bts
}
