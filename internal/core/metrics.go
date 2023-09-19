// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

type TweetMetricServantA interface {
	UpdateRankScore(metric *cs.TweetMetric) error
	AddTweetMetric(postId int64) error
	DeleteTweetMetric(postId int64) error
}
