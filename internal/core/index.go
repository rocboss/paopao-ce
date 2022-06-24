package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

type IndexPostsService interface {
	IndexPosts(userId int64, offset int, limit int) ([]*model.PostFormated, error)
}
