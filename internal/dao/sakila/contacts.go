// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
)

var (
	_ core.ContactManageService = (*contactManageSrv)(nil)
)

type contact struct {
	UserId       int64 `db:"user_id" json:"user_id"`
	FriendId     int64 `db:"friend_id" json:"friend_id"`
	Status       int8  `json:"status"` // 1请求好友, 2已同意好友, 3已拒绝好友, 4已删除好友
	IsBlack      int8  `json:"is_black"`
	NoticeEnable int8  `json:"notice_enable"`
}

type contactManageSrv struct {
	*sqlxSrv
	q *cc.ContactManager
}

func (s *contactManageSrv) RequestingFriend(userId int64, friendId int64, greetings string) (err error) {
	_, err = s.q.RejectFriend.Query(userId, friendId, greetings)
	return
}

func (s *contactManageSrv) AddFriend(userId int64, friendId int64) (err error) {
	_, err = s.q.AddFriend.Exec(userId, friendId)
	return
}

func (s *contactManageSrv) RejectFriend(userId int64, friendId int64) (err error) {
	_, err = s.q.RejectFriend.Exec(userId, friendId)
	return
}

func (s *contactManageSrv) DeleteFriend(userId int64, friendId int64) (err error) {
	_, err = s.q.DelFriend.Exec(userId, friendId)
	return
}

func (s *contactManageSrv) GetContacts(userId int64, offset int, limit int) (*ms.ContactList, error) {
	res := &ms.ContactList{
		Contacts: []ms.ContactItem{},
	}
	if err := s.q.GetContacts.Select(&res.Contacts, userId, limit, offset); err != nil {
		return nil, err
	}
	if err := s.q.TotalContactsById.Get(&res.Total, userId); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *contactManageSrv) IsFriend(userId int64, friendId int64) bool {
	ct := &contact{
		UserId:   friendId,
		FriendId: userId,
	}
	err := s.q.GetUserFriend.Get(ct, userId, friendId)
	if err == nil && ct.Status == cs.ContactStatusAgree {
		return true
	}
	return false
}

func newContactManageService(db *sqlx.DB) core.ContactManageService {
	return &contactManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildContactManager),
	}
}
