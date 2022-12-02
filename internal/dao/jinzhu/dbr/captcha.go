// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "gorm.io/gorm"

type Captcha struct {
	*Model
	Phone     string `json:"phone"`
	Captcha   string `json:"captcha"`
	UseTimes  int    `json:"use_times"`
	ExpiredOn int64  `json:"expired_on"`
}

func (c *Captcha) Create(db *gorm.DB) (*Captcha, error) {
	err := db.Create(&c).Error

	return c, err
}

func (c *Captcha) Update(db *gorm.DB) error {
	return db.Model(&Captcha{}).Where("id = ? AND is_del = ?", c.Model.ID, 0).Save(c).Error
}

func (c *Captcha) Get(db *gorm.DB) (*Captcha, error) {
	var captcha Captcha
	if c.Model != nil && c.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", c.ID, 0)
	}
	if c.Phone != "" {
		db = db.Where("phone = ?", c.Phone)
	}

	err := db.Last(&captcha).Error
	if err != nil {
		return &captcha, err
	}

	return &captcha, nil
}
