// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"encoding/json"
	"fmt"

	"github.com/alimy/tryst/event"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/sirupsen/logrus"
)

const (
	_tweetActionCreate uint8 = iota
	_tweetActionDelete
)

const (
	_commentActionCreate uint8 = iota
	_commentActionDelete
	_commentActionThumbsUp
	_commentActionThumbsDown
	_commentActionReplyCreate
	_commentActionReplyDelete
	_commentActionReplyThumbsUp
	_commentActionReplyThumbsDown
	_commentActionHighlight
)

const (
	_messageActionCreate uint8 = iota
	_messageActionRead
	_messageActionFollow
	_messageActionSendWhisper
)

const (
	_trendsActionCreateTweet uint8 = iota
	_trendsActionDeleteTweet
	_trendsActionFollowUser
	_trendsActionUnfollowUser
	_trendsActionAddFriend
	_trendsActionDeleteFriend
)

type cacheUnreadMsgEvent struct {
	event.UnimplementedEvent
	ds  core.DataService
	wc  core.WebCache
	uid int64
}

type createMessageEvent struct {
	event.UnimplementedEvent
	ds      core.DataService
	wc      core.WebCache
	message *ms.Message
}

type tweetActionEvent struct {
	event.UnimplementedEvent
	ac       core.AppCache
	userId   int64
	username string
	action   uint8
}

type commentActionEvent struct {
	event.UnimplementedEvent
	ds        core.DataService
	ac        core.AppCache
	tweetId   int64
	commentId int64
	action    uint8
}

type messageActionEvent struct {
	event.UnimplementedEvent
	wc     core.WebCache
	action uint8
	userId []int64
}

type trendsActionEvent struct {
	event.UnimplementedEvent
	ac      core.AppCache
	ds      core.DataService
	action  uint8
	userIds []int64
}

type changeUserEvent struct {
	*cache.BaseCacheEvent
	userId   int64
	username string
}

func onChangeUsernameEvent(id int64, name string) {
	events.OnEvent(&changeUserEvent{
		BaseCacheEvent: cache.NewBaseCacheEvent(_ac),
		userId:         id,
		username:       name,
	})
}

func onTrendsActionEvent(action uint8, userIds ...int64) {
	events.OnEvent(&trendsActionEvent{
		ac:      _ac,
		ds:      _ds,
		action:  action,
		userIds: userIds,
	})
}

func onTweetActionEvent(action uint8, userId int64, username string) {
	events.OnEvent(&tweetActionEvent{
		ac:       _ac,
		action:   action,
		userId:   userId,
		username: username,
	})
}

func onMessageActionEvent(action uint8, userIds ...int64) {
	events.OnEvent(&messageActionEvent{
		wc:     _wc,
		action: action,
		userId: userIds,
	})
}

func onCommentActionEvent(tweetId int64, commentId int64, action uint8) {
	events.OnEvent(&commentActionEvent{
		ds:        _ds,
		ac:        _ac,
		tweetId:   tweetId,
		commentId: commentId,
		action:    action,
	})
}

func onCacheUnreadMsgEvent(uid int64) {
	events.OnEvent(&cacheUnreadMsgEvent{
		ds:  _ds,
		wc:  _wc,
		uid: uid,
	})
}

func onCreateMessageEvent(data *ms.Message) {
	events.OnEvent(&createMessageEvent{
		ds:      _ds,
		wc:      _wc,
		message: data,
	})
}

func (e *cacheUnreadMsgEvent) Name() string {
	return "cacheUnreadMsgEvent"
}

func (e *cacheUnreadMsgEvent) Action() error {
	if e.wc.ExistUnreadMsgCountResp(e.uid) {
		// do nothing
		return nil
	}
	count, err := e.ds.GetUnreadCount(e.uid)
	if err != nil {
		return fmt.Errorf("cacheUnreadMsgEvent action occurs error: %w", err)
	}
	resp := &joint.JsonResp{
		Code: 0,
		Msg:  "success",
		Data: &web.GetUnreadMsgCountResp{
			Count: count,
		},
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cacheUnreadMsgEvent action marshal resp occurs error: %w", err)
	}
	if err = e.wc.PutUnreadMsgCountResp(e.uid, data); err != nil {
		return fmt.Errorf("cacheUnreadMsgEvent action put resp data to redis cache occurs error: %w", err)
	}
	return nil
}

func (e *createMessageEvent) Name() string {
	return "createMessageEvent"
}

func (e *createMessageEvent) Action() (err error) {
	if _, err = e.ds.CreateMessage(e.message); err == nil {
		err = e.wc.DelUnreadMsgCountResp(e.message.ReceiverUserID)
	}
	return
}

func (e *commentActionEvent) Name() string {
	return "updateCommentMetricEvent"
}

func (e *commentActionEvent) Action() (err error) {
	// logrus.Debugf("trigger commentActionEvent action commentId[%d]", e.commentId)
	switch e.action {
	case _commentActionCreate:
		err = e.ds.AddCommentMetric(e.commentId)
		e.expireAllStyleComments()
	case _commentActionDelete:
		err = e.ds.DeleteCommentMetric(e.commentId)
		e.expireAllStyleComments()
	case _commentActionReplyCreate, _commentActionReplyDelete:
		err = e.updateCommentMetric()
		e.expireAllStyleComments()
	case _commentActionThumbsUp, _commentActionThumbsDown:
		err = e.updateCommentMetric()
		e.expireHotsComments()
	case _commentActionHighlight:
		e.expireAllStyleComments()
	default:
		// nothing
	}
	return
}

func (e *commentActionEvent) expireHotsComments() {
	e.ac.DelAny(fmt.Sprintf("%s%d:%s:*", conf.PrefixTweetComment, e.tweetId, conf.InfixCommentHots))
}

func (e *commentActionEvent) expireAllStyleComments() {
	e.ac.DelAny(fmt.Sprintf("%s%d:*", conf.PrefixTweetComment, e.tweetId))
}

func (e *commentActionEvent) updateCommentMetric() error {
	// logrus.Debug("trigger commentActionEvent action[updateCommentMetric]")
	comment, err := e.ds.GetCommentByID(e.commentId)
	if err != nil {
		return err
	}
	e.ds.UpdateCommentMetric(&cs.CommentMetric{
		CommentId:       e.commentId,
		ReplyCount:      comment.ReplyCount,
		ThumbsUpCount:   comment.ThumbsUpCount,
		ThumbsDownCount: comment.ThumbsDownCount,
	})
	return nil
}

func (e *messageActionEvent) Name() string {
	return "expireMessagesEvent"
}

func (e *messageActionEvent) Action() (err error) {
	for _, userId := range e.userId {
		switch e.action {
		case _messageActionRead,
			_messageActionSendWhisper:
			// 清除未读消息数缓存，不需要处理错误
			e.wc.DelUnreadMsgCountResp(userId)
		case _messageActionCreate,
			_messageActionFollow:
			fallthrough
		default:
			// TODO
		}
		//清除该用户所有消息缓存
		err = e.wc.DelAny(fmt.Sprintf("%s%d:*", conf.PrefixMessages, userId))
	}
	return
}

func (e *trendsActionEvent) Name() string {
	return "trendsActionEvent"
}

func (e *trendsActionEvent) Action() (err error) {
	switch e.action {
	case _trendsActionCreateTweet:
		logrus.Debug("trigger trendsActionEvent by create tweet ")
		e.updateUserMetric(cs.MetricActionCreateTweet)
		e.expireFriendTrends()
	case _trendsActionDeleteTweet:
		logrus.Debug("trigger trendsActionEvent by delete tweet ")
		e.updateUserMetric(cs.MetricActionDeleteTweet)
		e.expireFriendTrends()
	case _trendsActionAddFriend, _trendsActionDeleteFriend,
		_trendsActionFollowUser, _trendsActionUnfollowUser:
		e.expireMyTrends()
	default:
		// nothing
	}
	return
}

func (e *trendsActionEvent) updateUserMetric(action uint8) {
	for _, userId := range e.userIds {
		e.ds.UpdateUserMetric(userId, action)
	}
}

func (e *trendsActionEvent) expireFriendTrends() {
	for _, userId := range e.userIds {
		if friendIds, err := e.ds.MyFriendIds(userId); err == nil {
			for _, id := range friendIds {
				e.ac.DelAny(fmt.Sprintf("%s%d:*", conf.PrefixIdxTrends, id))
			}
		}
	}
}

func (e *trendsActionEvent) expireMyTrends() {
	for _, userId := range e.userIds {
		e.ac.DelAny(fmt.Sprintf("%s%d:*", conf.PrefixIdxTrends, userId))
	}
}

func (e *tweetActionEvent) Name() string {
	return "tweetActionEvent"
}

func (e *tweetActionEvent) Action() (err error) {
	switch e.action {
	case _tweetActionCreate, _tweetActionDelete:
		e.ac.Delete(conf.KeyUserProfileByName.Get(e.username))
	default:
		// nothing
	}
	return
}

func (e *changeUserEvent) Name() string {
	return "changeUserEvent"
}

func (e *changeUserEvent) Action() error {
	return e.ExpireUserData(e.userId, e.username)
}
