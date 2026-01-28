// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserBlock struct {
	*Model
	UserId       int64 `json:"user_id"`
	TargetUserId int64 `json:"target_user_id"`
	Reason       string `json:"reason"`
}

func (UserBlock) TableName() string {
	return "p_user_block"
}

func (b *UserBlock) GetUserBlock(db *gorm.DB, userId, targetUserId int64) (*UserBlock, error) {
	var block UserBlock
	err := db.Unscoped().Where("user_id = ? AND target_user_id = ?", userId, targetUserId).First(&block).Error
	if err != nil {
		logrus.Debugf("UserBlock.GetUserBlock get block error:%s", err)
		return nil, err
	}
	return &block, nil
}

func (b *UserBlock) DelUserBlock(db *gorm.DB, userId, targetUserId int64) error {
	return db.Unscoped().Where("user_id = ? AND target_user_id = ?", userId, targetUserId).Delete(b).Error
}

func (b *UserBlock) ListUserBlocks(db *gorm.DB, userId int64, limit int, offset int) (res []*UserBlock, total int64, err error) {
	db = db.Model(b).Where("user_id=?", userId)
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

func (b *UserBlock) IsBlocked(db *gorm.DB, userId, targetUserId int64) bool {
	if _, err := b.GetUserBlock(db, userId, targetUserId); err == nil {
		return true
	}
	return false
}

func (b *UserBlock) Create(db *gorm.DB) (*UserBlock, error) {
	err := db.Create(b).Error
	return b, err
}

func (b *UserBlock) UpdateInUnscoped(db *gorm.DB) error {
	return db.Unscoped().Save(b).Error
}
