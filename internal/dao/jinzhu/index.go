package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/model/rest"
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ core.IndexPostsService = (*indexPostsServant)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsServant)(nil)
)

type indexPostsServant struct {
	ams core.AuthorizationManageService
	ths core.TweetHelpService
	db  *gorm.DB
}

type simpleIndexPostsServant struct {
	ths core.TweetHelpService
	db  *gorm.DB
}

func newIndexPostsService(db *gorm.DB) core.IndexPostsService {
	return &indexPostsServant{
		ams: NewAuthorizationManageService(),
		ths: newTweetHelpService(db),
		db:  db,
	}
}

func newSimpleIndexPostsService(db *gorm.DB) core.IndexPostsService {
	return &simpleIndexPostsServant{
		ths: newTweetHelpService(db),
		db:  db,
	}
}

// IndexPosts 根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (s *indexPostsServant) IndexPosts(user *model.User, offset int, limit int) (*rest.IndexTweetsResp, error) {
	predicates := model.Predicates{
		"ORDER": types.AnySlice{"is_top DESC, latest_replied_on DESC"},
	}
	if user == nil {
		predicates["visibility = ?"] = []types.Any{model.PostVisitPublic}
	} else if !user.IsAdmin {
		friendIds, _ := s.ams.BeFriendIds(user.ID)
		friendIds = append(friendIds, user.ID)
		args := types.AnySlice{model.PostVisitPublic, model.PostVisitPrivate, user.ID, model.PostVisitFriend, friendIds}
		predicates["visibility = ? OR (visibility = ? AND user_id = ?) OR (visibility = ? AND user_id IN ?)"] = args
	}

	posts, err := (&model.Post{}).Fetch(s.db, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("gormIndexPostsServant.IndexPosts err: %v", err)
		return nil, err
	}
	formatPosts, err := s.ths.MergePosts(posts)
	if err != nil {
		return nil, err
	}

	total, err := (&model.Post{}).CountBy(s.db, predicates)
	if err != nil {
		return nil, err
	}

	return &rest.IndexTweetsResp{
		Tweets: formatPosts,
		Total:  total,
	}, nil
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (s *simpleIndexPostsServant) IndexPosts(_user *model.User, offset int, limit int) (*rest.IndexTweetsResp, error) {
	predicates := model.Predicates{
		"visibility = ?": []types.Any{model.PostVisitPublic},
		"ORDER":          []types.Any{"is_top DESC, latest_replied_on DESC"},
	}

	posts, err := (&model.Post{}).Fetch(s.db, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("gormSimpleIndexPostsServant.IndexPosts err: %v", err)
		return nil, err
	}

	formatPosts, err := s.ths.MergePosts(posts)
	if err != nil {
		return nil, err
	}

	total, err := (&model.Post{}).CountBy(s.db, predicates)
	if err != nil {
		return nil, err
	}

	return &rest.IndexTweetsResp{
		Tweets: formatPosts,
		Total:  total,
	}, nil
}
