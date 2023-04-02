// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"context"

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

func NewIndexAction(act IdxAct, post *dbr.Post) *IndexAction {
	return &IndexAction{
		Act:  act,
		Post: post,
	}
}

// CacheIndexService cache index service interface
type CacheIndexService interface {
	IndexPostsService

	SendAction(act IdxAct, post *dbr.Post)
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
