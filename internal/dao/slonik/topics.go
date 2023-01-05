// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/ce/postgres"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	db *pgx.Conn
	q  dbr.Querier
}

func (s *topicServant) CreateTag(tag *core.Tag) (*core.Tag, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *topicServant) DeleteTag(tag *core.Tag) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *topicServant) GetTags(conditions *core.ConditionsT, offset, limit int) ([]*core.Tag, error) {
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
		db: db,
		q:  dbr.New(db),
	}
}
