// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core/cs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostStar struct {
	*Model
	Post   *Post `json:"-"`
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
}

func (p *PostStar) Get(db *gorm.DB) (*PostStar, error) {
	var star PostStar
	tn := db.NamingStrategy.TableName("PostStar") + "."

	if p.Model != nil && p.ID > 0 {
		db = db.Where(tn+"id = ? AND "+tn+"is_del = ?", p.ID, 0)
	}
	if p.PostID > 0 {
		db = db.Where(tn+"post_id = ?", p.PostID)
	}
	if p.UserID > 0 {
		db = db.Where(tn+"user_id = ?", p.UserID)
	}

	db = db.Joins("Post").Where("visibility <> ? OR (visibility = ? AND ? = ?)", PostVisitPrivate, PostVisitPrivate, clause.Column{Table: "Post", Name: "user_id"}, p.UserID).Order(clause.OrderByColumn{Column: clause.Column{Table: "Post", Name: "id"}, Desc: true})
	if err := db.First(&star).Error; err != nil {
		return nil, err
	}
	return &star, nil
}

func (p *PostStar) Create(db *gorm.DB) (*PostStar, error) {
	err := db.Omit("Post").Create(&p).Error

	return p, err
}

func (p *PostStar) Delete(db *gorm.DB) error {
	return db.Model(&PostStar{}).Omit("Post").Where("id = ? AND is_del = ?", p.Model.ID, 0).Updates(map[string]any{
		"deleted_on": time.Now().Unix(),
		"is_del":     1,
	}).Error
}

func (p *PostStar) List(db *gorm.DB, conditions *ConditionsT, typ cs.RelationTyp, limit int, offset int) (res []*PostStar, err error) {
	tn := db.NamingStrategy.TableName("PostStar") + "."
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if p.UserID > 0 {
		db = db.Where(tn+"user_id = ?", p.UserID)
	}
	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(tn+k, v)
		}
	}
	db = db.Joins("Post")
	switch typ {
	case cs.RelationAdmin:
		// admin have all permition to visit all type tweets
	case cs.RelationFriend:
		db = db.Where("visibility = ? OR visibility = ?", PostVisitPublic, PostVisitFriend)
	case cs.RelationSelf:
		db = db.Where("visibility <> ? OR (visibility = ? AND ? = ?)", PostVisitPrivate, PostVisitPrivate, clause.Column{Table: "Post", Name: "user_id"}, p.UserID)
	default:
		db = db.Where("visibility=?", PostVisitPublic)
	}
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Table: "Post", Name: "id"}, Desc: true})
	err = db.Find(&res).Error
	return
}

func (p *PostStar) Count(db *gorm.DB, typ cs.RelationTyp, conditions *ConditionsT) (res int64, err error) {
	tn := db.NamingStrategy.TableName("PostStar") + "."
	if p.PostID > 0 {
		db = db.Where(tn+"post_id = ?", p.PostID)
	}
	if p.UserID > 0 {
		db = db.Where(tn+"user_id = ?", p.UserID)
	}
	for k, v := range *conditions {
		if k != "ORDER" {
			db = db.Where(tn+k, v)
		}
	}
	db = db.Joins("Post")
	switch typ {
	case cs.RelationAdmin:
		// admin have all permition to visit all type tweets
	case cs.RelationFriend:
		db = db.Where("visibility = ? OR visibility = ?", PostVisitPublic, PostVisitFriend)
	case cs.RelationSelf:
		db = db.Where("visibility <> ? OR (visibility = ? AND ? = ?)", PostVisitPrivate, PostVisitPrivate, clause.Column{Table: "Post", Name: "user_id"}, p.UserID)
	default:
		db = db.Where("visibility=?", PostVisitPublic)
	}
	err = db.Model(p).Count(&res).Error
	return
}
