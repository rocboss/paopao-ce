// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

// TweetTimelineService 广场首页推文时间线服务
type TweetTimelineService interface {
	TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error)
}

// TweetTimelineServantA 广场首页推文时间线服务(版本A)
type TweetTimelineServantA interface {
	TweetTimeline(userId int64, offset int, limit int) (*cs.TweetBox, error)
}
