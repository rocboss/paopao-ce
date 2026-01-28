// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package cs contain core data service interface type
// model define

package cs

// RelationTyp 用户关系类型
type RelationTyp uint8

const (
	RelationUnknown RelationTyp = iota
	RelationSelf
	RelationFriend
	RelationFollower
	RelationFollowing
	RelationAdmin
	RelationGuest
)

// String 返回关系类型的字符串表示
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
