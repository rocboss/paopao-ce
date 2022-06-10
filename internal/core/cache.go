package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

const (
	IdxActNop IndexActionT = iota + 1
	IdxActCreatePost
	IdxActUpdatePost
	IdxActDeletePost
	IdxActStickPost
)

type IndexActionT uint8

func (a IndexActionT) String() string {
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
	default:
		return "unknow action"
	}
}

// CacheIndexService cache index service interface
type CacheIndexService interface {
	IndexPosts(offset int, limit int) ([]*model.PostFormated, bool)
	SendAction(active IndexActionT)
}
