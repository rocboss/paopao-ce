package model

import "gorm.io/gorm"

// 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源

type PostContentT int

const (
	CONTENT_TYPE_TITLE PostContentT = iota + 1
	CONTENT_TYPE_TEXT
	CONTENT_TYPE_IMAGE
	CONTENT_TYPE_VIDEO
	CONTENT_TYPE_AUDIO
	CONTENT_TYPE_LINK
	CONTENT_TYPE_ATTACHMENT
	CONTENT_TYPE_CHARGE_ATTACHMENT
)

type PostContent struct {
	*Model
	PostID  int64        `json:"post_id"`
	UserID  int64        `json:"user_id"`
	Content string       `json:"content"`
	Type    PostContentT `json:"type"`
	Sort    int64        `json:"sort"`
}
type PostContentFormated struct {
	ID      int64        `json:"id"`
	PostID  int64        `json:"post_id"`
	Content string       `json:"content"`
	Type    PostContentT `json:"type"`
	Sort    int64        `json:"sort"`
}

func (p *PostContent) Create(db *gorm.DB) (*PostContent, error) {
	err := db.Create(&p).Error

	return p, err
}

func (p *PostContent) Format() *PostContentFormated {
	if p.Model == nil {
		return nil
	}
	return &PostContentFormated{
		ID:      p.ID,
		PostID:  p.ID,
		Content: p.Content,
		Type:    p.Type,
		Sort:    p.Sort,
	}
}

func (p *PostContent) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*PostContent, error) {
	var contents []*PostContent
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if p.PostID > 0 {
		db = db.Where("id = ?", p.PostID)
	}

	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&contents).Error; err != nil {
		return nil, err
	}

	return contents, nil
}

func (p *PostContent) Get(db *gorm.DB) (*PostContent, error) {
	var content PostContent
	if p.Model != nil && p.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", p.ID, 0)
	} else {
		return nil, gorm.ErrRecordNotFound
	}

	err := db.First(&content).Error
	if err != nil {
		return &content, err
	}

	return &content, nil
}
