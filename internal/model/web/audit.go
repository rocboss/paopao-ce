// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

const (
	AuditStyleUnknown AuditStyle = iota
	AuditStyleUserTweet
	AuditStyleUserTweetComment
	AuditStyleUserTweetReply
)

const (
	AuditHookCtxKey  = "audit_ctx_key"
	OnlineUserCtxKey = "online_user_ctx_key"
)

type AuditStyle uint8

type AuditMetaInfo struct {
	Style AuditStyle
	Id    int64
}

func (s AuditStyle) String() (res string) {
	switch s {
	case AuditStyleUserTweet:
		res = "UserTweet"
	case AuditStyleUserTweetComment:
		res = "UserTweetComment"
	case AuditStyleUserTweetReply:
		res = "UserTweetReply"
	case AuditStyleUnknown:
		fallthrough
	default:
		res = "Unknown"
	}
	return
}
