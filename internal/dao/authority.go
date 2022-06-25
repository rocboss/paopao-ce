package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
)

func newSimpleAuthorizationManageService() *simpleAuthorizationManageService {
	return &simpleAuthorizationManageService{
		db: conf.DBEngine,
	}
}

func (s *simpleAuthorizationManageService) IsAllow(user *model.User, action *core.Action) bool {
	// user is activation if had bind phone
	isActivation := (len(user.Phone) != 0)
	isFriend := s.isFriend(action.UserId)
	// TODO: just use defaut act authorization chek rule now
	return action.Act.IsAllow(user, action.UserId, isFriend, isActivation)
}

// GetFriendFilter _userId保留未来使用
func (s *simpleAuthorizationManageService) GetFriendFilter(_userId int64) core.FriendFilter {
	// TODO: just return an empty friend fileter now
	return core.FriendFilter{}
}

func (s *simpleAuthorizationManageService) GetFriendIds(_userId int64) []int64 {
	// TODO: just retrun empty now
	return nil
}

func (s *simpleAuthorizationManageService) isFriend(_userId int64) bool {
	// friend with all world now
	return true
}
