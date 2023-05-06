// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

type (
	Tag         = dbr.Tag
	TagFormated = dbr.TagFormated
)

// TopicService 话题服务
type TopicService interface {
	CreateTag(tag *Tag) (*Tag, error)
	DeleteTag(tag *Tag) error
	GetTags(conditions *ConditionsT, offset, limit int) ([]*Tag, error)
	GetHotTags(userId int64, limit int, offset int) ([]*TagFormated, error)
	GetNewestTags(userId int64, limit int, offset int) ([]*TagFormated, error)
	GetFollowTags(userId int64, limit int, offset int) ([]*TagFormated, error)
	GetTagsByKeyword(keyword string) ([]*Tag, error)
	FollowTopic(userId int64, topicId int64) error
	UnfollowTopic(userId int64, topicId int64) error
	StickTopic(userId int64, topicId int64) (int8, error)
}
