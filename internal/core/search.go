// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

const (
	SearchTypeDefault SearchType = "search"
	SearchTypeTag     SearchType = "tag"
)

const (
	PostVisitPublic    = dbr.PostVisitPublic
	PostVisitPrivate   = dbr.PostVisitPrivate
	PostVisitFriend    = dbr.PostVisitFriend
	PostVisitFollowing = dbr.PostVisitFollowing
	PostVisitInvalid   = dbr.PostVisitInvalid
)

type (
	// PostVisibleT 可访问类型，0公开，1私密，2好友
	PostVisibleT = dbr.PostVisibleT

	SearchType string

	QueryReq struct {
		Query      string
		Visibility []PostVisibleT
		Type       SearchType
	}

	QueryResp struct {
		Items []*ms.PostFormated
		Total int64
	}

	TsDocItem struct {
		Post    *ms.Post
		Content string
	}
)

// TweetSearchService tweet search service interface
type TweetSearchService interface {
	IndexName() string
	AddDocuments(data []TsDocItem, primaryKey ...string) (bool, error)
	DeleteDocuments(identifiers []string) error
	Search(user *ms.User, q *QueryReq, offset, limit int) (*QueryResp, error)
}
