package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
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

type act uint8

type FriendFilter map[int64]types.Empty

type Action struct {
	Act    act
	UserId int64
}

// AuthorizationManageService 授权管理服务
type AuthorizationManageService interface {
	IsAllow(user *model.User, action *Action) bool
	GetFriendFilter(userId int64) FriendFilter
	GetFriendIds(userId int64) []int64
}

func (f FriendFilter) IsFriend(userId int64) bool {
	// _, yesno := f[userId]
	// return yesno
	// so, you are friend with all world now
	return true
}

// IsAllow default true if user is admin
func (a act) IsAllow(user *model.User, userId int64, isFriend bool, isActivation bool) bool {
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
