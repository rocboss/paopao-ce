// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"strings"
	"time"

	"github.com/alimy/cfg"
	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pg"
)

var (
	_ core.TweetService       = (*tweetSrv)(nil)
	_ core.TweetManageService = (*tweetManageSrv)(nil)
	_ core.TweetHelpService   = (*tweetHelpSrv)(nil)
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

// MergePosts post数据整合
func (s *tweetHelpSrv) MergePosts(posts []*ms.Post) ([]*ms.PostFormated, error) {
	if len(posts) == 0 {
		return nil, nil
	}
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
	if len(posts) == 0 {
		return nil, nil
	}
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
	err = s.db.InSelect(&res, s.q.GetPostContentByIds, ids)
	return
}

func (s *tweetHelpSrv) getUsersByIDs(ids []int64) (res []*dbr.UserFormated, err error) {
	err = s.db.InSelect(&res, s.q.GetUsersByIds, ids)
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
	s.db.Withx(func(tx *sqlx.Tx) error {
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
	err = tx.InSelect(&mediaContents, s.q.CommentMediaFromCommentIds, commentIds)
	if err != nil {
		return nil, err
	}
	// 删评论
	now := time.Now().Unix()
	if _, err = tx.Stmtx(s.q.DelCommentByPostId).Exec(now, postId); err != nil {
		return nil, err
	}
	// 删评论内容
	if _, err = tx.InExec(s.q.DelCommentContentByCommentIds, now, commentIds); err != nil {
		return nil, err
	}
	// 删评论的评论
	if _, err = tx.InExec(s.q.DelReplyByCommentIds, now, commentIds); err != nil {
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

func (s *tweetManageSrv) HighlightPost(userId, postId int64) (is_essence int, err error) {
	err = s.db.Withx(func(tx *sqlx.Tx) error {
		var postUserId int64
		err = tx.Stmtx(s.q.PostHighlightStatus).QueryRowx(postId).Scan(&postUserId, &is_essence)
		if err != nil {
			return err
		}
		if postUserId != userId {
			return cs.ErrNoPermission
		}
		if _, err = tx.Stmtx(s.q.HighlightPost).Exec(postId); err != nil {
			return err
		}
		is_essence = 1 - is_essence
		return nil
	})
	return
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
	s.db.Withx(func(tx *sqlx.Tx) error {
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
	res = &ms.Post{}
	err = s.q.GetPostById.Get(res, id)
	return
}

func (s *tweetSrv) GetPosts(c ms.ConditionsT, offset, limit int) (res []*ms.Post, err error) {
	userId, exist := c["user_id"]
	visibility := c["visibility IN ?"]
	if exist {
		err = s.db.InSelect(&res, s.q.GetUserPosts, userId, visibility, limit, offset)
	} else {
		err = s.db.InSelect(&res, s.q.GetAnyPosts, visibility, limit, offset)
	}
	return
}

func (s *tweetSrv) GetPostCount(c ms.ConditionsT) (res int64, err error) {
	userId, exist := c["user_id"]
	visibility := c["visibility IN ?"]
	if exist {
		err = s.db.InGet(&res, s.q.GetUserPostCount, userId, visibility)
	} else {
		err = s.db.InGet(&res, s.q.GetAnyPostCount, visibility)
	}
	return
}

func (s *tweetSrv) GetUserPostStar(postID, userID int64) (res *ms.PostStar, err error) {
	res = &ms.PostStar{}
	err = s.q.GetUserPostStar.Get(res, postID, userID, userID)
	return
}

func (s *tweetSrv) GetUserPostStars(userID int64, limit int, offset int) (res []*ms.PostStar, err error) {
	err = s.q.GetUserPostStars.Select(&res, userID, userID, limit, offset)
	return
}

func (s *tweetSrv) ListUserStarTweets(user *cs.VistUser, limit int, offset int) (res []*ms.PostStar, total int64, err error) {
	switch user.RelTyp {
	case cs.RelationAdmin:
		if err = s.q.UserStarTweetsByAdmin.Select(&res, user.UserId, limit, offset); err == nil {
			err = s.q.UserStarTweetsCountByAdmin.Get(&total, user.UserId)
		}
	case cs.RelationSelf:
		// Fixed: 这里的查询有Bug，没有考虑查询Friend star 时如果对方已经不是你好友时，应该去除这个star
		if err = s.q.UserStarTweetsBySelf.Select(&res, user.UserId, user.UserId, limit, offset); err == nil {
			err = s.q.UserStarTweetsCountBySelf.Get(&total, user.UserId, user.UserId)
		}
	case cs.RelationFriend:
		if err = s.q.UserStarTweetsByFriend.Select(&res, user.UserId, limit, offset); err == nil {
			err = s.q.UserStarTweetsByFriend.Get(&total, user.UserId)
		}
	case cs.RelationGuest:
		fallthrough
	default:
		if err = s.q.UserStarTweetsByGuest.Select(&res, user.UserId, limit, offset); err == nil {
			err = s.q.UserStarTweetsCountByGuest.Get(&total, user.UserId)
		}
	}
	return
}

func (s *tweetSrv) ListUserMediaTweets(user *cs.VistUser, limit int, offset int) (res []*ms.Post, total int64, err error) {
	gStmt, cStmt := s.q.UserMediaTweetsByGuest, s.q.UserMediaTweetsCountByGuest
	switch user.RelTyp {
	case cs.RelationAdmin, cs.RelationSelf:
		gStmt, cStmt = s.q.UserMediaTweetsBySelf, s.q.UserMediaTweetsCountBySelf
	case cs.RelationFriend:
		gStmt, cStmt = s.q.UserMediaTweetsByFriend, s.q.UserMediaTweetsCountByFriend
	case cs.RelationGuest:
		fallthrough
	default:
		// nothing
	}
	if err = gStmt.Select(&res, user.UserId, limit, offset); err == nil {
		err = cStmt.Get(&total, user.UserId)
	}
	return
}

func (s *tweetSrv) ListUserCommentTweets(user *cs.VistUser, limit int, offset int) (res []*ms.Post, total int64, err error) {
	gStmt, cStmt := s.q.UserCommentTweetsByGuest, s.q.UserCommentTweetsCountByGuest
	switch user.RelTyp {
	case cs.RelationAdmin, cs.RelationSelf:
		gStmt, cStmt = s.q.UserCommentTweetsBySelf, s.q.UserCommentTweetsCountBySelf
	case cs.RelationFriend:
		gStmt, cStmt = s.q.UserCommentTweetsByFriend, s.q.UserCommentTweetsCountByFriend
	case cs.RelationGuest:
		fallthrough
	default:
		// nothing
	}
	if err = gStmt.Select(&res, user.UserId, limit, offset); err == nil {
		err = cStmt.Get(&total, user.UserId)
	}
	return
}

func (s *tweetSrv) GetUserPostStarCount(userID int64) (res int64, err error) {
	err = s.q.GetUserPostStarCount.Get(&res, userID, userID)
	return
}

func (s *tweetSrv) GetUserPostCollection(postID, userID int64) (res *ms.PostCollection, err error) {
	res = &ms.PostCollection{}
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
	res = &ms.PostAttachmentBill{}
	err = s.q.GetPostAttachmentBill.Get(res, postID, userID)
	return
}

func (s *tweetSrv) GetPostContentsByIDs(ids []int64) (res []*ms.PostContent, err error) {
	err = s.db.InSelect(&res, s.q.GetPostContentsByIds, ids)
	return
}

func (s *tweetSrv) GetPostContentByID(id int64) (res *ms.PostContent, err error) {
	res = &ms.PostContent{}
	err = s.q.GetPostContentById.Get(res, id)
	return
}

func newTweetService(db *sqlx.DB) core.TweetService {
	return &tweetSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildTweet),
	}
}

func newTweetManageService(db *sqlx.DB, cacheIndex core.CacheIndexService) (s core.TweetManageService) {
	tms := &tweetManageSrv{
		sqlxSrv: newSqlxSrv(db),
		cis:     cacheIndex,
		q:       ccBuild(db, cc.BuildTweetManage),
	}
	s = tms
	if cfg.Any("PostgreSQL", "PgSQL", "Postgres") {
		s = &pgTweetManageSrv{
			tweetManageSrv: tms,
			p:              pgBuild(db, pg.BuildTweetManage),
		}
	}
	return
}

func newTweetHelpService(db *sqlx.DB) core.TweetHelpService {
	return &tweetHelpSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuildFn(db, cc.BuildTweetHelp),
	}
}
