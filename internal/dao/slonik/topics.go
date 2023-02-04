// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/ce/postgres"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	*pgxServant
}

func (s *topicServant) UpsertTags(userId int64, tags []string) ([]*core.Tag, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *topicServant) DecrTagsById(ids []int64) error {
	// return s.with(func(c context.Context, q dbr.Querier) error {
	// 	tagIds, err := q.TagsByIdA(c, pgtype.Array[int64](ids))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	err := q.DecrTagsById(c, &dbr.DecrTagsByIdParams{ModifiedOn: time.Now().Unix(), Ids: tagIds})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })
	return debug.ErrNotImplemented
}

func (s *topicServant) GetTags(category core.TagCategory, offset int, limit int) (res []*core.Tag, _ error) {
	ctx := context.Background()
	switch category {
	case core.TagCategoryHot:
		tags, err := s.q.HotTags(ctx, &dbr.HotTagsParams{Offset: int32(offset), Limit: int32(limit)})
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			res = append(res, &core.Tag{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	case core.TagCategoryNew:
		tags, err := s.q.NewestTags(ctx, &dbr.NewestTagsParams{Offset: int32(offset), Limit: int32(limit)})
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			res = append(res, &core.Tag{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	}
	return
}

func (s *topicServant) GetTagsByKeyword(keyword string) (res []*core.Tag, _ error) {
	ctx := context.Background()
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		tags, err := s.q.TagsByKeywordA(ctx)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			res = append(res, &core.Tag{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	} else {
		tags, err := s.q.TagsByKeywordB(ctx, keyword)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			res = append(res, &core.Tag{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	}
	return
}

func newTopicService(db *pgx.Conn) core.TopicService {
	return &topicServant{
		pgxServant: newPgxServant(db),
	}
}
