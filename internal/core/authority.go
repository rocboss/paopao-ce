// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/pkg/types"
)

const (
	ActRegisterUser act = iota
	ActCreatePublicTweet
	ActCreatePublicAttachment
	ActCreatePublicPicture
	ActCreatePublicVideo
	ActCreatePrivateTweet
	ActCreatePrivateAttachment
	ActCreatePrivatePicture
	ActCreatePrivateVideo
	ActCreateFriendTweet
	ActCreateFriendAttachment
	ActCreateFriendPicture
	ActCreateFriendVideo
	ActCreatePublicComment
	ActCreatePublicPicureComment
	ActCreateFriendComment
	ActCreateFriendPicureComment
	ActCreatePrivateComment
	ActCreatePrivatePicureComment
	ActStickTweet
	ActTopTweet
	ActLockTweet
	ActVisibleTweet
	ActDeleteTweet
	ActCreateActivationCode
)

type (
	act uint8

	FriendFilter map[int64]types.Empty
	FriendSet    map[string]types.Empty

	Action struct {
		Act    act
		UserId int64
	}
)

func (f FriendFilter) IsFriend(userId int64) bool {
	_, yeah := f[userId]
	return yeah
}

// IsAllow default true if user is admin
func (a act) IsAllow(user *User, userId int64, isFriend bool, isActivation bool) bool {
	if user.IsAdmin {
		return true
	}
	if user.ID == userId && isActivation {
		switch a {
		case ActCreatePublicTweet,
			ActCreatePublicAttachment,
			ActCreatePublicPicture,
			ActCreatePublicVideo,
			ActCreatePrivateTweet,
			ActCreatePrivateAttachment,
			ActCreatePrivatePicture,
			ActCreatePrivateVideo,
			ActCreateFriendTweet,
			ActCreateFriendAttachment,
			ActCreateFriendPicture,
			ActCreateFriendVideo,
			ActCreatePrivateComment,
			ActCreatePrivatePicureComment,
			ActStickTweet,
			ActLockTweet,
			ActVisibleTweet,
			ActDeleteTweet:
			return true
		}
	}

	if user.ID == userId && !isActivation {
		switch a {
		case ActCreatePrivateTweet,
			ActCreatePrivateComment,
			ActStickTweet,
			ActLockTweet,
			ActDeleteTweet:
			return true
		}
	}

	if isFriend && isActivation {
		switch a {
		case ActCreatePublicComment,
			ActCreatePublicPicureComment,
			ActCreateFriendComment,
			ActCreateFriendPicureComment:
			return true
		}
	}

	if !isFriend && isActivation {
		switch a {
		case ActCreatePublicComment,
			ActCreatePublicPicureComment:
			return true
		}
	}

	return false
}

// AuthorizationManageService 授权管理服务
type AuthorizationManageService interface {
	IsAllow(user *User, action *Action) bool
	BeFriendFilter(userId int64) FriendFilter
	BeFriendIds(userId int64) ([]int64, error)
	MyFriendSet(userId int64) FriendSet
}
