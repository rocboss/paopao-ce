// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

const (
	MsgTypePost             = dbr.MsgTypePost
	MsgtypeComment          = dbr.MsgtypeComment
	MsgTypeReply            = dbr.MsgTypeReply
	MsgTypeWhisper          = dbr.MsgTypeWhisper
	MsgTypeRequestingFriend = dbr.MsgTypeRequestingFriend
	MsgTypeSystem           = dbr.MsgTypeSystem

	MsgStatusUnread = dbr.MsgStatusUnread
	MsgStatusReaded = dbr.MsgStatusReaded
)

type (
	Message         = dbr.Message
	MessageFormated = dbr.MessageFormated
)

// MessageService 消息服务
type MessageService interface {
	CreateMessage(msg *Message) (*Message, error)
	GetUnreadCount(userID int64) (int64, error)
	GetMessageByID(id int64) (*Message, error)
	ReadMessage(message *Message) error
	GetMessages(conditions *ConditionsT, offset, limit int) ([]*MessageFormated, error)
	GetMessageCount(conditions *ConditionsT) (int64, error)
}
