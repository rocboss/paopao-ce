// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	ContactStatusRequesting int8 = iota + 1
	ContactStatusAgree
	ContactStatusReject
	ContactStatusDeleted
)

type Contact struct {
	*Model
	User         *User  `json:"-" gorm:"foreignKey:ID;references:FriendId"`
	UserId       int64  `json:"user_id"`
	FriendId     int64  `json:"friend_id"`
	GroupId      int64  `json:"group_id"`
	Remark       string `json:"remark"`
	Status       int8   `json:"status"` // 1请求好友, 2已同意好友, 3已拒绝好友, 4已删除好友
	IsTop        int8   `json:"is_top"`
	IsBlack      int8   `json:"is_black"`
	NoticeEnable int8   `json:"notice_enable"`
}

func (c *Contact) FetchUser(db *gorm.DB) (*Contact, error) {
	var contact Contact
	err := db.Omit("User").Unscoped().Where("user_id = ? AND friend_id = ?", c.UserId, c.FriendId).First(&contact).Error
	if err != nil {
		logrus.Debugf("Contact.FetchUser fetch user error:%s", err)
		return nil, err
	}
	return &contact, nil
}

func (c *Contact) GetByUserFriend(db *gorm.DB) (*Contact, error) {
	var contact Contact
	err := db.Omit("User").Where("user_id = ? AND friend_id = ?", c.UserId, c.FriendId).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (c *Contact) FetchByUserFriendAll(db *gorm.DB) ([]*Contact, error) {
	var contacts []*Contact
	if err := db.Omit("User").
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
			c.UserId, c.FriendId, c.FriendId, c.UserId).
		Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (c *Contact) List(db *gorm.DB, conditions ConditionsT, offset, limit int) ([]*Contact, error) {
	var contacts []*Contact
	var err error
	tn := db.NamingStrategy.TableName("Contact") + "."

	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	for k, v := range conditions {
		if k != "ORDER" {
			db = db.Where(tn+k, v)
		}
	}

	db.Joins("User").Order(clause.OrderByColumn{Column: clause.Column{Name: "nickname"}, Desc: false})
	if err = db.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (c *Contact) BeFriendIds(db *gorm.DB) (ids []int64, err error) {
	if err = db.Model(c).Omit("User").Select("user_id").Where("friend_id = ? AND status = ?", c.FriendId, ContactStatusAgree).Find(&ids).Error; err != nil {
		return nil, err
	}
	return
}

func (c *Contact) MyFriendIds(db *gorm.DB) (ids []string, err error) {
	if err = db.Model(c).Omit("User").Select("friend_id").Where("user_id = ? AND status = ?", c.UserId, ContactStatusAgree).Find(&ids).Error; err != nil {
		return nil, err
	}
	return
}

func (m *Contact) Count(db *gorm.DB, conditions ConditionsT) (int64, error) {
	var count int64

	for k, v := range conditions {
		if k != "ORDER" {
			db = db.Where(k, v)
		}
	}
	if err := db.Model(m).Omit("User").Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (c *Contact) Create(db *gorm.DB) (*Contact, error) {
	err := db.Omit("User").Create(&c).Error
	return c, err
}

func (c *Contact) Update(db *gorm.DB) error {
	return db.Model(&Contact{}).Omit("User").Where("id = ?", c.Model.ID).Save(c).Error
}

func (c *Contact) UpdateInUnscoped(db *gorm.DB) error {
	return db.Unscoped().Omit("User").Save(c).Error
}
