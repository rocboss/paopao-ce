// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.UserLevelService = (*userLevelSrv)(nil)
)

// userLevelSrv 实现用户等级服务接口
type userLevelSrv struct {
	db *gorm.DB
}

// newUserLevelService 创建用户等级服务实例
func newUserLevelService(db *gorm.DB) core.UserLevelService {
	return &userLevelSrv{
		db: db,
	}
}

// GetLevelByExperience 根据经验值获取用户等级
func (s *userLevelSrv) GetLevelByExperience(experience int) (*dbr.UserLevel, error) {
	level := &dbr.UserLevel{}
	return level.GetByExperience(s.db, experience)
}

// GetAllLevels 获取所有等级配置
func (s *userLevelSrv) GetAllLevels() ([]*dbr.UserLevel, error) {
	level := &dbr.UserLevel{}
	return level.GetAll(s.db)
}

// CreateLevel 创建等级配置
func (s *userLevelSrv) CreateLevel(level *dbr.UserLevel) (*dbr.UserLevel, error) {
	return level.Create(s.db)
}

// UpdateLevel 更新等级配置
func (s *userLevelSrv) UpdateLevel(level *dbr.UserLevel) error {
	return level.Update(s.db)
}

// DeleteLevel 删除等级配置
func (s *userLevelSrv) DeleteLevel(level int) error {
	return (&dbr.UserLevel{Level: level}).Delete(s.db)
}
