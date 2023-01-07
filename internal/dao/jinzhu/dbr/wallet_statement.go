// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "gorm.io/gorm"

type WalletStatement struct {
	*Model
	UserID          int64  `json:"user_id"`
	ChangeAmount    int64  `json:"change_amount"`
	BalanceSnapshot int64  `json:"balance_snapshot"`
	Reason          string `json:"reason"`
	PostID          int64  `json:"post_id"`
}

func (w *WalletStatement) Get(db *gorm.DB) (*WalletStatement, error) {
	var ws WalletStatement
	if w.Model != nil && w.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", w.ID, 0)
	}
	if w.PostID > 0 {
		db = db.Where("post_id = ?", w.PostID)
	}
	if w.UserID > 0 {
		db = db.Where("user_id = ?", w.UserID)
	}

	err := db.First(&ws).Error
	if err != nil {
		return &ws, err
	}

	return &ws, nil
}

func (w *WalletStatement) Create(db *gorm.DB) (*WalletStatement, error) {
	err := db.Create(&w).Error

	return w, err
}

func (w *WalletStatement) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*WalletStatement, error) {
	var records []*WalletStatement
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if w.UserID > 0 {
		db = db.Where("user_id = ?", w.UserID)
	}

	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (w *WalletStatement) Count(db *gorm.DB, conditions *ConditionsT) (int64, error) {
	var count int64
	if w.PostID > 0 {
		db = db.Where("post_id = ?", w.PostID)
	}
	if w.UserID > 0 {
		db = db.Where("user_id = ?", w.UserID)
	}
	for k, v := range *conditions {
		if k != "ORDER" {
			db = db.Where(k, v)
		}
	}
	if err := db.Model(w).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
