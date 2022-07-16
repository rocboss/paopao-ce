package core

import "github.com/rocboss/paopao-ce/internal/model"

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
	Post *model.Post
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

func NewIndexAction(act IdxAct, post *model.Post) *IndexAction {
	return &IndexAction{
		Act:  act,
		Post: post,
	}
}

// CacheIndexService cache index service interface
type CacheIndexService interface {
	IndexPostsService

	SendAction(act IdxAct, post *model.Post)
}
