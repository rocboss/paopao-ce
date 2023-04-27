// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

type CommentThumbs struct {
	UserID       int64 `json:"user_id"`
	TweetID      int64 `json:"tweet_id"`
	CommentID    int64 `json:"comment_id"`
	ReplyID      int64 `json:"reply_id"`
	CommentType  int8  `json:"comment_type"`
	IsThumbsUp   int8  `json:"is_thumbs_up"`
	IsThumbsDown int8  `json:"is_thumbs_down"`
}

type CommentThumbsList []*CommentThumbs

type CommentThumbsMap map[int64]*CommentThumbs
