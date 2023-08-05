// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import "github.com/rocboss/paopao-ce/internal/core/ms"

// UserManageService 用户管理服务
type UserManageService interface {
	GetUserByID(id int64) (*ms.User, error)
	GetUserByUsername(username string) (*ms.User, error)
	GetUserByPhone(phone string) (*ms.User, error)
	GetUsersByIDs(ids []int64) ([]*ms.User, error)
	GetUsersByKeyword(keyword string) ([]*ms.User, error)
	CreateUser(user *ms.User) (*ms.User, error)
	UpdateUser(user *ms.User) error
}

// ContactManageService 联系人管理服务
type ContactManageService interface {
	RequestingFriend(userId int64, friendId int64, greetings string) error
	AddFriend(userId int64, friendId int64) error
	RejectFriend(userId int64, friendId int64) error
	DeleteFriend(userId int64, friendId int64) error
	GetContacts(userId int64, offset int, limit int) (*ms.ContactList, error)
	IsFriend(userID int64, friendID int64) bool
}
