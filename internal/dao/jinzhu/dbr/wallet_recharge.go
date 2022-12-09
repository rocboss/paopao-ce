// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "gorm.io/gorm"

type WalletRecharge struct {
	*Model
	UserID      int64  `json:"user_id"`
	Amount      int64  `json:"amount"`
	TradeNo     string `json:"trade_no"`
	TradeStatus string `json:"trade_status"`
}

func (p *WalletRecharge) Get(db *gorm.DB) (*WalletRecharge, error) {
	var pas WalletRecharge
	if p.Model != nil && p.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", p.ID, 0)
	}
	if p.UserID > 0 {
		db = db.Where("user_id = ?", p.UserID)
	}

	err := db.First(&pas).Error
	if err != nil {
		return &pas, err
	}

	return &pas, nil
}

func (p *WalletRecharge) Create(db *gorm.DB) (*WalletRecharge, error) {
	err := db.Create(&p).Error

	return p, err
}
