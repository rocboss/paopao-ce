// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.TweetService       = (*tweetSrv)(nil)
	_ core.TweetManageService = (*tweetManageSrv)(nil)
	_ core.TweetHelpService   = (*tweetHelpSrv)(nil)

	_ core.TweetServantA       = (*tweetSrvA)(nil)
	_ core.TweetManageServantA = (*tweetManageSrvA)(nil)
	_ core.TweetHelpServantA   = (*tweetHelpSrvA)(nil)
)

type tweetSrv struct {
	*sqlxSrv
	q *cc.Tweet
}

type tweetManageSrv struct {
	*sqlxSrv
	cis core.CacheIndexService
	q   *cc.TweetManage
}

type tweetHelpSrv struct {
	*sqlxSrv
	q *cc.TweetHelp
}

type tweetSrvA struct {
	*sqlxSrv
	q *cc.TweetA
}

type tweetManageSrvA struct {
	*sqlxSrv
	cis core.CacheIndexService
	q   *cc.TweetManageA
}

type tweetHelpSrvA struct {
	*sqlxSrv
	q *cc.TweetHelpA
}

// MergePosts post数据整合
func (s *tweetHelpSrv) MergePosts(posts []*ms.Post) ([]*ms.PostFormated, error) {
	postIds := make([]int64, 0, len(posts))
	userIds := make([]int64, 0, len(posts))
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}
	postContents, err := s.getPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}
	users, err := s.getUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*dbr.UserFormated, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}
	contentMap := make(map[int64][]*dbr.PostContentFormated, len(postContents))
	for _, content := range postContents {
		contentMap[content.PostID] = append(contentMap[content.PostID], content)
	}
	// 数据整合
	postsFormated := make([]*dbr.PostFormated, 0, len(posts))
	for _, post := range posts {
		postFormated := post.Format()
		postFormated.User = userMap[post.UserID]
		postFormated.Contents = contentMap[post.ID]
		postsFormated = append(postsFormated, postFormated)
	}
	return postsFormated, nil
}

// RevampPosts post数据整形修复
func (s *tweetHelpSrv) RevampPosts(posts []*ms.PostFormated) ([]*ms.PostFormated, error) {
	postIds := make([]int64, 0, len(posts))
	userIds := make([]int64, 0, len(posts))
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}
	postContents, err := s.getPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}
	users, err := s.getUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*dbr.UserFormated, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}
	contentMap := make(map[int64][]*dbr.PostContentFormated, len(postContents))
	for _, content := range postContents {
		contentMap[content.PostID] = append(contentMap[content.PostID], content)
	}
	// 数据整合
	for _, post := range posts {
		post.User = userMap[post.UserID]
		post.Contents = contentMap[post.ID]
	}
	return posts, nil
}

func (s *tweetHelpSrv) getPostContentsByIDs(ids []int64) ([]*dbr.PostContentFormated, error) {
	res := []*dbr.PostContentFormated{}
	err := s.inSelect(s.db, &res, s.q.GetPostContentByIds, ids)
	return res, err
}

func (s *tweetHelpSrv) getUsersByIDs(ids []int64) ([]*dbr.UserFormated, error) {
	res := []*dbr.UserFormated{}
	err := s.inSelect(s.db, &res, s.q.GetUsersByIds, ids)
	return res, err
}

func (s *tweetManageSrv) CreatePostCollection(postID, userID int64) (*ms.PostCollection, error) {
	res, err := s.q.AddPostCollection.Exec(postID, userID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &ms.PostCollection{
		Model: &dbr.Model{
			ID: id,
		},
		PostID: postID,
		UserID: userID,
	}, nil
}

func (s *tweetManageSrv) DeletePostCollection(r *ms.PostCollection) error {
	_, err := s.q.DelPostCollection.Exec(r.ID)
	return err
}

func (s *tweetManageSrv) CreatePostContent(r *ms.PostContent) (*ms.PostContent, error) {
	res, err := s.q.AddPostContent.Exec(r.PostID, r.UserID, r.Content, r.Type, r.Sort)
	if err != nil {
		return nil, err
	}
	r.Model = &dbr.Model{}
	r.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *tweetManageSrv) CreateAttachment(r *ms.Attachment) (int64, error) {
	res, err := s.q.AddAttachment.Exec(r.UserID, r.FileSize, r.ImgWidth, r.ImgHeight, r.Type, r.Content)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *tweetManageSrv) CreatePost(r *ms.Post) (*ms.Post, error) {
	res, err := s.q.AddPost.Exec(r.UserID, r.Visibility, r.Tags, r.AttachmentPrice, r.IP, r.IPLoc)
	if err != nil {
		return nil, err
	}
	if id, err := res.LastInsertId(); err == nil {
		r.Model = &dbr.Model{
			ID: id,
		}
	} else {
		return nil, err
	}
	return r, nil
}

func (s *tweetManageSrv) DeletePost(post *ms.Post) ([]string, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetManageSrv) LockPost(r *ms.Post) error {
	_, err := s.q.LockPost.Exec(r.ID)
	return err
}

func (s *tweetManageSrv) StickPost(r *ms.Post) error {
	_, err := s.q.StickPost.Exec(r.ID)
	return err
}

func (s *tweetManageSrv) VisiblePost(r *ms.Post, visibility core.PostVisibleT) error {
	_, err := s.q.VisiblePost.Exec(r.ID, visibility)
	return err
}

func (s *tweetManageSrv) UpdatePost(r *ms.Post) error {
	if _, err := s.q.UpdatePost.Exec(r.CommentCount, r.UpvoteCount, r.ID); err != nil {
		return err
	}
	s.cis.SendAction(core.IdxActUpdatePost, r)
	return nil
}

func (s *tweetManageSrv) CreatePostStar(postID, userID int64) (*ms.PostStar, error) {
	res, err := s.q.AddPostStar.Exec(postID, userID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &ms.PostStar{
		Model: &dbr.Model{
			ID: id,
		},
		PostID: postID,
		UserID: userID,
	}, nil
}

func (s *tweetManageSrv) DeletePostStar(r *ms.PostStar) error {
	_, err := s.q.DelPostStar.Exec(r.ID)
	return err
}

func (s *tweetSrv) GetPostByID(id int64) (*ms.Post, error) {
	res := &ms.Post{}
	err := s.q.GetPostById.Get(res, id)
	return res, err
}

func (s *tweetSrv) GetPosts(conditions *ms.ConditionsT, offset, limit int) ([]*ms.Post, error) {
	// TODO 需要精细化接口
	debug.NotImplemented()
	return nil, nil
}

func (s *tweetSrv) GetPostCount(conditions *ms.ConditionsT) (int64, error) {
	// TODO 需要精细化接口
	debug.NotImplemented()
	return 0, nil
}

func (s *tweetSrv) GetUserPostStar(postID, userID int64) (*ms.PostStar, error) {
	res := &ms.PostStar{}
	err := s.q.GetUserPostStar.Get(res, postID, userID)
	return res, err
}

func (s *tweetSrv) GetUserPostStars(userID int64, offset, limit int) ([]*ms.PostStar, error) {
	res := []*ms.PostStar{}
	err := s.q.GetUserPostStar.Select(&res, userID, limit, offset)
	return res, err
}

func (s *tweetSrv) GetUserPostStarCount(userID int64) (res int64, err error) {
	err = s.q.GetUserPostStarCount.Get(&res, userID)
	return
}

func (s *tweetSrv) GetUserPostCollection(postID, userID int64) (*ms.PostCollection, error) {
	res := &ms.PostCollection{}
	err := s.q.GetUserPostCollection.Get(res, postID, userID)
	return res, err
}

func (s *tweetSrv) GetUserPostCollections(userID int64, offset, limit int) ([]*ms.PostCollection, error) {
	res := []*ms.PostCollection{}
	err := s.q.GetUserPostCollections.Select(&res, userID, limit, offset)
	return res, err
}

func (s *tweetSrv) GetUserPostCollectionCount(userID int64) (res int64, err error) {
	err = s.q.GetUserPostCollectionCount.Get(&res)
	return
}

func (s *tweetSrv) GetPostAttatchmentBill(postID, userID int64) (*ms.PostAttachmentBill, error) {
	res := &ms.PostAttachmentBill{}
	err := s.q.GetPostAttachmentBill.Get(res, postID, userID)
	return res, err
}

func (s *tweetSrv) GetPostContentsByIDs(ids []int64) ([]*ms.PostContent, error) {
	res := []*ms.PostContent{}
	err := s.inSelect(s.db, &res, s.q.GetPostContetnsByIds, ids)
	return res, err
}

func (s *tweetSrv) GetPostContentByID(id int64) (*ms.PostContent, error) {
	res := &ms.PostContent{}
	err := s.q.GetPostContentById.Get(res, id)
	return res, err
}

func (s *tweetSrvA) TweetInfoById(id int64) (*cs.TweetInfo, error) {
	res := &cs.TweetInfo{}
	err := s.q.TweetInfoById.Get(res, id)
	return res, err
}

func (s *tweetSrvA) TweetItemById(id int64) (*cs.TweetItem, error) {
	res := &cs.TweetItem{}
	err := s.q.TweetItemById.Get(res, id)
	if err != nil {
		return nil, err
	}
	// TODO need add contents info to res
	return res, nil
}

func (s *tweetSrvA) UserTweets(visitorId, userId int64) (res cs.TweetList, err error) {
	res = cs.TweetList{}
	switch s.checkRelationBy(visitorId, userId) {
	case 0: // admin
		err = s.q.UserTweetsByAdmin.Select(&res, userId)
	case 1: // self
		err = s.q.UserTweetsBySelf.Select(&res, userId)
	case 2: // friend
		err = s.q.UserTweetsByFriend.Select(&res, userId)
	case 3: // follower
		fallthrough
	default: // guest
		err = s.q.UserTweetsByGuest.Select(&res, userId)
	}
	if err != nil {
		return nil, err
	}
	// TODO need add contents info to res
	return res, nil
}

// checkRelationBy check the relation of visitor with user
func (s *tweetSrvA) checkRelationBy(visitorId, userId int64) uint {
	// TODO
	return 0
}

func (s *tweetSrvA) ReactionByTweetId(userId int64, tweetId int64) (*cs.ReactionItem, error) {
	res := &cs.ReactionItem{}
	err := s.q.ReactionByTweetId.Get(res, userId, tweetId)
	return res, err
}

func (s *tweetSrvA) UserReactions(userId int64, limit int, offset int) (cs.ReactionList, error) {
	res := cs.ReactionList{}
	err := s.q.UserReactions.Select(&res, userId)
	return res, err
}

func (s *tweetSrvA) FavoriteByTweetId(userId int64, tweetId int64) (*cs.FavoriteItem, error) {
	res := &cs.FavoriteItem{}
	err := s.q.FavoriteByTweetId.Get(res, userId, tweetId)
	return res, err
}

func (s *tweetSrvA) UserFavorites(userId int64, limit int, offset int) (cs.FavoriteList, error) {
	res := cs.FavoriteList{}
	err := s.q.UserFavorites.Select(&res, userId)
	return res, err
}

func (s *tweetSrvA) AttachmentByTweetId(userId int64, tweetId int64) (*cs.AttachmentBill, error) {
	res := &cs.AttachmentBill{}
	err := s.q.AttachmentByTweetId.Get(res, userId, tweetId)
	return res, err
}

func (s *tweetManageSrvA) CreateAttachment(obj *cs.Attachment) (int64, error) {
	// TODO
	return 0, debug.ErrNotImplemented
}

func (s *tweetManageSrvA) CreateTweet(userId int64, req *cs.NewTweetReq) (*cs.TweetItem, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageSrvA) DeleteTweet(userId int64, tweetId int64) ([]string, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetManageSrvA) LockTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) StickTweet(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) VisibleTweet(userId int64, visibility cs.TweetVisibleType) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) CreateReaction(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) DeleteReaction(userId int64, reactionId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) CreateFavorite(userId int64, tweetId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetManageSrvA) DeleteFavorite(userId int64, favoriteId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *tweetHelpSrvA) RevampTweets(tweets cs.TweetList) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func (s *tweetHelpSrvA) MergeTweets(tweets cs.TweetInfo) (cs.TweetList, error) {
	// TODO
	return nil, debug.ErrNotImplemented
}

func newTweetService(db *sqlx.DB) core.TweetService {
	return &tweetSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweet),
	}
}

func newTweetManageService(db *sqlx.DB, cacheIndex core.CacheIndexService) core.TweetManageService {
	return &tweetManageSrv{
		sqlxSrv: newSqlxSrv(db),
		cis:     cacheIndex,
		q:       mustBuild(db, cc.BuildTweetManage),
	}
}

func newTweetHelpService(db *sqlx.DB) core.TweetHelpService {
	return &tweetHelpSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweetHelp),
	}
}

func newTweetServantA(db *sqlx.DB) core.TweetServantA {
	return &tweetSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweetA),
	}
}

func newTweetManageServantA(db *sqlx.DB, cacheIndex core.CacheIndexService) core.TweetManageServantA {
	return &tweetManageSrvA{
		sqlxSrv: newSqlxSrv(db),
		cis:     cacheIndex,
		q:       mustBuild(db, cc.BuildTweetManageA),
	}
}

func newTweetHelpServantA(db *sqlx.DB) core.TweetHelpServantA {
	return &tweetHelpSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildTweetHelpA),
	}
}
