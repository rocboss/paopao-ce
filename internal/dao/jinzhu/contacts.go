// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ core.ContactManageService = (*contactManageServant)(nil)
)

type contactManageServant struct {
	db *gorm.DB
}

func newContactManageService(db *gorm.DB) core.ContactManageService {
	return &contactManageServant{
		db: db,
	}
}

func (s *contactManageServant) fetchOrNewContact(db *gorm.DB, userId int64, friendId int64, status int8) (*dbr.Contact, error) {
	contact := &dbr.Contact{
		UserId:   userId,
		FriendId: friendId,
	}
	contact, err := contact.FetchUser(db)
	if err != nil {
		contact = &dbr.Contact{
			UserId:   userId,
			FriendId: friendId,
			Status:   status,
		}
		if contact, err = contact.Create(db); err != nil {
			logrus.Errorf("contactManageServant.fetchOrNewContact create new contact err:%s", err)
			return nil, err
		}
	}
	return contact, nil
}

func (s *contactManageServant) RequestingFriend(userId int64, friendId int64, greetings string) (err error) {
	db := s.db.Begin()
	defer func() {
		if err == nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()

	contact, e := s.fetchOrNewContact(db, userId, friendId, dbr.ContactStatusRequesting)
	if e != nil {
		err = e
		return
	}

	// 如果已经好友，啥也不干
	if contact.Status == dbr.ContactStatusAgree {
		return nil
	} else if contact.Status == dbr.ContactStatusReject || contact.Status == dbr.ContactStatusDeleted {
		contact.Status = dbr.ContactStatusRequesting
		contact.IsDel = 0 // remove deleted flag if needed
		if err = contact.UpdateInUnscoped(db); err != nil {
			logrus.Errorf("contactManageServant.RequestingFriend update exsit contact err:%s", err)
			return
		}
	}

	msg := &dbr.Message{
		SenderUserID:   userId,
		ReceiverUserID: friendId,
		Type:           dbr.MsgTypeRequestingFriend,
		Brief:          "请求添加好友，并附言:",
		Content:        greetings,
		ReplyID:        int64(dbr.ContactStatusRequesting),
	}
	if _, err = msg.Create(db); err != nil {
		logrus.Errorf("contactManageServant.RequestingFriend create message err:%s", err)
		return
	}
	return nil
}

func (s *contactManageServant) AddFriend(userId int64, friendId int64) (err error) {
	db := s.db.Begin()
	defer func() {
		if err == nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()

	contact := &dbr.Contact{
		UserId:   friendId,
		FriendId: userId,
	}
	if contact, err = contact.GetByUserFriend(db); err != nil {
		return
	}
	// 如果还不是请求好友，啥也不干
	if contact.Status != dbr.ContactStatusRequesting {
		logrus.Debugf("contactManageServant.AddFriend not reuesting status now so skip")
		return nil
	}
	contact.Status = dbr.ContactStatusAgree
	if err = contact.Update(db); err != nil {
		return err
	}

	contact, err = s.fetchOrNewContact(db, userId, friendId, dbr.ContactStatusAgree)
	if err != nil {
		return
	}

	// 如果已经好友，啥也不干
	if contact.Status != dbr.ContactStatusAgree {
		contact.Status = dbr.ContactStatusAgree
		contact.IsDel = 0 // remove deleted flag
		if err = contact.UpdateInUnscoped(db); err != nil {
			logrus.Errorf("contactManageServant.AddFriend update contact err:%s", err)
			return
		}
	}

	args := []any{userId, friendId, friendId, userId, dbr.MsgTypeRequestingFriend, dbr.ContactStatusRequesting}
	msgs, e := (&dbr.Message{}).FetchBy(db, dbr.Predicates{
		"((sender_user_id = ? AND receiver_user_id = ?) OR (sender_user_id = ? AND receiver_user_id = ?)) AND type = ? AND reply_id = ?": args,
	})
	if e != nil {
		err = e
		return
	}
	for _, msg := range msgs {
		msg.ReplyID = int64(dbr.ContactStatusAgree)
		if err = msg.Update(db); err != nil {
			return
		}
	}
	return nil
}

func (s *contactManageServant) RejectFriend(userId int64, friendId int64) (err error) {
	db := s.db.Begin()
	defer func() {
		if err == nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()

	contact := &dbr.Contact{
		UserId:   friendId,
		FriendId: userId,
	}
	if contact, err = contact.GetByUserFriend(db); err != nil {
		return
	}
	// 如果还不是请求好友，啥也不干
	if contact.Status != dbr.ContactStatusRequesting {
		return nil
	}
	contact.Status = dbr.ContactStatusReject
	if err = contact.Update(db); err != nil {
		return err
	}

	args := []any{friendId, userId, dbr.MsgTypeRequestingFriend, dbr.ContactStatusRequesting}
	msgs, e := (&dbr.Message{}).FetchBy(db, dbr.Predicates{
		"sender_user_id = ? AND receiver_user_id = ? AND type = ? AND reply_id = ?": args,
	})
	if e != nil {
		err = e
		return
	}
	for _, msg := range msgs {
		msg.ReplyID = int64(dbr.ContactStatusReject)
		if err = msg.Update(db); err != nil {
			return
		}
	}
	return nil
}

func (s *contactManageServant) DeleteFriend(userId int64, friendId int64) (err error) {
	db := s.db.Begin()
	defer func() {
		if err == nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()

	contact := &dbr.Contact{
		UserId:   userId,
		FriendId: friendId,
	}
	contacts, e := contact.FetchByUserFriendAll(db)
	if e != nil {
		return e
	}

	for _, contact := range contacts {
		// 如果还不是好友，啥也不干
		if contact.Status != dbr.ContactStatusAgree {
			continue
		}
		contact.Status = dbr.ContactStatusDeleted
		contact.DeletedOn = time.Now().Unix()
		contact.IsDel = 1
		if err = contact.Update(db); err != nil {
			return
		}
	}
	return nil
}

func (s *contactManageServant) GetContacts(userId int64, offset int, limit int) (*core.ContactList, error) {
	contact := &dbr.Contact{}
	condition := dbr.ConditionsT{
		"user_id": userId,
		"status":  dbr.ContactStatusAgree,
	}
	contacts, err := contact.List(s.db, condition, offset, limit)
	if err != nil {
		return nil, err
	}
	total, err := contact.Count(s.db, condition)
	if err != nil {
		return nil, err
	}
	resp := &core.ContactList{
		Contacts: make([]core.ContactItem, 0, len(contacts)),
		Total:    total,
	}
	for _, c := range contacts {
		if c.User != nil {
			resp.Contacts = append(resp.Contacts, core.ContactItem{
				UserId:   c.FriendId,
				UserName: c.User.Username,
				Nickname: c.User.Nickname,
				Avatar:   c.User.Avatar,
				Phone:    c.User.Phone,
			})
		}
	}
	return resp, nil
}

func (s *contactManageServant) IsFriend(userId int64, friendId int64) bool {
	contact := &dbr.Contact{
		UserId:   friendId,
		FriendId: userId,
	}
	contact, err := contact.GetByUserFriend(s.db)
	if err == nil && contact.Status == dbr.ContactStatusAgree {
		return true
	}
	return false
}
