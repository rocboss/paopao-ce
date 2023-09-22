// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/alimy/tryst/event"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type updateCommentMetricEvent struct {
	event.UnimplementedEvent
	db        *gorm.DB
	cms       core.CommentMetricServantA
	commentId int64
}

func onUpdateCommentMetricEvent(commentId int64, db *gorm.DB, cms core.CommentMetricServantA) {
	events.OnEvent(&updateCommentMetricEvent{
		db:        db,
		cms:       cms,
		commentId: commentId,
	})
}

func (e *updateCommentMetricEvent) Name() string {
	return "updateCommentMetricEvent"
}

func (e *updateCommentMetricEvent) Action() (err error) {
	logrus.Debugf("trigger updateCommentMetricEvent action commentId[%d]", e.commentId)
	comment := dbr.Comment{}
	if err = e.db.Table(_comment_).Where("id=?", e.commentId).First(&comment).Error; err != nil {
		return
	}
	e.cms.UpdateCommentMetric(&cs.CommentMetric{
		CommentId:       e.commentId,
		ReplyCount:      comment.ReplyCount,
		ThumbsUpCount:   comment.ThumbsUpCount,
		ThumbsDownCount: comment.ThumbsDownCount,
	})
	return
}
