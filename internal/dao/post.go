package dao

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/model"
)

func (d *Dao) CreatePost(post *model.Post) (*model.Post, error) {
	post.LatestRepliedOn = time.Now().Unix()
	return post.Create(d.engine)
}

func (d *Dao) DeletePost(post *model.Post) error {
	return post.Delete(d.engine)
}

func (d *Dao) LockPost(post *model.Post) error {
	post.IsLock = 1 - post.IsLock
	return post.Update(d.engine)
}

func (d *Dao) GetPostByID(id int64) (*model.Post, error) {
	post := &model.Post{
		Model: &model.Model{
			ID: id,
		},
	}
	return post.Get(d.engine)
}

func (d *Dao) GetPosts(conditions *model.ConditionsT, offset, limit int) ([]*model.Post, error) {
	return (&model.Post{}).List(d.engine, conditions, offset, limit)
}

func (d *Dao) GetPostCount(conditions *model.ConditionsT) (int64, error) {
	return (&model.Post{}).Count(d.engine, conditions)
}

func (d *Dao) UpdatePost(post *model.Post) error {
	return post.Update(d.engine)
}

func (d *Dao) GetUserPostStar(postID, userID int64) (*model.PostStar, error) {
	star := &model.PostStar{
		PostID: postID,
		UserID: userID,
	}

	return star.Get(d.engine)
}

func (d *Dao) GetUserPostStars(userID int64, offset, limit int) ([]*model.PostStar, error) {
	star := &model.PostStar{
		UserID: userID,
	}

	return star.List(d.engine, &model.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *Dao) GetUserPostStarCount(userID int64) (int64, error) {
	star := &model.PostStar{
		UserID: userID,
	}
	return star.Count(d.engine, &model.ConditionsT{})
}

func (d *Dao) CreatePostStar(postID, userID int64) (*model.PostStar, error) {
	star := &model.PostStar{
		PostID: postID,
		UserID: userID,
	}

	return star.Create(d.engine)
}

func (d *Dao) DeletePostStar(p *model.PostStar) error {
	return p.Delete(d.engine)
}

func (d *Dao) GetUserPostCollection(postID, userID int64) (*model.PostCollection, error) {
	star := &model.PostCollection{
		PostID: postID,
		UserID: userID,
	}

	return star.Get(d.engine)
}

func (d *Dao) GetUserPostCollections(userID int64, offset, limit int) ([]*model.PostCollection, error) {
	collection := &model.PostCollection{
		UserID: userID,
	}

	return collection.List(d.engine, &model.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *Dao) GetUserPostCollectionCount(userID int64) (int64, error) {
	collection := &model.PostCollection{
		UserID: userID,
	}
	return collection.Count(d.engine, &model.ConditionsT{})
}
func (d *Dao) GetUserWalletBills(userID int64, offset, limit int) ([]*model.WalletStatement, error) {
	statement := &model.WalletStatement{
		UserID: userID,
	}

	return statement.List(d.engine, &model.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *Dao) GetUserWalletBillCount(userID int64) (int64, error) {
	statement := &model.WalletStatement{
		UserID: userID,
	}
	return statement.Count(d.engine, &model.ConditionsT{})
}

func (d *Dao) CreatePostCollection(postID, userID int64) (*model.PostCollection, error) {
	collection := &model.PostCollection{
		PostID: postID,
		UserID: userID,
	}

	return collection.Create(d.engine)
}

func (d *Dao) DeletePostCollection(p *model.PostCollection) error {
	return p.Delete(d.engine)
}

func (d *Dao) GetPostAttatchmentBill(postID, userID int64) (*model.PostAttachmentBill, error) {
	bill := &model.PostAttachmentBill{
		PostID: postID,
		UserID: userID,
	}

	return bill.Get(d.engine)
}
