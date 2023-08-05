// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ms

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
