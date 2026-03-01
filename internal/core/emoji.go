// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

// EmojiService 表情包服务接口
type EmojiService interface {
	// 上传表情包
	UploadEmoji(userID int64, name, url string, width, height int, size int64, emojiType string) (*dbr.Emoji, error)

	// 获取表情包列表
	GetEmojiList(offset, limit int) ([]*dbr.Emoji, error)

	// 获取用户上传的表情包列表
	GetUserEmojiList(userID int64, offset, limit int) ([]*dbr.Emoji, error)

	// 获取用户收藏的表情包列表
	GetUserCollectedEmojiList(userID int64, offset, limit int) ([]*dbr.Emoji, error)

	// 收藏表情包
	CollectEmoji(userID, emojiID int64) error

	// 取消收藏表情包
	UncollectEmoji(userID, emojiID int64) error

	// 检查表情包是否被用户收藏
	IsEmojiCollected(userID, emojiID int64) (bool, error)

	// 获取表情包详情
	GetEmojiByID(emojiID int64) (*dbr.Emoji, error)

	// 删除表情包
	DeleteEmoji(userID, emojiID int64) error
}
