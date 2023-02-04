// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
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
	debug.NotImplemented()
	return nil, nil
}

func (s *topicServant) DecrTagsById(ids []int64) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *topicServant) GetTags(category core.TagCategory, offset int, limit int) ([]*core.Tag, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *topicServant) GetTagsByKeyword(keyword string) ([]*core.Tag, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func newTopicService(db *pgx.Conn) core.TopicService {
	return &topicServant{
		pgxServant: newPgxServant(db),
	}
}
