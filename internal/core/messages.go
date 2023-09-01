// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

// MessageService 消息服务
type MessageService interface {
	CreateMessage(msg *ms.Message) (*ms.Message, error)
	GetUnreadCount(userID int64) (int64, error)
	GetMessageByID(id int64) (*ms.Message, error)
	ReadMessage(message *ms.Message) error
	GetMessages(userId int64, offset, limit int) ([]*ms.MessageFormated, error)
	GetMessageCount(userId int64) (int64, error)
}
