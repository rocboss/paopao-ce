// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "gorm.io/gorm"

type AttachmentType int

const (
	AttachmentTypeImage AttachmentType = iota + 1
	AttachmentTypeVideo
	AttachmentTypeOther
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
