// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

// PostVisibleT 可访问类型，0公开，1私密，2好友
type PostVisibleT uint8

const (
	PostVisitPublic PostVisibleT = iota
	PostVisitPrivate
	PostVisitFriend
	PostVisitInvalid
)

type Post struct {
	*Model
	UserID          int64        `json:"user_id"`
	CommentCount    int64        `json:"comment_count"`
	CollectionCount int64        `json:"collection_count"`
	ShareCount      int64        `json:"share_count"`
	UpvoteCount     int64        `json:"upvote_count"`
	Visibility      PostVisibleT `json:"visibility"`
	IsTop           int          `json:"is_top"`
	IsEssence       int          `json:"is_essence"`
	IsLock          int          `json:"is_lock"`
	LatestRepliedOn int64        `json:"latest_replied_on"`
	Tags            string       `json:"tags"`
	AttachmentPrice int64        `json:"attachment_price"`
	IP              string       `json:"ip"`
	IPLoc           string       `json:"ip_loc"`
}

type PostFormated struct {
	ID              int64                  `json:"id"`
	UserID          int64                  `json:"user_id"`
	User            *UserFormated          `json:"user"`
	Contents        []*PostContentFormated `json:"contents"`
	CommentCount    int64                  `json:"comment_count"`
	CollectionCount int64                  `json:"collection_count"`
	ShareCount      int64                  `json:"share_count"`
	UpvoteCount     int64                  `json:"upvote_count"`
	Visibility      PostVisibleT           `json:"visibility"`
	IsTop           int                    `json:"is_top"`
	IsEssence       int                    `json:"is_essence"`
	IsLock          int                    `json:"is_lock"`
	LatestRepliedOn int64                  `json:"latest_replied_on"`
	CreatedOn       int64                  `json:"created_on"`
	ModifiedOn      int64                  `json:"modified_on"`
	Tags            map[string]int8        `json:"tags"`
	AttachmentPrice int64                  `json:"attachment_price"`
	IPLoc           string                 `json:"ip_loc"`
}

func (p *Post) Format() *PostFormated {
	if p.Model != nil {
		tagsMap := map[string]int8{}
		for _, tag := range strings.Split(p.Tags, ",") {
			tagsMap[tag] = 1
		}
		return &PostFormated{
			ID:              p.ID,
			UserID:          p.UserID,
			User:            &UserFormated{},
			Contents:        []*PostContentFormated{},
			CommentCount:    p.CommentCount,
			CollectionCount: p.CollectionCount,
			ShareCount:      p.ShareCount,
			UpvoteCount:     p.UpvoteCount,
			Visibility:      p.Visibility,
			IsTop:           p.IsTop,
			IsEssence:       p.IsEssence,
			IsLock:          p.IsLock,
			LatestRepliedOn: p.LatestRepliedOn,
			CreatedOn:       p.CreatedOn,
			ModifiedOn:      p.ModifiedOn,
			AttachmentPrice: p.AttachmentPrice,
			Tags:            tagsMap,
			IPLoc:           p.IPLoc,
		}
	}

	return nil
}

func (p *Post) Create(db *gorm.DB) (*Post, error) {
	err := db.Create(&p).Error

	return p, err
}

func (s *Post) Delete(db *gorm.DB) error {
	return db.Model(s).Where("id = ?", s.Model.ID).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}

func (p *Post) Get(db *gorm.DB) (*Post, error) {
	var post Post
	if p.Model != nil && p.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", p.ID, 0)
	} else {
		return nil, gorm.ErrRecordNotFound
	}

	err := db.First(&post).Error
	if err != nil {
		return &post, err
	}

	return &post, nil
}

func (p *Post) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*Post, error) {
	var posts []*Post
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if p.UserID > 0 {
		db = db.Where("user_id = ?", p.UserID)
	}
	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Post) Fetch(db *gorm.DB, predicates Predicates, offset, limit int) ([]*Post, error) {
	var posts []*Post
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if p.UserID > 0 {
		db = db.Where("user_id = ?", p.UserID)
	}
	for query, args := range predicates {
		if query == "ORDER" {
			db = db.Order(args[0])
		} else {
			db = db.Where(query, args...)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Post) CountBy(db *gorm.DB, predicates Predicates) (count int64, err error) {
	for query, args := range predicates {
		if query != "ORDER" {
			db = db.Where(query, args...)
		}
	}
	err = db.Model(p).Count(&count).Error
	return
}

func (p *Post) Count(db *gorm.DB, conditions *ConditionsT) (int64, error) {
	var count int64
	if p.UserID > 0 {
		db = db.Where("user_id = ?", p.UserID)
	}
	for k, v := range *conditions {
		if k != "ORDER" {
			db = db.Where(k, v)
		}
	}
	if err := db.Model(p).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (p *Post) Update(db *gorm.DB) error {
	return db.Model(&Post{}).Where("id = ? AND is_del = ?", p.Model.ID, 0).Save(p).Error
}

func (p PostVisibleT) String() string {
	switch p {
	case PostVisitPublic:
		return "public"
	case PostVisitPrivate:
		return "private"
	case PostVisitFriend:
		return "friend"
	case PostVisitInvalid:
		return "invalid"
	default:
		return "unknow"
	}
}
