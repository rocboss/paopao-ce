package dao

import (
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/sirupsen/logrus"
)

func (d *dataServant) IndexPosts(offset int, limit int) ([]*model.PostFormated, error) {
	if d.useCacheIndex {
		posts := d.atomicIndex.Load().([]*model.PostFormated)
		start := offset * limit
		end := start + limit
		if len(posts) >= end {
			logrus.Debugln("get index posts from cached")
			return posts[start:end], nil
		}
	}
	logrus.Debugf("get index posts from database but useCacheIndex: %t", d.useCacheIndex)
	return d.getIndexPosts(offset, limit)
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

func (d *dataServant) getIndexPosts(offset int, limit int) ([]*model.PostFormated, error) {
	posts, err := (&model.Post{}).List(d.engine, &model.ConditionsT{
		"ORDER": "is_top DESC, latest_replied_on DESC",
	}, offset, limit)
	if err != nil {
		logrus.Debugf("getIndexPosts err: %v", err)
		return nil, err
	}
	return d.MergePosts(posts)
}

func (d *dataServant) indexActive(act indexActionT) {
	select {
	case d.indexActionCh <- act:
		logrus.Debugf("send indexAction by chan: %s", act)
	default:
		go func(ch chan<- indexActionT, act indexActionT) {
			logrus.Debugf("send indexAction by goroutine: %s", act)
			ch <- act
		}(d.indexActionCh, act)
	}
}

func (d *dataServant) startIndexPosts() {
	var err error
	for {
		select {
		case <-d.checkTick.C:
			if len(d.indexPosts) == 0 {
				logrus.Debugf("index posts by checkTick")
				if d.indexPosts, err = d.getIndexPosts(0, d.maxIndexSize); err == nil {
					d.atomicIndex.Store(d.indexPosts)
				} else {
					logrus.Errorf("get index posts err: %v", err)
				}
			}
		case <-d.expireIndexTick.C:
			logrus.Debugf("expire index posts by expireIndexTick")
			if len(d.indexPosts) != 0 {
				d.indexPosts = nil
				d.atomicIndex.Store(d.indexPosts)
			}
		case action := <-d.indexActionCh:
			switch action {
			case idxActCreatePost, idxActUpdatePost, idxActDeletePost, idxActStickPost:
				logrus.Debugf("remove index posts by action %s", action)
				// prevent many update post in least time
				if len(d.indexPosts) != 0 {
					d.indexPosts = nil
					d.atomicIndex.Store(d.indexPosts)
				}
			default:
				// nop
			}
		}
	}
}
