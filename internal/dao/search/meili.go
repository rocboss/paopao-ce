// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package search

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/json"
	"github.com/sirupsen/logrus"
)

var (
	_ core.TweetSearchService = (*meiliTweetSearchServant)(nil)
	_ core.VersionInfo        = (*meiliTweetSearchServant)(nil)
)

type meiliTweetSearchServant struct {
	tweetSearchFilter

	client        *meilisearch.Client
	index         *meilisearch.Index
	publicFilter  string
	privateFilter string
	friendFilter  string
}

type postInfo struct {
	ID              int64             `json:"id"`
	UserID          int64             `json:"user_id"`
	CommentCount    int64             `json:"comment_count"`
	CollectionCount int64             `json:"collection_count"`
	UpvoteCount     int64             `json:"upvote_count"`
	Visibility      core.PostVisibleT `json:"visibility"`
	IsTop           int               `json:"is_top"`
	IsEssence       int               `json:"is_essence"`
	IsLock          int               `json:"is_lock"`
	LatestRepliedOn int64             `json:"latest_replied_on"`
	CreatedOn       int64             `json:"created_on"`
	ModifiedOn      int64             `json:"modified_on"`
	AttachmentPrice int64             `json:"attachment_price"`
	IPLoc           string            `json:"ip_loc"`
}

func (s *meiliTweetSearchServant) Name() string {
	return "Meili"
}

func (s *meiliTweetSearchServant) Version() *semver.Version {
	return semver.MustParse("v0.2.2")
}

func (s *meiliTweetSearchServant) IndexName() string {
	return s.index.UID
}

func (s *meiliTweetSearchServant) AddDocuments(data []core.TsDocItem, _primaryKey ...string) (bool, error) {
	docs := s.toDocs(data)
	if len(docs) == 0 {
		return true, nil
	}
	if _, err := s.index.AddDocuments(docs); err != nil {
		logrus.Errorf("meiliTweetSearchServant.AddDocuments error: %s", err)
		return false, err
	}
	return true, nil
}

func (s *meiliTweetSearchServant) DeleteDocuments(identifiers []string) error {
	task, err := s.index.DeleteDocuments(identifiers)
	if err != nil {
		logrus.Errorf("meiliTweetSearchServant.DeleteDocuments error: %v", err)
		return err
	}
	logrus.Debugf("meiliTweetSearchServant.DeleteDocuments task: (taskUID:%d, indexUID:%s, status:%s)", task.TaskUID, task.IndexUID, task.Status)
	return nil
}

func (s *meiliTweetSearchServant) Search(user *core.User, q *core.QueryReq, offset, limit int) (resp *core.QueryResp, err error) {
	if q.Type == core.SearchTypeDefault && q.Query != "" {
		resp, err = s.queryByContent(user, q, offset, limit)
	} else if q.Type == core.SearchTypeTag && q.Query != "" {
		resp, err = s.queryByTag(user, q, offset, limit)
	} else {
		resp, err = s.queryAny(user, offset, limit)
	}
	if err != nil {
		logrus.Errorf("meiliTweetSearchServant.search searchType:%s query:%s error:%v", q.Type, q.Query, err)
		return
	}

	logrus.Debugf("meiliTweetSearchServant.Search type:%s query:%s resp Hits:%d NbHits:%d offset: %d limit:%d ", q.Type, q.Query, len(resp.Items), resp.Total, offset, limit)
	s.filterResp(user, resp)
	return
}

func (s *meiliTweetSearchServant) queryByContent(user *core.User, q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	request := &meilisearch.SearchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Sort:   []string{"is_top:desc", "latest_replied_on:desc"},
	}

	filter := s.filterList(user)
	if len(filter) > 0 {
		request.Filter = filter
	}

	// logrus.Debugf("meiliTweetSearchServant.queryByContent query:%s request%+v", q.Query, request)
	resp, err := s.index.Search(q.Query, request)
	if err != nil {
		return nil, err
	}

	return s.postsFrom(resp)
}

func (s *meiliTweetSearchServant) queryByTag(user *core.User, q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	request := &meilisearch.SearchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Sort:   []string{"is_top:desc", "latest_replied_on:desc"},
	}

	filter := s.filterList(user)
	tagFilter := []string{"tags=" + q.Query}
	if len(filter) > 0 {
		request.Filter = [][]string{tagFilter, {filter}}
	} else {
		request.Filter = tagFilter
	}

	// logrus.Debugf("meiliTweetSearchServant.queryByTag query:%s request%+v", q.Query, request)
	resp, err := s.index.Search("#"+q.Query, request)
	if err != nil {
		return nil, err
	}

	return s.postsFrom(resp)
}

func (s *meiliTweetSearchServant) queryAny(user *core.User, offset, limit int) (*core.QueryResp, error) {
	request := &meilisearch.SearchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Sort:   []string{"is_top:desc", "latest_replied_on:desc"},
	}

	filter := s.filterList(user)
	if len(filter) > 0 {
		request.Filter = filter
	}

	resp, err := s.index.Search("", request)
	if err != nil {
		return nil, err
	}

	return s.postsFrom(resp)
}

func (s *meiliTweetSearchServant) filterList(user *core.User) string {
	if user == nil {
		return s.publicFilter
	}

	if user.IsAdmin {
		return ""
	}

	return fmt.Sprintf("%s OR %s OR (%s%d)", s.publicFilter, s.friendFilter, s.privateFilter, user.ID)
}

func (s *meiliTweetSearchServant) postsFrom(resp *meilisearch.SearchResponse) (*core.QueryResp, error) {
	posts := make([]*core.PostFormated, 0, len(resp.Hits))
	for _, hit := range resp.Hits {
		raw, err := json.Marshal(hit)
		if err != nil {
			return nil, err
		}
		p := &postInfo{}
		if err = json.Unmarshal(raw, p); err != nil {
			return nil, err
		}
		posts = append(posts, &core.PostFormated{
			ID:              p.ID,
			UserID:          p.UserID,
			CommentCount:    p.CommentCount,
			CollectionCount: p.CollectionCount,
			UpvoteCount:     p.UpvoteCount,
			Visibility:      p.Visibility,
			IsTop:           p.IsTop,
			IsEssence:       p.IsEssence,
			IsLock:          p.IsLock,
			LatestRepliedOn: p.LatestRepliedOn,
			CreatedOn:       p.CreatedOn,
			ModifiedOn:      p.ModifiedOn,
			AttachmentPrice: p.AttachmentPrice,
			IPLoc:           p.IPLoc,
		})
	}
	return &core.QueryResp{
		Items: posts,
		Total: resp.EstimatedTotalHits,
	}, nil
}

func (s *meiliTweetSearchServant) toDocs(data []core.TsDocItem) []map[string]any {
	docs := make([]map[string]any, 0, len(data))
	for _, d := range data {
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
			"tags":              strings.Split(d.Post.Tags, ","),
			"ip_loc":            d.Post.IPLoc,
			"latest_replied_on": d.Post.LatestRepliedOn,
			"attachment_price":  d.Post.AttachmentPrice,
			"created_on":        d.Post.CreatedOn,
			"modified_on":       d.Post.ModifiedOn,
		})
	}
	return docs
}
