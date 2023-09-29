// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/slonik/sqlc/auto/pgc"
	"github.com/sirupsen/logrus"
)

var (
	_ core.ContactManageService = (*contactManageSrv)(nil)
)

type contactManageSrv struct {
	*pgcSrv
}

func (s *contactManageSrv) RequestingFriend(userId int64, friendId int64, greetings string) (err error) {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		now := time.Now().Unix()
		contact, err := q.UpsertContact(c, &pgc.UpsertContactParams{
			UserID:    userId,
			FriendID:  friendId,
			Status:    int16(cs.ContactStatusRequesting),
			CreatedOn: now,
		})
		if err != nil {
			return err
		}
		// 如果已经好友，啥也不干
		if contact.Status == int16(cs.ContactStatusAgree) {
			return nil
		} else if contact.Status == int16(cs.ContactStatusReject) || contact.Status == int16(cs.ContactStatusDeleted) {
			contact.Status = int16(cs.ContactStatusRequesting)
			if err = q.FreshContactStatus(c, &pgc.FreshContactStatusParams{
				Status:     contact.Status,
				ModifiedOn: now,
				ID:         contact.ID,
			}); err != nil {
				return err
			}
		}
		if _, err = q.CreateMessage(c, &pgc.CreateMessageParams{
			SenderUserID:   userId,
			ReceiverUserID: friendId,
			Type:           int16(ms.MsgTypeRequestingFriend),
			Brief:          "请求添加好友，并附言:",
			Content:        greetings,
			ReplyID:        int64(cs.ContactStatusRequesting),
			CreatedOn:      now,
		}); err != nil {
			return err
		}
		return nil
	})
}

func (s *contactManageSrv) AddFriend(userId int64, friendId int64) (err error) {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		contact, err := q.GetUserFriend(c, &pgc.GetUserFriendParams{
			UserID:   friendId,
			FriendID: userId,
		})
		if err != nil {
			return err
		}
		// 如果还不是请求好友，啥也不干
		if contact.Status != int16(cs.ContactStatusRequesting) {
			logrus.Debugf("contactManageSrv.AddFriend not reuesting status now so skip")
			return nil
		}
		now := time.Now().Unix()
		if err = q.FreshContactStatus(c, &pgc.FreshContactStatusParams{
			Status:     int16(cs.ContactStatusAgree),
			ModifiedOn: now,
			ID:         contact.ID,
		}); err != nil {
			return err
		}
		var row *pgc.UpsertContactRow
		row, err = q.UpsertContact(c, &pgc.UpsertContactParams{
			UserID:    userId,
			FriendID:  friendId,
			Status:    int16(cs.ContactStatusAgree),
			CreatedOn: now,
		})
		if err != nil {
			return err
		}
		// 如果已经好友，啥也不干
		if row.Status != int16(cs.ContactStatusAgree) {
			if err = q.FreshContactStatus(c, &pgc.FreshContactStatusParams{
				Status:     int16(cs.ContactStatusAgree),
				ModifiedOn: now,
				ID:         row.ID,
			}); err != nil {
				return err
			}
		}
		if err = q.AddFriendMsgsUpdate(c, &pgc.AddFriendMsgsUpdateParams{
			ReplyID:        int64(cs.ContactStatusAgree),
			ModifiedOn:     now,
			SenderUserID:   userId,
			ReceiverUserID: friendId,
			Type:           int16(ms.MsgTypeRequestingFriend),
			ReplyID_2:      int64(cs.ContactStatusRequesting),
		}); err != nil {
			return err
		}
		return nil
	})
}

func (s *contactManageSrv) RejectFriend(userId int64, friendId int64) (err error) {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		contact, err := q.GetUserFriend(c, &pgc.GetUserFriendParams{
			UserID:   friendId,
			FriendID: userId,
		})
		if err != nil {
			return err
		}
		// 如果还不是请求好友，啥也不干
		if contact.Status != int16(cs.ContactStatusRequesting) {
			return nil
		}
		now := time.Now().Unix()
		if err = q.FreshContactStatus(c, &pgc.FreshContactStatusParams{
			Status:     int16(cs.ContactStatusReject),
			ModifiedOn: now,
			ID:         contact.ID,
		}); err != nil {
			return err
		}
		if err = q.RejectFriendMsgsUpdate(c, &pgc.RejectFriendMsgsUpdateParams{
			ReplyID:        int64(cs.ContactStatusReject),
			ModifiedOn:     now,
			SenderUserID:   friendId,
			ReceiverUserID: userId,
			Type:           int16(ms.MsgTypeRequestingFriend),
			ReplyID_2:      int64(cs.ContactStatusRequesting),
		}); err != nil {
			return err
		}
		return nil
	})
}

func (s *contactManageSrv) DeleteFriend(userId int64, friendId int64) (err error) {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		contacts, err := q.GetContacts(c, &pgc.GetContactsParams{
			UserID:   userId,
			FriendID: friendId,
		})
		if err != nil {
			return err
		}
		ids := make([]int64, 0, 2)
		for _, contact := range contacts {
			// 如果还不是好友，啥也不干
			if contact.Status != int16(cs.ContactStatusAgree) {
				continue
			}
			ids = append(ids, contact.ID)
		}
		if err = q.DeleteFriendByIds(c, &pgc.DeleteFriendByIdsParams{
			DeletedOn: time.Now().Unix(),
			Ids:       ids,
		}); err != nil {
			return err
		}
		return nil
	})
}

func (s *contactManageSrv) GetContacts(userId int64, offset int, limit int) (res *ms.ContactList, err error) {
	res = &ms.ContactList{}
	if res.Total, err = s.q.CountFriendsById(context.Background(), userId); err == nil && res.Total > 0 {
		var rows []*pgc.ListFriendRow
		rows, err = s.q.ListFriend(context.Background(), &pgc.ListFriendParams{
			UserID: userId,
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			return
		}
		for _, contact := range rows {
			res.Contacts = append(res.Contacts, ms.ContactItem{
				UserId:   contact.UserID,
				Username: contact.Username,
				Nickname: contact.Nickname,
				Avatar:   contact.Avatar,
				Phone:    contact.Phone,
			})
		}
	}
	return
}

func (s *contactManageSrv) IsFriend(userId int64, friendId int64) bool {
	ok, _ := s.q.IsFriend(context.Background(), &pgc.IsFriendParams{
		UserID:   userId,
		FriendID: friendId,
	})
	return ok
}

func newContactManageService(db *pgx.Conn) core.ContactManageService {
	return &contactManageSrv{
		pgcSrv: newPgcSrv(db),
	}
}
