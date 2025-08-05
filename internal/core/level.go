// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

// UserLevelService 用户等级服务接口
type UserLevelService interface {
	// GetLevelByExperience 根据经验值获取用户等级
	GetLevelByExperience(experience int) (*dbr.UserLevel, error)

	// GetAllLevels 获取所有等级配置
	GetAllLevels() ([]*dbr.UserLevel, error)

	// CreateLevel 创建等级配置
	CreateLevel(level *dbr.UserLevel) (*dbr.UserLevel, error)

	// UpdateLevel 更新等级配置
	UpdateLevel(level *dbr.UserLevel) error

	// DeleteLevel 删除等级配置
	DeleteLevel(level int) error
}