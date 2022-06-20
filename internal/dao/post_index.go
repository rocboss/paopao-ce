package dao

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/sirupsen/logrus"
)

func (d *dataServant) IndexPosts(userId int64, offset int, limit int) ([]*model.PostFormated, error) {
	if d.useCacheIndex {
		if posts, err := d.cacheIndex.IndexPosts(userId, offset, limit); err == nil {
			logrus.Debugf("get index posts from cached by userId: %d", userId)
			return posts, nil
		}
	}
	logrus.Debugf("get index posts from database but useCacheIndex: %t", d.useCacheIndex)
	return d.getIndexPosts(userId, offset, limit)
}

func (d *dataServant) MergePosts(posts []*model.Post) ([]*model.PostFormated, error) {
	postIds := make([]int64, 0, len(posts))
	userIds := make([]int64, 0, len(posts))
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}

	postContents, err := d.GetPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := d.GetUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int64]*model.UserFormated, len(users))
	for _, user := range users {
		userMap[user.ID] = user.Format()
	}

	contentMap := make(map[int64][]*model.PostContentFormated, len(postContents))
	for _, content := range postContents {
		contentMap[content.PostID] = append(contentMap[content.PostID], content.Format())
	}

	// 数据整合
	postsFormated := make([]*model.PostFormated, 0, len(posts))
	for _, post := range posts {
		postFormated := post.Format()
		postFormated.User = userMap[post.UserID]
		postFormated.Contents = contentMap[post.ID]
		postsFormated = append(postsFormated, postFormated)
	}
	return postsFormated, nil
}

// getIndexPosts _userId保留未来使用
// TODO: 未来可能根据userId查询广场推文列表，简单做到不同用户的主页都是不同的；
func (d *dataServant) getIndexPosts(_userId int64, offset int, limit int) ([]*model.PostFormated, error) {
	posts, err := (&model.Post{}).List(d.engine, &model.ConditionsT{
		"visibility IN ?": []model.PostVisibleT{model.PostVisitPublic, model.PostVisitFriend},
		"ORDER":           "is_top DESC, latest_replied_on DESC",
	}, offset, limit)
	if err != nil {
		logrus.Debugf("getIndexPosts err: %v", err)
		return nil, err
	}
	return d.MergePosts(posts)
}

func (d *dataServant) indexActive(act core.IndexActionT) {
	if d.useCacheIndex {
		d.cacheIndex.SendAction(act)
	}
}
