// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	RelationUnknown RelationTyp = iota
	RelationSelf
	RelationFriend
	RelationFollower
	RelationFollowing
	RelationAdmin
	RelationGuest
)

type (
	// UserInfoList 用户信息列表
	UserInfoList []*UserInfo

	//
	RelationTyp uint8

	VistUser struct {
		Username string
		UserId   int64
		RelTyp   RelationTyp
	}
)

// UserInfo 用户基本信息
type UserInfo struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
	Status    int    `json:"status"`
	Avatar    string `json:"avatar"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedOn int64  `json:"created_on"`
}

type UserProfile struct {
	ID          int64  `json:"id" db:"id"`
	Nickname    string `json:"nickname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Status      int    `json:"status"`
	Avatar      string `json:"avatar"`
	Balance     int64  `json:"balance"`
	IsAdmin     bool   `json:"is_admin"`
	CreatedOn   int64  `json:"created_on"`
	TweetsCount int    `json:"tweets_count"`
}

func (t RelationTyp) String() string {
	switch t {
	case RelationSelf:
		return "self"
	case RelationFriend:
		return "friend"
	case RelationFollower:
		return "follower"
	case RelationFollowing:
		return "following"
	case RelationAdmin:
		return "admin"
	case RelationUnknown:
		fallthrough
	default:
		return "unknown"
	}
}
