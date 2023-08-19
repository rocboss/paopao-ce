// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
	"github.com/sirupsen/logrus"
)

var (
	_ core.ContactManageService = (*contactManageSrv)(nil)
)

type contactManageSrv struct {
	*sqlxSrv
	q *cc.ContactManager
}

func (s *contactManageSrv) RequestingFriend(userId int64, friendId int64, greetings string) error {
	return s.db.Withx(func(tx *sqlx.Tx) error {
		contact, err := s.fetchOrNewContact(tx, userId, friendId, cs.ContactStatusRequesting)
		if err != nil {
			return err
		}
		now := time.Now().Unix()
		// 如果已经好友，啥也不干
		if contact.Status == cs.ContactStatusAgree {
			return nil
		} else if contact.Status == cs.ContactStatusReject || contact.Status == cs.ContactStatusDeleted {
			contact.Status = cs.ContactStatusRequesting
			contact.IsDel = 0 // remove deleted flag if needed
			args := []any{cs.ContactStatusRequesting, now, contact.ID}
			if _, err = tx.Stmtx(s.q.FreshContactStatus).Exec(args...); err != nil {
				return err
			}
		}
		msg := &ms.Message{
			SenderUserID:   userId,
			ReceiverUserID: friendId,
			Type:           ms.MsgTypeRequestingFriend,
			Brief:          "请求添加好友，并附言:",
			Content:        greetings,
			ReplyID:        int64(cs.ContactStatusRequesting),
			Model: &ms.Model{
				CreatedOn: now,
			},
		}
		if _, err = tx.NamedStmt(s.q.CreateMessage).Exec(msg); err != nil {
			logrus.Errorf("contactManageSrv.RequestingFriend create message err:%s", err)
			return err
		}
		return nil
	})
}

func (s *contactManageSrv) AddFriend(userId int64, friendId int64) error {
	return s.db.Withx(func(tx *sqlx.Tx) error {
		contact := &cs.Contact{}
		err := tx.Stmtx(s.q.GetUserFriend).Get(contact, userId, friendId)
		if err != nil {
			return err
		}
		// 如果还不是请求好友，啥也不干
		if contact.Status != cs.ContactStatusRequesting {
			logrus.Debugf("contactManageSrv.AddFriend not reuesting status now so skip")
			return nil
		}
		now := time.Now().Unix()
		args := []any{cs.ContactStatusAgree, now, contact.ID}
		if _, err = tx.Stmtx(s.q.FreshContactStatus).Exec(args...); err != nil {
			return err
		}
		contact, err = s.fetchOrNewContact(tx, userId, friendId, cs.ContactStatusAgree)
		if err != nil {
			return err
		}
		// 如果已经好友，啥也不干
		if contact.Status != cs.ContactStatusAgree {
			args = []any{cs.ContactStatusAgree, now, contact.ID}
			if _, err = tx.Stmtx(s.q.FreshContactStatus).Exec(args); err != nil {
				return err
			}
		}
		args = []any{cs.ContactStatusAgree, now, userId, friendId, friendId, userId, ms.MsgTypeRequestingFriend, cs.ContactStatusRequesting}
		if _, err = tx.Stmtx(s.q.AddFriendMsgsUpdate).Exec(args...); err != nil {
			return err
		}
		return nil
	})
}

func (s *contactManageSrv) RejectFriend(userId int64, friendId int64) error {
	return s.db.Withx(func(tx *sqlx.Tx) error {
		contact := &cs.Contact{}
		err := tx.Stmtx(s.q.GetUserFriend).Get(contact, userId, friendId)
		if err != nil {
			return err
		}
		// 如果还不是请求好友，啥也不干
		if contact.Status != cs.ContactStatusRequesting {
			return nil
		}
		now := time.Now().Unix()
		args := []any{cs.ContactStatusReject, now, contact.ID}
		if _, err = tx.Stmtx(s.q.FreshContactStatus).Exec(args...); err != nil {
			return err
		}
		args = []any{cs.ContactStatusReject, now, friendId, userId, ms.MsgTypeRequestingFriend, cs.ContactStatusRequesting}
		if _, err = tx.Stmtx(s.q.RejectFriendMsgsUpdate).Exec(args...); err != nil {
			return err
		}
		return nil
	})
}

func (s *contactManageSrv) DeleteFriend(userId int64, friendId int64) error {
	return s.db.Withx(func(tx *sqlx.Tx) error {
		var contacts []cs.Contact
		err := tx.Stmtx(s.q.GetContacts).Select(&contacts, userId, friendId, friendId, userId)
		if err != nil {
			return err
		}
		for _, contact := range contacts {
			// 如果还不是好友，啥也不干
			if contact.Status != cs.ContactStatusAgree {
				continue
			}
			if _, err = tx.Stmtx(s.q.DelFriend).Exec(time.Now().Unix(), contact.ID); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *contactManageSrv) GetContacts(userId int64, offset int, limit int) (res *ms.ContactList, err error) {
	res = &ms.ContactList{
		Contacts: []ms.ContactItem{},
	}
	if err = s.q.ListFriend.Select(&res.Contacts, userId, limit, offset); err != nil {
		return
	}
	if err = s.q.TotalFriendsById.Get(&res.Total, userId); err != nil {
		return
	}
	return
}

func (s *contactManageSrv) IsFriend(userId int64, friendId int64) (res bool) {
	s.q.IsFriend.Get(&res, userId, friendId)
	return
}

func (s *contactManageSrv) fetchOrNewContact(tx *sqlx.Tx, userId int64, friendId int64, status int8) (res *cs.Contact, err error) {
	res = &cs.Contact{}
	if err = tx.Stmtx(s.q.GetContact).Get(res, userId, friendId); err == nil {
		return
	}
	result, xerr := tx.Stmtx(s.q.CreateContact).Exec(userId, friendId, status, time.Now().Unix())
	if xerr != nil {
		return nil, xerr
	}
	id, xerr := result.LastInsertId()
	if xerr != nil {
		return nil, xerr
	}
	return &cs.Contact{
		ID:       id,
		UserId:   userId,
		FriendId: friendId,
		Status:   status,
	}, nil
}

func newContactManageService(db *sqlx.DB) core.ContactManageService {
	return &contactManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildContactManager),
	}
}
