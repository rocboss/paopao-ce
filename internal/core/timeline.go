// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

// IndexPostsService 广场首页推文列表服务
type IndexPostsService interface {
	IndexPosts(user *ms.User, offset int, limit int) (*ms.IndexTweetList, error)
}

// IndexPostsServantA 广场首页推文列表服务(版本A)
type IndexPostsServantA interface {
	IndexPosts(user *ms.User, limit int, offset int) (*cs.TweetBox, error)
}
