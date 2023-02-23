// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

const (
	SearchTypeDefault SearchType = "search"
	SearchTypeTag     SearchType = "tag"
)

const (
	PostVisitPublic  = dbr.PostVisitPublic
	PostVisitPrivate = dbr.PostVisitPrivate
	PostVisitFriend  = dbr.PostVisitFriend
	PostVisitInvalid = dbr.PostVisitInvalid
)

type (
	PostVisibleT = dbr.PostVisibleT

	SearchType string

	QueryReq struct {
		Query      string
		Visibility []PostVisibleT
		Type       SearchType
	}

	QueryResp struct {
		Items []*PostFormated
		Total int64
	}

	TsDocItem struct {
		Post    *Post
		Content string
	}
)

// TweetSearchService tweet search service interface
type TweetSearchService interface {
	IndexName() string
	AddDocuments(data []TsDocItem, primaryKey ...string) (bool, error)
	DeleteDocuments(identifiers []string) error
	Search(user *User, q *QueryReq, offset, limit int) (*QueryResp, error)
}
