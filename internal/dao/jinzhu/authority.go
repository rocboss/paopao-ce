package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"gorm.io/gorm"
)

var (
	_ core.AuthorizationManageService = (*authorizationManageServant)(nil)
)

type authorizationManageServant struct {
	db *gorm.DB
}

func (s *authorizationManageServant) IsAllow(user *model.User, action *core.Action) bool {
	// user is activation if had bind phone
	isActivation := (len(user.Phone) != 0)
	isFriend := s.isFriend(user.ID, action.UserId)
	// TODO: just use defaut act authorization chek rule now
	return action.Act.IsAllow(user, action.UserId, isFriend, isActivation)
}

func (s *authorizationManageServant) BeFriendFilter(userId int64) core.FriendFilter {
	// just empty now
	return core.FriendFilter{}
}

func (s *authorizationManageServant) BeFriendIds(userId int64) ([]int64, error) {
	// just empty now
	return []int64{}, nil
}

func (s *authorizationManageServant) isFriend(userId int64, friendId int64) bool {
	// just true now
	return true
}
