package search

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"github.com/sirupsen/logrus"
)

func NewMeiliTweetSearchService(ams core.AuthorizationManageService) (core.TweetSearchService, core.VersionInfo) {
	s := conf.MeiliSetting
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   s.Endpoint(),
		APIKey: s.ApiKey,
	})

	if _, err := client.Index(s.Index).FetchInfo(); err != nil {
		logrus.Debugf("create index because fetch index info error: %v", err)
		client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        s.Index,
			PrimaryKey: "id",
		})
		searchableAttributes := []string{"content", "tags"}
		sortableAttributes := []string{"is_top", "latest_replied_on"}
		filterableAttributes := []string{"tags", "visibility", "user_id"}

		index := client.Index(s.Index)
		index.UpdateSearchableAttributes(&searchableAttributes)
		index.UpdateSortableAttributes(&sortableAttributes)
		index.UpdateFilterableAttributes(&filterableAttributes)
	}

	mts := &meiliTweetSearchServant{
		tweetSearchFilter: tweetSearchFilter{
			ams: ams,
		},
		client:        client,
		index:         client.Index(s.Index),
		publicFilter:  fmt.Sprintf("visibility=%d", model.PostVisitPublic),
		privateFilter: fmt.Sprintf("visibility=%d AND user_id=", model.PostVisitPrivate),
		friendFilter:  fmt.Sprintf("visibility=%d", model.PostVisitFriend),
	}
	return mts, mts
}

func NewZincTweetSearchService(ams core.AuthorizationManageService) (core.TweetSearchService, core.VersionInfo) {
	s := conf.ZincSetting
	zts := &zincTweetSearchServant{
		tweetSearchFilter: tweetSearchFilter{
			ams: ams,
		},
		indexName:     s.Index,
		client:        zinc.NewClient(s),
		publicFilter:  fmt.Sprintf("visibility:%d", model.PostVisitPublic),
		privateFilter: fmt.Sprintf("visibility:%d AND user_id:%%d", model.PostVisitPrivate),
		friendFilter:  fmt.Sprintf("visibility:%d", model.PostVisitFriend),
	}
	zts.createIndex()

	return zts, zts
}

func NewBridgeTweetSearchService(ts core.TweetSearchService) core.TweetSearchService {
	capacity := conf.TweetSearchSetting.MaxUpdateQPS
	if capacity < 10 {
		capacity = 10
	} else if capacity > 10000 {
		capacity = 10000
	}
	bts := &bridgeTweetSearchServant{
		ts:           ts,
		updateDocsCh: make(chan *documents, capacity),
	}

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
