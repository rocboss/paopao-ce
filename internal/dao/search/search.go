// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package search

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
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
		logrus.Debugf("create meili index because fetch index info error: %v", err)
		if _, err := client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        s.Index,
			PrimaryKey: "id",
		}); err == nil {
			settings := meilisearch.Settings{
				SearchableAttributes: []string{"content", "tags"},
				SortableAttributes:   []string{"is_top", "latest_replied_on"},
				FilterableAttributes: []string{"tags", "visibility", "user_id"},
			}
			if _, err = client.Index(s.Index).UpdateSettings(&settings); err != nil {
				logrus.Errorf("update meili settings error: %s", err)
			}
		} else {
			logrus.Errorf("create meili index error: %s", err)
		}
	}

	mts := &meiliTweetSearchServant{
		tweetSearchFilter: tweetSearchFilter{
			ams: ams,
		},
		client:        client,
		index:         client.Index(s.Index),
		publicFilter:  fmt.Sprintf("visibility=%d", core.PostVisitPublic),
		privateFilter: fmt.Sprintf("visibility=%d AND user_id=", core.PostVisitPrivate),
		friendFilter:  fmt.Sprintf("visibility=%d", core.PostVisitFriend),
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
		client:        zinc.NewClient(s.Endpoint(), s.User, s.Password),
		publicFilter:  fmt.Sprintf("visibility:%d", core.PostVisitPublic),
		privateFilter: fmt.Sprintf("visibility:%d AND user_id:%%d", core.PostVisitPrivate),
		friendFilter:  fmt.Sprintf("visibility:%d", core.PostVisitFriend),
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
		ts:               ts,
		updateDocsCh:     make(chan *documents, capacity),
		updateDocsTempCh: make(chan *documents, 100),
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
