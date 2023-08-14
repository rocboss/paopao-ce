// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Following struct {
	*Model
	User     *User `json:"-" gorm:"foreignKey:ID;references:FollowId"`
	UserId   int64 `json:"user_id"`
	FollowId int64 `json:"friend_id"`
}

func (f *Following) GetFollowing(db *gorm.DB, userId, followId int64) (*Following, error) {
	var following Following
	err := db.Omit("User").Unscoped().Where("user_id = ? AND follow_id = ?", userId, followId).First(&following).Error
	if err != nil {
		logrus.Debugf("Following.GetFollowing get following error:%s", err)
		return nil, err
	}
	return &following, nil
}

func (f *Following) DelFollowing(db *gorm.DB, userId, followId int64) error {
	return db.Omit("User").Unscoped().Where("user_id = ? AND follow_id = ?", userId, followId).Delete(f).Error
}

func (f *Following) ListFollows(db *gorm.DB, userId int64, limit int, offset int) (res []*Following, total int64, err error) {
	db = db.Model(f).Where("user_id=?", userId)
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	db.Joins("User").Order(clause.OrderByColumn{Column: clause.Column{Table: "User", Name: "nickname"}, Desc: false})
	if err = db.Find(&res).Error; err != nil {
		return
	}
	return
}

func (f *Following) ListFollowingIds(db *gorm.DB, userId int64, limit, offset int) (ids []int64, total int64, err error) {
	db = db.Model(f).Where("follow_id=?", userId)
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if err = db.Omit("User").Select("user_id").Find(&ids).Error; err != nil {
		return
	}
	return
}

func (f *Following) FollowCount(db *gorm.DB, userId int64) (follows int64, followings int64, err error) {
	if err = db.Model(f).Where("user_id=?", userId).Count(&follows).Error; err != nil {
		return
	}
	if err = db.Model(f).Where("follow_id=?", userId).Count(&followings).Error; err != nil {
		return
	}
	return
}

func (s *Following) IsFollow(db *gorm.DB, userId int64, followId int64) bool {
	if _, err := s.GetFollowing(db, userId, followId); err == nil {
		return true
	}
	return false
}

func (f *Following) Create(db *gorm.DB) (*Following, error) {
	err := db.Omit("User").Create(f).Error
	return f, err
}

func (c *Following) UpdateInUnscoped(db *gorm.DB) error {
	return db.Unscoped().Omit("User").Save(c).Error
}
