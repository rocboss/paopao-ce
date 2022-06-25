package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

type IndexPostsService interface {
	IndexPosts(user *model.User, offset int, limit int) ([]*model.PostFormated, error)
}
