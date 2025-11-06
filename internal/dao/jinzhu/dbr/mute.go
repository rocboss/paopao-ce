// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserMute struct {
	*Model
	UserId       int64 `json:"user_id"`
	TargetUserId int64 `json:"target_user_id"`
	Reason       string `json:"reason"`
}

func (UserMute) TableName() string {
	return "p_user_mute"
}

func (m *UserMute) GetUserMute(db *gorm.DB, userId, targetUserId int64) (*UserMute, error) {
	var mute UserMute
	err := db.Unscoped().Where("user_id = ? AND target_user_id = ?", userId, targetUserId).First(&mute).Error
	if err != nil {
		logrus.Debugf("UserMute.GetUserMute get mute error:%s", err)
		return nil, err
	}
	return &mute, nil
}

func (m *UserMute) DelUserMute(db *gorm.DB, userId, targetUserId int64) error {
	return db.Unscoped().Where("user_id = ? AND target_user_id = ?", userId, targetUserId).Delete(m).Error
}

func (m *UserMute) ListUserMutes(db *gorm.DB, userId int64, limit int, offset int) (res []*UserMute, total int64, err error) {
	db = db.Model(m).Where("user_id=?", userId)
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if err = db.Find(&res).Error; err != nil {
		return
	}
	return
}

func (m *UserMute) IsMuted(db *gorm.DB, userId, targetUserId int64) bool {
	if _, err := m.GetUserMute(db, userId, targetUserId); err == nil {
		return true
	}
	return false
}

func (m *UserMute) Create(db *gorm.DB) (*UserMute, error) {
	err := db.Create(m).Error
	return m, err
}

func (m *UserMute) UpdateInUnscoped(db *gorm.DB) error {
	return db.Unscoped().Save(m).Error
}
