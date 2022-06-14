package dao

import (
	"strings"
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
)

func (d *dataServant) CreatePost(post *model.Post) (*model.Post, error) {
	post.LatestRepliedOn = time.Now().Unix()
	p, err := post.Create(d.engine)
	if err != nil {
		return nil, err
	}
	d.indexActive(core.IdxActCreatePost)
	return p, nil
}

func (d *dataServant) DeletePost(post *model.Post) error {
	if err := post.Delete(d.engine); err != nil {
		return err
	}
	d.indexActive(core.IdxActDeletePost)
	return nil
}

func (d *dataServant) LockPost(post *model.Post) error {
	post.IsLock = 1 - post.IsLock
	return post.Update(d.engine)
}

func (d *dataServant) StickPost(post *model.Post) error {
	post.IsTop = 1 - post.IsTop
	if err := post.Update(d.engine); err != nil {
		return err
	}
	d.indexActive(core.IdxActStickPost)
	return nil
}

func (d *dataServant) VisiblePost(post *model.Post, visibility model.PostVisibleT) error {
	oldVisibility := post.Visibility
	post.Visibility = visibility
	// TODO: 这个判断是否可以不要呢
	if oldVisibility == visibility {
		return nil
	}
	// 私密推文 特殊处理
	if visibility == model.PostVisitPrivate {
		// 强制取消置顶
		// TODO: 置顶推文用户是否有权设置成私密？ 后续完善
		post.IsTop = 0
	}
	db := d.engine.Begin()
	err := post.Update(db)
	if err != nil {
		db.Rollback()
		return err
	}

	// tag处理
	tags := strings.Split(post.Tags, ",")
	for _, t := range tags {
		tag := &model.Tag{
			Tag: t,
		}
		// TODO: 暂时宽松不处理错误，这里或许可以有优化，后续完善
		if oldVisibility == model.PostVisitPrivate {
			// 从私密转为非私密才需要重新创建tag
			d.createTag(db, tag)
		} else if visibility == model.PostVisitPrivate {
			// 从非私密转为私密才需要删除tag
			d.deleteTag(db, tag)
		}
	}
	db.Commit()
	d.indexActive(core.IdxActVisiblePost)
	return nil
}

func (d *dataServant) GetPostByID(id int64) (*model.Post, error) {
	post := &model.Post{
		Model: &model.Model{
			ID: id,
		},
	}
	return post.Get(d.engine)
}

func (d *dataServant) GetPosts(conditions *model.ConditionsT, offset, limit int) ([]*model.Post, error) {
	return (&model.Post{}).List(d.engine, conditions, offset, limit)
}

func (d *dataServant) GetPostCount(conditions *model.ConditionsT) (int64, error) {
	return (&model.Post{}).Count(d.engine, conditions)
}

func (d *dataServant) UpdatePost(post *model.Post) error {
	if err := post.Update(d.engine); err != nil {
		return err
	}
	d.indexActive(core.IdxActUpdatePost)
	return nil
}

func (d *dataServant) GetUserPostStar(postID, userID int64) (*model.PostStar, error) {
	star := &model.PostStar{
		PostID: postID,
		UserID: userID,
	}

	return star.Get(d.engine)
}

func (d *dataServant) GetUserPostStars(userID int64, offset, limit int) ([]*model.PostStar, error) {
	star := &model.PostStar{
		UserID: userID,
	}

	return star.List(d.engine, &model.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *dataServant) GetUserPostStarCount(userID int64) (int64, error) {
	star := &model.PostStar{
		UserID: userID,
	}
	return star.Count(d.engine, &model.ConditionsT{})
}

func (d *dataServant) CreatePostStar(postID, userID int64) (*model.PostStar, error) {
	star := &model.PostStar{
		PostID: postID,
		UserID: userID,
	}

	return star.Create(d.engine)
}

func (d *dataServant) DeletePostStar(p *model.PostStar) error {
	return p.Delete(d.engine)
}

func (d *dataServant) GetUserPostCollection(postID, userID int64) (*model.PostCollection, error) {
	star := &model.PostCollection{
		PostID: postID,
		UserID: userID,
	}

	return star.Get(d.engine)
}

func (d *dataServant) GetUserPostCollections(userID int64, offset, limit int) ([]*model.PostCollection, error) {
	collection := &model.PostCollection{
		UserID: userID,
	}

	return collection.List(d.engine, &model.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *dataServant) GetUserPostCollectionCount(userID int64) (int64, error) {
	collection := &model.PostCollection{
		UserID: userID,
	}
	return collection.Count(d.engine, &model.ConditionsT{})
}
func (d *dataServant) GetUserWalletBills(userID int64, offset, limit int) ([]*model.WalletStatement, error) {
	statement := &model.WalletStatement{
		UserID: userID,
	}

	return statement.List(d.engine, &model.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *dataServant) GetUserWalletBillCount(userID int64) (int64, error) {
	statement := &model.WalletStatement{
		UserID: userID,
	}
	return statement.Count(d.engine, &model.ConditionsT{})
}

func (d *dataServant) CreatePostCollection(postID, userID int64) (*model.PostCollection, error) {
	collection := &model.PostCollection{
		PostID: postID,
		UserID: userID,
	}

	return collection.Create(d.engine)
}

func (d *dataServant) DeletePostCollection(p *model.PostCollection) error {
	return p.Delete(d.engine)
}

func (d *dataServant) GetPostAttatchmentBill(postID, userID int64) (*model.PostAttachmentBill, error) {
	bill := &model.PostAttachmentBill{
		PostID: postID,
		UserID: userID,
	}

	return bill.Get(d.engine)
}
