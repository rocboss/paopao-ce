// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

// TopicService 话题服务
type TopicService interface {
	UpsertTags(userId int64, tags []string) (cs.TagInfoList, error)
	DecrTagsById(ids []int64) error
	ListTags(typ cs.TagType, limit int, offset int) (cs.TagList, error)
	TagsByKeyword(keyword string) (cs.TagInfoList, error)
	GetHotTags(userId int64, limit int, offset int) (cs.TagList, error)
	GetNewestTags(userId int64, limit int, offset int) (cs.TagList, error)
	GetFollowTags(userId int64, is_pin bool, limit int, offset int) (cs.TagList, error)
	FollowTopic(userId int64, topicId int64) error
	UnfollowTopic(userId int64, topicId int64) error
	StickTopic(userId int64, topicId int64) (int8, error)
	PinTopic(userId int64, topicId int64) (int8, error)
}

// TopicServantA 话题服务(版本A)
type TopicServantA interface {
	UpsertTags(userId int64, tags []string) (cs.TagInfoList, error)
	DecrTagsById(ids []int64) error
	ListTags(typ cs.TagType, limit int, offset int) (cs.TagList, error)
	TagsByKeyword(keyword string) (cs.TagInfoList, error)
}
