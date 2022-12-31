// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "gorm.io/gorm"

const (
	UserStatusNormal int = iota + 1
	UserStatusClosed
)

type User struct {
	*Model
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	Balance  int64  `json:"balance"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserFormated struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"is_admin"`
}

func (u *User) Format() *UserFormated {
	if u.Model != nil {
		return &UserFormated{
			ID:       u.ID,
			Nickname: u.Nickname,
			Username: u.Username,
			Status:   u.Status,
			Avatar:   u.Avatar,
			IsAdmin:  u.IsAdmin,
		}
	}

	return nil
}

func (u *User) Get(db *gorm.DB) (*User, error) {
	var user User
	if u.Model != nil && u.Model.ID > 0 {
		db = db.Where("id= ? AND is_del = ?", u.Model.ID, 0)
	} else if u.Phone != "" {
		db = db.Where("phone = ? AND is_del = ?", u.Phone, 0)
	} else {
		db = db.Where("username = ? AND is_del = ?", u.Username, 0)
	}

	err := db.First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (u *User) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*User, error) {
	var users []*User
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) Create(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error

	return u, err
}

func (u *User) Update(db *gorm.DB) error {
	return db.Model(&User{}).Where("id = ? AND is_del = ?", u.Model.ID, 0).Save(u).Error
}
