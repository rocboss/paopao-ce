// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
)

var (
	_ core.TweetManageService = (*pgcTweetManageSrv)(nil)
)

type pgcTweetManageSrv struct {
	*tweetManageSrv
	p *pgc.TweetManage
}

func (s *pgcTweetManageSrv) CreatePost(r *ms.Post) (*ms.Post, error) {
	now := time.Now().Unix()
	r.Model = &ms.Model{CreatedOn: now}
	r.LatestRepliedOn = now
	err := s.p.AddPost.Get(&r.ID, r)
	if err != nil {
		return nil, err
	}
	s.cis.SendAction(core.IdxActCreatePost, r)
	return r, nil
}

func (s *pgcTweetManageSrv) CreatePostContent(r *ms.PostContent) (*ms.PostContent, error) {
	r.Model = &ms.Model{CreatedOn: time.Now().Unix()}
	if err := s.p.AddPostContent.Get(&r.ID, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (s *pgcTweetManageSrv) CreateAttachment(r *ms.Attachment) (res int64, err error) {
	args := []any{r.UserID, r.FileSize, r.ImgWidth, r.ImgHeight, r.Type, r.Content, time.Now().Unix()}
	err = s.p.AddAttachment.Get(&res, args...)
	return
}

func (s *pgcTweetManageSrv) CreatePostStar(postID, userID int64) (res *ms.PostStar, err error) {
	err = stmtGet(s.p.AddPostStar, &res, postID, userID, time.Now().Unix())
	return
}

func (s *pgcTweetManageSrv) CreatePostCollection(postID, userID int64) (res *ms.PostCollection, err error) {
	err = stmtGet(s.p.AddPostCollection, &res, postID, userID, time.Now().Unix())
	return
}
