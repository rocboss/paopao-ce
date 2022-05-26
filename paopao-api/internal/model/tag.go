package model

import "gorm.io/gorm"

type Tag struct {
	*Model
	UserID   int64  `json:"user_id"`
	Tag      string `json:"tag"`
	QuoteNum int64  `json:"quote_num"`
}
type TagFormated struct {
	ID       int64         `json:"id"`
	UserID   int64         `json:"user_id"`
	User     *UserFormated `json:"user"`
	Tag      string        `json:"tag"`
	QuoteNum int64         `json:"quote_num"`
}

func (t *Tag) Format() *TagFormated {
	if t.Model == nil {
		return &TagFormated{}
	}

	return &TagFormated{
		ID:       t.ID,
		UserID:   t.UserID,
		User:     &UserFormated{},
		Tag:      t.Tag,
		QuoteNum: t.QuoteNum,
	}
}

func (t *Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag Tag
	if t.Model != nil && t.Model.ID > 0 {
		db = db.Where("id= ? AND is_del = ?", t.Model.ID, 0)
	} else {
		db = db.Where("tag = ? AND is_del = ?", t.Tag, 0)
	}

	err := db.First(&tag).Error
	if err != nil {
		return &tag, err
	}

	return &tag, nil
}

func (t *Tag) Create(db *gorm.DB) (*Tag, error) {
	err := db.Create(&t).Error

	return t, err
}

func (t *Tag) Update(db *gorm.DB) error {
	return db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.Model.ID, 0).Save(t).Error
}

func (t *Tag) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if t.UserID > 0 {
		db = db.Where("user_id = ?", t.UserID)
	}
	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
