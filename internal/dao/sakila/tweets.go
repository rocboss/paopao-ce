// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

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

func (s *tweetHelpSrv) getPostContentsByIDs(ids []int64) (res []*dbr.PostContentFormated, err error) {
	err = s.inSelect(&res, s.q.GetPostContentByIds, ids)
	return
}

func (s *tweetHelpSrv) getUsersByIDs(ids []int64) (res []*dbr.UserFormated, err error) {
	err = s.inSelect(&res, s.q.GetUsersByIds, ids)
	return
}

func (s *tweetManageSrv) CreatePostCollection(postID, userID int64) (*ms.PostCollection, error) {
	now := time.Now().Unix()
	res, err := s.q.AddPostCollection.Exec(postID, userID, now)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &ms.PostCollection{
		Model: &dbr.Model{
			ID:        id,
			CreatedOn: now,
		},
		PostID: postID,
		UserID: userID,
	}, nil
}

func (s *tweetManageSrv) DeletePostCollection(r *ms.PostCollection) error {
	_, err := s.q.DelPostCollection.Exec(time.Now().Unix(), r.ID)
	return err
}

func (s *tweetManageSrv) CreatePostContent(r *ms.PostContent) (*ms.PostContent, error) {
	r.Model = &ms.Model{CreatedOn: time.Now().Unix()}
	res, err := s.q.AddPostContent.Exec(r)
	if err != nil {
		return nil, err
	}
	r.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *tweetManageSrv) CreateAttachment(r *ms.Attachment) (int64, error) {
	args := []any{r.UserID, r.FileSize, r.ImgWidth, r.ImgHeight, r.Type, r.Content, time.Now().Unix()}
	res, err := s.q.AddAttachment.Exec(args...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *tweetManageSrv) CreatePost(r *ms.Post) (*ms.Post, error) {
	now := time.Now().Unix()
	r.Model = &dbr.Model{CreatedOn: now}
	r.LatestRepliedOn = now
	res, err := s.q.AddPost.Exec(r)
	if err != nil {
		return nil, err
	}
	if r.ID, err = res.LastInsertId(); err != nil {
		return nil, err
	}
	s.cis.SendAction(core.IdxActCreatePost, r)
	return r, nil
}

func (s *tweetManageSrv) DeletePost(post *ms.Post) (mediaContents []string, err error) {
	s.with(func(tx *sqlx.Tx) error {
		// 获取多媒体内容
		if err = tx.Stmtx(s.q.MediaContentByPostId).Get(&mediaContents, post.ID); err != nil {
			return err
		}
		// 删推文
		now := time.Now().Unix()
		if _, err = tx.Stmtx(s.q.DelPostById).Exec(now, post.ID); err != nil {
			return err
		}
		// 删评论
		contents, err := s.deleteCommentByPostId(tx, post.ID)
		if err != nil {
			return err
		}
		mediaContents = append(mediaContents, contents...)
		if tags := strings.Split(post.Tags, ","); len(tags) > 0 {
			// 删tag，宽松处理错误，有错误不会回滚
			s.deleteTags(tx, tags)
		}
		return nil
	})
	s.cis.SendAction(core.IdxActDeletePost, post)
	return
}

func (s *tweetManageSrv) deleteCommentByPostId(tx *sqlx.Tx, postId int64) (mediaContents []string, err error) {
	var commentIds []int64
	// 获取推文的所有评论id
	err = tx.Stmtx(s.q.CommentIdsByPostId).Select(&commentIds, postId)
	if err != nil {
		return nil, err
	}
	// 获取评论的媒体内容
	err = s.inSelectx(tx, &mediaContents, s.q.CommentMediaFromCommentIds, commentIds)
	if err != nil {
		return nil, err
	}
	// 删评论
	now := time.Now().Unix()
	if _, err = tx.Stmtx(s.q.DelCommentByPostId).Exec(now, postId); err != nil {
		return nil, err
	}
	// 删评论内容
	if _, err = s.inExecx(tx, s.q.DelCommentContentByCommentIds, now, commentIds); err != nil {
		return nil, err
	}
	// 删评论的评论
	if _, err = s.inExecx(tx, s.q.DelReplyByCommentIds, now, commentIds); err != nil {
		return nil, err
	}
	return
}

func (s *tweetManageSrv) LockPost(r *ms.Post) error {
	_, err := s.q.LockPost.Exec(time.Now().Unix(), r.ID)
	return err
}

func (s *tweetManageSrv) StickPost(r *ms.Post) error {
	_, err := s.q.StickPost.Exec(time.Now().Unix(), r.ID)
	return err
}

func (s *tweetManageSrv) VisiblePost(post *ms.Post, visibility core.PostVisibleT) (err error) {
	oldVisibility := post.Visibility
	post.Visibility = visibility
	// TODO: 这个判断是否可以不要呢
	if oldVisibility == visibility {
		return nil
	}
	// 私密推文 特殊处理
	if visibility == ms.PostVisitPrivate {
		// 强制取消置顶
		// TODO: 置顶推文用户是否有权设置成私密？ 后续完善
		post.IsTop = 0
	}
	s.with(func(tx *sqlx.Tx) error {
		_, err = s.q.VisiblePost.Exec(visibility, post.IsTop, time.Now().Unix(), post.ID)
		// tag处理
		tags := strings.Split(post.Tags, ",")
		// TODO: 暂时宽松不处理错误，这里或许可以有优化，后续完善
		if oldVisibility == dbr.PostVisitPrivate {
			// 从私密转为非私密才需要重新创建tag
			s.createTags(tx, post.UserID, tags)
		} else if visibility == dbr.PostVisitPrivate {
			// 从非私密转为私密才需要删除tag
			s.deleteTags(tx, tags)
		}
		return nil
	})
	s.cis.SendAction(core.IdxActVisiblePost, post)
	return
}

func (s *tweetManageSrv) UpdatePost(r *ms.Post) error {
	r.ModifiedOn = time.Now().Unix()
	if _, err := s.q.UpdatePost.Exec(r); err != nil {
		return err
	}
	s.cis.SendAction(core.IdxActUpdatePost, r)
	return nil
}

func (s *tweetManageSrv) CreatePostStar(postID, userID int64) (*ms.PostStar, error) {
	now := time.Now().Unix()
	res, err := s.q.AddPostStar.Exec(postID, userID, now)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &ms.PostStar{
		Model: &dbr.Model{
			ID:        id,
			CreatedOn: now,
		},
		PostID: postID,
		UserID: userID,
	}, nil
}

func (s *tweetManageSrv) DeletePostStar(r *ms.PostStar) error {
	_, err := s.q.DelPostStar.Exec(time.Now().Unix(), r.ID)
	return err
}

func (s *tweetSrv) GetPostByID(id int64) (res *ms.Post, err error) {
	err = s.q.GetPostById.Get(res, id)
	return
}

func (s *tweetSrv) GetPosts(c ms.ConditionsT, offset, limit int) (res []*ms.Post, err error) {
	userId, exist := c["user_id"]
	visibility := c["visibility IN ?"]
	if exist {
		err = s.inSelect(&res, s.q.GetUserPosts, userId, visibility, limit, offset)
	} else {
		err = s.inSelect(&res, s.q.GetAnyPosts, visibility, limit, offset)
	}
	return
}

func (s *tweetSrv) GetPostCount(c ms.ConditionsT) (res int64, err error) {
	userId, exist := c["user_id"]
	visibility := c["visibility IN ?"]
	if exist {
		err = s.inGet(&res, s.q.GetUserPostCount, userId, visibility)
	} else {
		err = s.inGet(&res, s.q.GetAnyPostCount, visibility)
	}
	return
}

func (s *tweetSrv) GetUserPostStar(postID, userID int64) (res *ms.PostStar, err error) {
	err = s.q.GetUserPostStar.Get(res, postID, userID, userID)
	return
}

func (s *tweetSrv) GetUserPostStars(userID int64, offset, limit int) (res []*ms.PostStar, err error) {
	err = s.q.GetUserPostStar.Select(&res, userID, userID, limit, offset)
	return
}

func (s *tweetSrv) GetUserPostStarCount(userID int64) (res int64, err error) {
	err = s.q.GetUserPostStarCount.Get(&res, userID, userID)
	return
}

func (s *tweetSrv) GetUserPostCollection(postID, userID int64) (res *ms.PostCollection, err error) {
	err = s.q.GetUserPostCollection.Get(res, postID, userID, userID)
	return
}

func (s *tweetSrv) GetUserPostCollections(userID int64, offset, limit int) (res []*ms.PostCollection, err error) {
	err = s.q.GetUserPostCollections.Select(&res, userID, userID, limit, offset)
	return res, err
}

func (s *tweetSrv) GetUserPostCollectionCount(userID int64) (res int64, err error) {
	err = s.q.GetUserPostCollectionCount.Get(&res, userID, userID)
	return
}

func (s *tweetSrv) GetPostAttatchmentBill(postID, userID int64) (res *ms.PostAttachmentBill, err error) {
	err = s.q.GetPostAttachmentBill.Get(res, postID, userID)
	return
}

func (s *tweetSrv) GetPostContentsByIDs(ids []int64) (res []*ms.PostContent, err error) {
	err = s.inSelect(&res, s.q.GetPostContentsByIds, ids)
	return
}

func (s *tweetSrv) GetPostContentByID(id int64) (res *ms.PostContent, err error) {
	err = s.q.GetPostContentById.Get(res, id)
	return
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
		q:       mustBuildFn(db, cc.BuildTweetHelp),
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
