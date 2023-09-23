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

type CommentMetric struct {
	*Model
	CommentId        int64
	RankScore        int64
	IncentiveScore   int
	DecayFactor      int
	MotivationFactor int
}

type UserMetric struct {
	*Model
	UserId         int64
	TweetsCount    int
	LatestTrendsOn int64
}

func (m *PostMetric) Create(db *gorm.DB) (*PostMetric, error) {
	err := db.Create(&m).Error
	return m, err
}

func (m *PostMetric) Delete(db *gorm.DB) error {
	return db.Model(m).Where("post_id", m.PostId).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}

func (m *CommentMetric) Create(db *gorm.DB) (*CommentMetric, error) {
	err := db.Create(&m).Error
	return m, err
}

func (m *CommentMetric) Delete(db *gorm.DB) error {
	return db.Model(m).Where("comment_id", m.CommentId).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}

func (m *UserMetric) Create(db *gorm.DB) (*UserMetric, error) {
	err := db.Create(&m).Error
	return m, err
}

func (m *UserMetric) Delete(db *gorm.DB) error {
	return db.Model(m).Where("user_id", m.UserId).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}
