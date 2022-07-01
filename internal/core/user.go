package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// UserManageService 用户管理服务
type UserManageService interface {
	GetUserByID(id int64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByPhone(phone string) (*model.User, error)
	GetUsersByIDs(ids []int64) ([]*model.User, error)
	GetUsersByKeyword(keyword string) ([]*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	IsFriend(userID int64, friendID int64) bool
}
