// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package search

import (
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/json"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"github.com/sirupsen/logrus"
)

var (
	_ core.TweetSearchService = (*zincTweetSearchServant)(nil)
	_ core.VersionInfo        = (*zincTweetSearchServant)(nil)
)

type zincTweetSearchServant struct {
	tweetSearchFilter

	indexName     string
	client        *zinc.ZincClient
	publicFilter  string
	privateFilter string
	friendFilter  string
}

func (s *zincTweetSearchServant) Name() string {
	return "Zinc"
}

func (s *zincTweetSearchServant) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}

func (s *zincTweetSearchServant) IndexName() string {
	return s.indexName
}

func (s *zincTweetSearchServant) AddDocuments(data []core.TsDocItem, primaryKey ...string) (bool, error) {
	if len(data) == 0 {
		return true, nil
	}

	buf := make([]map[string]any, 0, len(data)+1)
	if len(primaryKey) > 0 {
		buf = append(buf, map[string]any{
			"index": map[string]any{
				"_index": s.indexName,
				"_id":    primaryKey[0],
			},
		})
	} else {
		buf = append(buf, map[string]any{
			"index": map[string]any{
				"_index": s.indexName,
			},
		})
	}
	docs := s.toDocs(data)
	buf = append(buf, docs...)

	return s.client.BulkPushDoc(buf)
}

func (s *zincTweetSearchServant) DeleteDocuments(identifiers []string) error {
	for _, id := range identifiers {
		if err := s.client.DelDoc(s.indexName, id); err != nil {
			return err
		}
	}
	return nil
}

func (s *zincTweetSearchServant) Search(user *core.User, q *core.QueryReq, offset, limit int) (resp *core.QueryResp, err error) {
	if q.Type == core.SearchTypeDefault && q.Query != "" {
		resp, err = s.queryByContent(user, q, offset, limit)
	} else if q.Type == core.SearchTypeTag && q.Query != "" {
		resp, err = s.queryByTag(user, q, offset, limit)
	} else {
		resp, err = s.queryAny(user, offset, limit)
	}
	if err != nil {
		logrus.Errorf("zincTweetSearchServant.search searchType:%s query:%s error:%v", q.Type, q.Query, err)
		return
	}

	logrus.Debugf("zincTweetSearchServant.Search type:%s query:%s resp Hits:%d NbHits:%d offset: %d limit:%d ", q.Type, q.Query, len(resp.Items), resp.Total, offset, limit)
	s.filterResp(user, resp)
	return
}

func (s *zincTweetSearchServant) queryByContent(user *core.User, q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	resp, err := s.client.EsQuery(s.indexName, map[string]any{
		"query": map[string]any{
			"match_phrase": map[string]any{
				"content": q.Query,
			},
		},
		"sort": []string{"-is_top", "-latest_replied_on"},
		"from": offset,
		"size": limit,
	})
	if err != nil {
		return nil, err
	}
	return s.postsFrom(resp)
}

func (s *zincTweetSearchServant) queryByTag(user *core.User, q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	resp, err := s.client.ApiQuery(s.indexName, map[string]any{
		"search_type": "querystring",
		"query": map[string]any{
			"term": "tags." + q.Query + ":1",
		},
		"sort_fields": []string{"-is_top", "-latest_replied_on"},
		"from":        offset,
		"max_results": limit,
	})
	if err != nil {
		return nil, err
	}
	return s.postsFrom(resp)
}

func (s *zincTweetSearchServant) queryAny(user *core.User, offset, limit int) (*core.QueryResp, error) {
	queryMap := map[string]any{
		"query": map[string]any{
			"match_all": map[string]string{},
		},
		"sort": []string{"-is_top", "-latest_replied_on"},
		"from": offset,
		"size": limit,
	}
	resp, err := s.client.EsQuery(s.indexName, queryMap)
	if err != nil {
		return nil, err
	}
	return s.postsFrom(resp)
}

func (s *zincTweetSearchServant) postsFrom(resp *zinc.QueryResultT) (*core.QueryResp, error) {
	posts := make([]*core.PostFormated, 0, len(resp.Hits.Hits))
	for _, hit := range resp.Hits.Hits {
		item := &core.PostFormated{}
		raw, err := json.Marshal(hit.Source)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(raw, item); err != nil {
			return nil, err
		}
		posts = append(posts, item)
	}

	return &core.QueryResp{
		Items: posts,
		Total: resp.Hits.Total.Value,
	}, nil
}

func (s *zincTweetSearchServant) createIndex() {
	// 不存在则创建索引
	s.client.CreateIndex(s.indexName, &zinc.ZincIndexProperty{
		"id": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Store:    true,
			Sortable: true,
		},
		"user_id": &zinc.ZincIndexPropertyT{
			Type:  "numeric",
			Index: true,
			Store: true,
		},
		"comment_count": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"collection_count": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"upvote_count": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"visibility": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"is_top": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"is_essence": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"content": &zinc.ZincIndexPropertyT{
			Type:           "text",
			Index:          true,
			Store:          true,
			Aggregatable:   true,
			Highlightable:  true,
			Analyzer:       "gse_search",
			SearchAnalyzer: "gse_standard",
		},
		"tags": &zinc.ZincIndexPropertyT{
			Type:  "keyword",
			Index: true,
			Store: true,
		},
		"ip_loc": &zinc.ZincIndexPropertyT{
			Type:  "keyword",
			Index: true,
			Store: true,
		},
		"latest_replied_on": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"attachment_price": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Sortable: true,
			Store:    true,
		},
		"created_on": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"modified_on": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
	})
}

func (s *zincTweetSearchServant) toDocs(data []core.TsDocItem) []map[string]any {
	docs := make([]map[string]any, 0, len(data))
	for _, d := range data {
		tagMaps := map[string]int8{}
		for _, tag := range strings.Split(d.Post.Tags, ",") {
			tagMaps[tag] = 1
		}
		docs = append(docs, map[string]any{
			"id":                d.Post.ID,
			"user_id":           d.Post.UserID,
			"comment_count":     d.Post.CommentCount,
			"collection_count":  d.Post.CollectionCount,
			"upvote_count":      d.Post.UpvoteCount,
			"visibility":        d.Post.Visibility,
			"is_top":            d.Post.IsTop,
			"is_essence":        d.Post.IsEssence,
			"content":           d.Content,
			"tags":              tagMaps,
			"ip_loc":            d.Post.IPLoc,
			"latest_replied_on": d.Post.LatestRepliedOn,
			"attachment_price":  d.Post.AttachmentPrice,
			"created_on":        d.Post.CreatedOn,
			"modified_on":       d.Post.ModifiedOn,
		})
	}
	return docs
}
