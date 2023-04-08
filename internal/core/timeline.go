// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

type IndexTweetList struct {
	Tweets []*PostFormated
	Total  int64
}

// IndexPostsService 广场首页推文列表服务
type IndexPostsService interface {
	IndexPosts(user *User, offset int, limit int) (*IndexTweetList, error)
}

// IndexPostsServantA 广场首页推文列表服务(版本A)
type IndexPostsServantA interface {
	IndexPosts(user *User, limit int, offset int) (*cs.TweetBox, error)
}
