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
	"github.com/rocboss/paopao-ce/internal/core/cs"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/sqlc/postgres"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.TopicService = (*topicSrv)(nil)
)

type topicSrv struct {
	*pgxSrv
}

// UpsertTags update/insert tags info.
// Assume tags slice is distinct elements.
func (s *topicSrv) UpsertTags(userId int64, tags []string) (res cs.TagInfoList, err error) {
	err = s.with(func(c context.Context, q dbr.Querier) error {
		now := time.Now().Unix()
		upTags, err := q.IncrTags(c, &dbr.IncrTagsParams{
			Tags:       tags,
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
				res = append(res, &cs.TagInfo{
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
			res = append(res, &cs.TagInfo{
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

func (s *topicSrv) DecrTagsById(ids []int64) error {
	return s.q.DecrTagsById(context.Background(), &dbr.DecrTagsByIdParams{
		Ids:        ids,
		ModifiedOn: time.Now().Unix(),
	})
}

func (s *topicSrv) ListTags(typ cs.TagType, limit int, offset int) (res cs.TagList, _ error) {
	ctx := context.Background()
	switch typ {
	case cs.TagTypeHot:
		tags, err := s.q.HotTags(ctx, &dbr.HotTagsParams{Limit: int32(limit), Offset: int32(offset)})
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			res = append(res, &cs.TagItem{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
				User: &cs.UserInfo{
					ID:       tag.UserID,
					Nickname: tag.Nickname,
					Username: tag.Username,
					Status:   int(tag.Status),
					Avatar:   tag.Avatar,
					IsAdmin:  tag.IsAdmin,
				},
			})
		}
	case cs.TagTypeNew:
		tags, err := s.q.NewestTags(ctx, &dbr.NewestTagsParams{Limit: int32(limit), Offset: int32(offset)})
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			res = append(res, &cs.TagItem{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
				User: &cs.UserInfo{
					ID:       tag.UserID,
					Nickname: tag.Nickname,
					Username: tag.Username,
					Status:   int(tag.Status),
					Avatar:   tag.Avatar,
					IsAdmin:  tag.IsAdmin,
				},
			})
		}
	}
	return
}

func (s *topicSrv) TagsByKeyword(keyword string) (res cs.TagInfoList, _ error) {
	ctx := context.Background()
	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		tags, err := s.q.TagsByKeywordA(ctx)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			res = append(res, &cs.TagInfo{
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
			res = append(res, &cs.TagInfo{
				ID:       tag.ID,
				UserID:   tag.UserID,
				Tag:      tag.Tag,
				QuoteNum: tag.QuoteNum,
			})
		}
	}
	return
}

func (s *topicSrv) GetHotTags(userId int64, limit int, offset int) (cs.TagList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *topicSrv) GetNewestTags(userId int64, limit int, offset int) (cs.TagList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *topicSrv) GetFollowTags(userId int64, limit int, offset int) (cs.TagList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *topicSrv) FollowTopic(userId int64, topicId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *topicSrv) UnfollowTopic(userId int64, topicId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *topicSrv) StickTopic(userId int64, topicId int64) (int8, error) {
	// TODO
	return 0, debug.ErrNotImplemented
}

func newTopicService(db *pgx.Conn) core.TopicService {
	return &topicSrv{
		pgxSrv: newPgxSrv(db),
	}
}
