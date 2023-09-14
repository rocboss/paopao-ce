// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

const (
	IdxActNop IdxAct = iota + 1
	IdxActCreatePost
	IdxActUpdatePost
	IdxActDeletePost
	IdxActStickPost
	IdxActVisiblePost
)

type IdxAct uint8

type IndexAction struct {
	Act  IdxAct
	Post *dbr.Post
}

type IndexActionA struct {
	Act   IdxAct
	Tweet *cs.TweetInfo
}

func (a IdxAct) String() string {
	switch a {
	case IdxActNop:
		return "no operator"
	case IdxActCreatePost:
		return "create post"
	case IdxActUpdatePost:
		return "update post"
	case IdxActDeletePost:
		return "delete post"
	case IdxActStickPost:
		return "stick post"
	case IdxActVisiblePost:
		return "visible post"
	default:
		return "unknow action"
	}
}

func NewIndexAction(act IdxAct, post *ms.Post) *IndexAction {
	return &IndexAction{
		Act:  act,
		Post: post,
	}
}

func NewIndexActionA(act IdxAct, tweet *cs.TweetInfo) *IndexActionA {
	return &IndexActionA{
		Act:   act,
		Tweet: tweet,
	}
}

// CacheIndexService cache index service interface
type CacheIndexService interface {
	IndexPostsService

	SendAction(act IdxAct, post *dbr.Post)
}

// CacheIndexServantA cache index service interface
type CacheIndexServantA interface {
	IndexPostsServantA

	SendAction(act IdxAct, tweet *cs.TweetInfo)
}

// RedisCache memory cache by Redis
type RedisCache interface {
	SetPushToSearchJob(ctx context.Context) error
	DelPushToSearchJob(ctx context.Context) error
	SetImgCaptcha(ctx context.Context, id string, value string) error
	GetImgCaptcha(ctx context.Context, id string) (string, error)
	DelImgCaptcha(ctx context.Context, id string) error
	GetCountSmsCaptcha(ctx context.Context, phone string) (int64, error)
	IncrCountSmsCaptcha(ctx context.Context, phone string) error
	GetCountLoginErr(ctx context.Context, id int64) (int64, error)
	DelCountLoginErr(ctx context.Context, id int64) error
	IncrCountLoginErr(ctx context.Context, id int64) error
	GetCountWhisper(ctx context.Context, uid int64) (int64, error)
	IncrCountWhisper(ctx context.Context, uid int64) error
	SetRechargeStatus(ctx context.Context, tradeNo string) error
	DelRechargeStatus(ctx context.Context, tradeNo string) error
}

type AppCache interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte, ex int64) error
	SetNx(key string, data []byte, ex int64) error
	Delete(key ...string) error
	DelAny(pattern string) error
	Exist(key string) bool
}

type WebCache interface {
	AppCache
	GetUnreadMsgCountResp(uid int64) ([]byte, error)
	PutUnreadMsgCountResp(uid int64, data []byte) error
	DelUnreadMsgCountResp(uid int64) error
	ExistUnreadMsgCountResp(uid int64) bool
}
