// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	ContactStatusRequesting int8 = iota + 1
	ContactStatusAgree
	ContactStatusReject
	ContactStatusDeleted
)

type Contact struct {
	ID           int64  `db:"id" json:"id"`
	UserId       int64  `db:"user_id" json:"user_id"`
	FriendId     int64  `db:"friend_id" json:"friend_id"`
	GroupId      int64  `json:"group_id"`
	Remark       string `json:"remark"`
	Status       int8   `json:"status"` // 1请求好友, 2已同意好友, 3已拒绝好友, 4已删除好友
	IsTop        int8   `json:"is_top"`
	IsBlack      int8   `json:"is_black"`
	NoticeEnable int8   `json:"notice_enable"`
	IsDel        int8   `json:"-"`
	DeletedOn    int64  `db:"-" json:"-"`
}
