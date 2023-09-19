// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package cs contain core data service interface type
// model define

package cs

type TweetMetric struct {
	PostId          int64
	CommentCount    int64
	UpvoteCount     int64
	CollectionCount int64
	ShareCount      int64
	ThumbdownCount  int64
	ThumbupCount    int64
}

func (m *TweetMetric) RankScore(motivationFactor int) int64 {
	if motivationFactor == 0 {
		motivationFactor = 1
	}
	return (m.CommentCount + m.UpvoteCount*2 + m.CollectionCount*4 + m.ShareCount*8) * int64(motivationFactor)
}
