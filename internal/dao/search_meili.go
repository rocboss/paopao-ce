package dao

import (
	"github.com/Masterminds/semver/v3"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/json"
	"github.com/sirupsen/logrus"
)

func newMeiliTweetSearchServant() *meiliTweetSearchServant {
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

		index := client.Index(s.Index)
		index.UpdateSearchableAttributes(&searchableAttributes)
		index.UpdateSortableAttributes(&sortableAttributes)
	}

	return &meiliTweetSearchServant{
		client: client,
		index:  client.Index(s.Index),
	}
}

func (s *meiliTweetSearchServant) Name() string {
	return "Meili"
}

func (s *meiliTweetSearchServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *meiliTweetSearchServant) IndexName() string {
	return s.index.UID
}

func (s *meiliTweetSearchServant) AddDocuments(data core.DocItems, primaryKey ...string) (bool, error) {
	if _, err := s.index.AddDocuments(data, primaryKey...); err != nil {
		logrus.Errorf("meiliTweetSearchServant.AddDocuments error: %v", err)
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
	logrus.Debugf("meiliTweetSearchServant.DeleteDocuments task: %+v", task.Details)
	return nil
}

func (s *meiliTweetSearchServant) Search(q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	if q.Type == core.SearchTypeDefault && q.Query != "" {
		return s.queryByContent(q, offset, limit)
	} else if q.Type == core.SearchTypeTag && q.Query != "" {
		return s.queryByTag(q, offset, limit)
	}
	return s.queryAny(offset, limit)
}

func (s *meiliTweetSearchServant) queryByContent(q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	resp, err := s.index.Search(q.Query, &meilisearch.SearchRequest{
		Offset:               int64(offset),
		Limit:                int64(limit),
		AttributesToRetrieve: []string{"*"},
		Sort:                 []string{"is_top:desc", "latest_replied_on:desc"},
	})
	if err != nil {
		return nil, err
	}
	return s.postsFrom(resp)
}

func (s *meiliTweetSearchServant) queryByTag(q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	resp, err := s.index.Search("#"+q.Query, &meilisearch.SearchRequest{
		Offset:               int64(offset),
		Limit:                int64(limit),
		AttributesToRetrieve: []string{"*"},
		Sort:                 []string{"is_top:desc", "latest_replied_on:desc"},
	})
	if err != nil {
		return nil, err
	}
	return s.postsFrom(resp)
}

func (s *meiliTweetSearchServant) queryAny(offset, limit int) (*core.QueryResp, error) {
	resp, err := s.index.Search("", &meilisearch.SearchRequest{
		Offset:  int64(offset),
		Limit:   int64(limit),
		Matches: true,
		Sort:    []string{"is_top:desc", "latest_replied_on:desc"},
	})
	if err != nil {
		return nil, err
	}
	return s.postsFrom(resp)
}

func (s *meiliTweetSearchServant) postsFrom(resp *meilisearch.SearchResponse) (*core.QueryResp, error) {
	logrus.Debugf("resp Hits:%d NbHits:%d offset: %d limit:%d ", len(resp.Hits), resp.NbHits, resp.Offset, resp.Limit)
	posts := make([]*model.PostFormated, 0, len(resp.Hits))
	for _, hit := range resp.Hits {
		item := &model.PostFormated{}
		raw, err := json.Marshal(hit)
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
		Total: resp.NbHits,
	}, nil
}
