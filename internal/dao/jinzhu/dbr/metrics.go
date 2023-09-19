// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"

	"gorm.io/gorm"
)

type PostMetric struct {
	*Model
	PostId           int64
	RankScore        int64
	IncentiveScore   int
	DecayFactor      int
	MotivationFactor int
}

func (p *PostMetric) Create(db *gorm.DB) (*PostMetric, error) {
	err := db.Create(&p).Error
	return p, err
}

func (p *PostMetric) Delete(db *gorm.DB) error {
	return db.Model(p).Where("post_id", p.PostId).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}
