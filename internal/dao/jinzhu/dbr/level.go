// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"gorm.io/gorm"
)

// UserLevel 用户等级配置模型
type UserLevel struct {
	Level         int    `gorm:"column:level;primaryKey" json:"level"`        // 等级
	MinExperience int    `gorm:"column:min_experience" json:"min_experience"` // 最小经验值
	MaxExperience int    `gorm:"column:max_experience" json:"max_experience"` // 最大经验值
	Name          string `gorm:"column:name" json:"name"`                     // 等级名称
	Description   string `gorm:"column:description" json:"description"`       // 等级描述
}

// TableName 指定表名映射
func (u *UserLevel) TableName() string {
	return "p_user_level"
}

// GetByExperience 根据经验值获取等级
func (u *UserLevel) GetByExperience(db *gorm.DB, experience int) (*UserLevel, error) {
	var level UserLevel
	err := db.Where("min_experience <= ? AND max_experience >= ?", experience, experience).First(&level).Error
	return &level, err
}

// GetAll 获取所有等级配置
func (u *UserLevel) GetAll(db *gorm.DB) ([]*UserLevel, error) {
	var levels []*UserLevel
	err := db.Order("level ASC").Find(&levels).Error
	return levels, err
}

// Create 创建等级配置
func (u *UserLevel) Create(db *gorm.DB) (*UserLevel, error) {
	err := db.Create(&u).Error
	return u, err
}

// Update 更新等级配置
func (u *UserLevel) Update(db *gorm.DB) error {
	return db.Model(&UserLevel{}).Where("level = ?", u.Level).Save(u).Error
}

// Delete 删除等级配置
func (u *UserLevel) Delete(db *gorm.DB) error {
	return db.Where("level = ?", u.Level).Delete(&UserLevel{}).Error
}
