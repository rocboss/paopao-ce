package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/model/rest"
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
}

// ContactManageService 联系人管理服务
type ContactManageService interface {
	RequestingFriend(userId int64, friendId int64, greetings string) error
	AddFriend(userId int64, friendId int64) error
	RejectFriend(userId int64, friendId int64) error
	DeleteFriend(userId int64, friendId int64) error
	GetContacts(userId int64, offset int, limit int) (*rest.ContactsResp, error)
	IsFriend(userID int64, friendID int64) bool
}
