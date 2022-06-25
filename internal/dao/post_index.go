package dao

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/sirupsen/logrus"
)

func (d *dataServant) IndexPosts(user *model.User, offset int, limit int) ([]*model.PostFormated, error) {
	if d.useCacheIndex {
		if posts, err := d.cacheIndex.IndexPosts(user, offset, limit); err == nil {
			logrus.Debugf("get index posts from cached")
			return posts, nil
		}
	}
	logrus.Debugf("get index posts from database but useCacheIndex: %t", d.useCacheIndex)
	return d.getIndexPostsFunc(user, offset, limit)
}

// getIndexPosts TODO: 未来可能根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (d *dataServant) getIndexPosts(user *model.User, offset int, limit int) ([]*model.PostFormated, error) {
	predicates := model.Predicates{
		"ORDER": types.AnySlice{"is_top DESC, latest_replied_on DESC"},
	}
	if user == nil {
		predicates["visibility = ?"] = types.AnySlice{model.PostVisitPublic}
	} else if !user.IsAdmin {
		frienddIds := d.ams.GetFriendIds(user.ID)
		if len(frienddIds) > 0 {
			args := types.AnySlice{model.PostVisitPublic, model.PostVisitPrivate, user.ID, model.PostVisitFriend, frienddIds}
			predicates["visibility = ? OR (visibility = ? AND user_id = ?) OR (visibility = ? AND user_id IN ?"] = args
		} else {
			// TODO: 目前不处理没朋友情况，默认用户与世界都是朋友，但是目前无从知晓朋友id，所以这里特殊处理，后续完善
			args := types.AnySlice{model.PostVisitPrivate, user.ID, model.PostVisitPublic, model.PostVisitFriend}
			predicates["(visibility = ? AND user_id = ?) OR visibility = ? OR visibility = ?"] = args
		}
	}

	posts, err := (&model.Post{}).Fetch(d.engine, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("dataServant.getIndexPosts err: %v", err)
		return nil, err
	}

	return d.MergePosts(posts)
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (d *dataServant) simpleCacheIndexGetPosts(_user *model.User, offset int, limit int) ([]*model.PostFormated, error) {
	predicates := model.Predicates{
		"visibility IN ?": types.AnySlice{[]model.PostVisibleT{model.PostVisitPublic, model.PostVisitPublic}},
		"ORDER":           types.AnySlice{"is_top DESC, latest_replied_on DESC"},
	}

	posts, err := (&model.Post{}).Fetch(d.engine, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("dataServant.simpleCacheIndexGetPosts err: %v", err)
		return nil, err
	}

	return d.MergePosts(posts)
}

func (d *dataServant) indexActive(act core.IndexActionT) {
	if d.useCacheIndex {
		d.cacheIndex.SendAction(act)
	}
}
