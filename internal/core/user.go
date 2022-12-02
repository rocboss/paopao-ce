// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

type (
	ContactItem struct {
		UserId   int64  `json:"user_id"`
		UserName string `json:"username"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		Phone    string `json:"phone"`
	}

	ContactList struct {
		Contacts []ContactItem `json:"contacts"`
		Total    int64         `json:"total"`
	}
)

// UserManageService 用户管理服务
type UserManageService interface {
	GetUserByID(id int64) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByPhone(phone string) (*User, error)
	GetUsersByIDs(ids []int64) ([]*User, error)
	GetUsersByKeyword(keyword string) ([]*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) error
}

// ContactManageService 联系人管理服务
type ContactManageService interface {
	RequestingFriend(userId int64, friendId int64, greetings string) error
	AddFriend(userId int64, friendId int64) error
	RejectFriend(userId int64, friendId int64) error
	DeleteFriend(userId int64, friendId int64) error
	GetContacts(userId int64, offset int, limit int) (*ContactList, error)
	IsFriend(userID int64, friendID int64) bool
}
