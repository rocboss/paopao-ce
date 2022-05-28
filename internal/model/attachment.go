package model

import "gorm.io/gorm"

type AttachmentType int

const (
	ATTACHMENT_TYPE_IMAGE AttachmentType = iota + 1
	ATTACHMENT_TYPE_VIDEO
	ATTACHMENT_TYPE_OTHER
)

type Attachment struct {
	*Model
	UserID    int64          `json:"user_id"`
	FileSize  int64          `json:"file_size"`
	ImgWidth  int            `json:"img_width"`
	ImgHeight int            `json:"img_height"`
	Type      AttachmentType `json:"type"`
	Content   string         `json:"content"`
}

func (a *Attachment) Create(db *gorm.DB) (*Attachment, error) {
	err := db.Create(&a).Error

	return a, err
}
