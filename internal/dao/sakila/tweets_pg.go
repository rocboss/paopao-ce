// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pg"
)

var (
	_ core.TweetManageService = (*pgTweetManageSrv)(nil)
)

type pgTweetManageSrv struct {
	*tweetManageSrv
	p *pg.TweetManage
}

func (s *pgTweetManageSrv) CreatePost(r *ms.Post) (*ms.Post, error) {
	now := time.Now().Unix()
	r.Model = &dbr.Model{CreatedOn: now}
	r.LatestRepliedOn = now
	err := s.p.AddPost.Get(&r.ID, r)
	if err != nil {
		return nil, err
	}
	s.cis.SendAction(core.IdxActCreatePost, r)
	return r, nil
}

func (s *pgTweetManageSrv) CreatePostContent(r *ms.PostContent) (*ms.PostContent, error) {
	r.Model = &ms.Model{CreatedOn: time.Now().Unix()}
	err := s.p.AddPostContent.Get(&r.ID, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
