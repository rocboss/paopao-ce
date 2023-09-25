// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	// 消息列表样式
	StyleMsgAll        MessageStyle = "all"
	StyleMsgSystem     MessageStyle = "system"
	StyleMsgWhisper    MessageStyle = "whisper"
	StyleMsgRequesting MessageStyle = "requesting"
	StyleMsgUnread     MessageStyle = "unread"
)

type MessageStyle string
