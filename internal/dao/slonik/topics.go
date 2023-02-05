// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/ce/postgres"
	"github.com/rocboss/paopao-ce/pkg/types"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	*pgxServant
}

// UpsertTags update/insert tags info.
// Assume tags slice is distinct elements.
func (s *topicServant) UpsertTags(userId int64, tags []string) (res []*core.Tag, err error) {
	err = s.with(func(c context.Context, q dbr.Querier) error {
		now := time.Now().Unix()
		upTags, err := q.IncrTags(c, &dbr.IncrTagsParams{
			Tags:       types.PgxArray(tags),
			ModifiedOn: now,
		})
		if err != nil {
			return err
		}
		if len(upTags) > 0 {
			for _, t := range upTags {
				for i := 0; i < len(tags); {
					if tags[i] == t.Tag {
						latestIdx := len(tags) - 1
						tags[i] = tags[latestIdx]
						tags = tags[:latestIdx]
						break
					}
				}
				res = append(res, &core.Tag{
					ID:       t.ID,
					UserID:   t.UserID,
					Tag:      t.Tag,
					QuoteNum: t.QuoteNum,
				})
			}
		}
		for _, tag := range tags {
			id, err := q.InsertTags(c, &dbr.InsertTagsParams{
				UserID:    userId,
				Tag:       tag,
				CreatedOn: now,
			})
			if err != nil {
				return err
			}
			res = append(res, &core.Tag{
				ID:       id,
				UserID:   userId,
				Tag:      tag,
				QuoteNum: 1,
			})
		}
		return nil
	})
	return
}

func (s *topicServant) DecrTagsById(ids []int64) error {
	return s.q.DecrTagsById(context.Background(), &dbr.DecrTagsByIdParams{
		Ids:        types.PgxArray(ids),
		ModifiedOn: time.Now().Unix(),
	})
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
