// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	db          *sqlx.DB
	stmtAddTag  *sqlx.Stmt
	stmtDelTag  *sqlx.Stmt
	stmtListTag *sqlx.Stmt
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

func newTopicService(db *sqlx.DB) core.TopicService {
	return &topicServant{
		db:          db,
		stmtAddTag:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtDelTag:  c(`SELECT * FROM @person WHERE first_name=?`),
		stmtListTag: c(`SELECT * FROM @person WHERE first_name=?`),
	}
}
