// Code generated by Yesql. DO NOT EDIT.
// versions:
// - Yesql v1.9.0

package pgc

import (
	"context"
	"fmt"

	"github.com/alimy/yesql"
	"github.com/bitbus/sqlx"
)

var (
	_ = fmt.Errorf("error for placeholder")
)

const (
	_CommentManage_CreateComment        = `INSERT INTO @comment (post_id, user_id, ip, ip_loc, created_on) VALUES (?, ?, ?, ?, ?) RETURNING *`
	_CommentManage_CreateCommentContent = `INSERT INTO @comment_content (comment_id, user_id, content, type, sort, created_on) VALUES (?, ?, ?, ?, ?, ?) RETURNING *`
	_CommentManage_CreateCommentReply   = `INSERT INTO @comment_reply (comment_id, user_id, content, at_user_id, ip, ip_loc, created_on) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING *`
	_ContactManager_CreateContact       = `INSERT INTO @contact (user_id, friend_id, status, created_on) VALUES (?, ?, ?, ?) RETURNING *`
	_Message_CreateMessage              = `INSERT INTO @message (sender_user_id, receiver_user_id, type, brief, content, post_id, comment_id, reply_id, created_on) VALUES (:sender_user_id, :receiver_user_id, :type, :brief, :content, :post_id, :comment_id, :reply_id, :created_on) RETURNING id`
	_Tweet_ListIndexHotsTweets          = `SELECT post.* FROM @post post LEFT JOIN @post_metric metric ON post.id=metric.post_id WHERE post.visibility>=90 AND post.is_del=0 ORDER BY post.is_top DESC, metric.rank_score DESC NULLS LAST, post.latest_replied_on DESC LIMIT ? OFFSET ?`
	_TweetManage_AddAttachment          = `INSERT INTO @attachment (user_id, file_size, img_width, img_height, type, content, created_on) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id`
	_TweetManage_AddPost                = `INSERT INTO @post (user_id, tags, ip, ip_loc, attachment_price, visibility, latest_replied_on, created_on) VALUES (:user_id, :tags, :ip, :ip_loc, :attachment_price, :visibility, :latest_replied_on, :created_on) RETURNING id`
	_TweetManage_AddPostCollection      = `INSERT INTO @post_collection (post_id, user_id, created_on) VALUES (?, ?, ?) RETURNING *`
	_TweetManage_AddPostContent         = `INSERT INTO @post_content (post_id, user_id, content, type, sort, created_on) VALUES (:post_id, :user_id, :content, :type, :sort, :created_on) RETURNING id`
	_TweetManage_AddPostStar            = `INSERT INTO @post_star (post_id, user_id, created_on) VALUES (?, ?, ?) RETURNING *`
	_UserManage_CreateUser              = `INSERT INTO @user (username, nickname, password, salt, avatar, status, created_on, balance) VALUES (:username, :nickname, :password, :salt, :avatar, :status, :created_on, 0) RETURNING id`
	_Wallet_CreateRecharge              = `INSERT INTO @wallet_recharge (user_id, amount, created_on) VALUES (?, ?, ?) RETURNING *`
)

// PreparexContext enhances the Conn interface with context.
type PreparexContext interface {
	// PrepareContext prepares a statement.
	// The provided context is used for the preparation of the statement, not for
	// the execution of the statement.
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)

	// PrepareNamedContext returns an sqlx.NamedStmt
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	// Rebind rebind query to adapte SQL Driver
	Rebind(query string) string
}

// PreparexBuilder preparex builder interface for sqlx
type PreparexBuilder interface {
	PreparexContext
	QueryHook(query string) string
}

type CommentManage struct {
	yesql.Namespace      `yesql:"comment_manage"`
	CreateComment        *sqlx.Stmt `yesql:"create_comment"`
	CreateCommentContent *sqlx.Stmt `yesql:"create_comment_content"`
	CreateCommentReply   *sqlx.Stmt `yesql:"create_comment_reply"`
}

type ContactManager struct {
	yesql.Namespace `yesql:"contact_manager"`
	CreateContact   *sqlx.Stmt `yesql:"create_contact"`
}

type Message struct {
	yesql.Namespace `yesql:"message"`
	CreateMessage   *sqlx.NamedStmt `yesql:"create_message"`
}

type Tweet struct {
	yesql.Namespace     `yesql:"tweet"`
	ListIndexHotsTweets *sqlx.Stmt `yesql:"list_index_hots_tweets"`
}

type TweetManage struct {
	yesql.Namespace   `yesql:"tweet_manage"`
	AddAttachment     *sqlx.Stmt      `yesql:"add_attachment"`
	AddPostCollection *sqlx.Stmt      `yesql:"add_post_collection"`
	AddPostStar       *sqlx.Stmt      `yesql:"add_post_star"`
	AddPost           *sqlx.NamedStmt `yesql:"add_post"`
	AddPostContent    *sqlx.NamedStmt `yesql:"add_post_content"`
}

type UserManage struct {
	yesql.Namespace `yesql:"user_manage"`
	CreateUser      *sqlx.NamedStmt `yesql:"create_user"`
}

type Wallet struct {
	yesql.Namespace `yesql:"wallet"`
	CreateRecharge  *sqlx.Stmt `yesql:"create_recharge"`
}

func BuildCommentManage(p PreparexBuilder, ctx ...context.Context) (obj *CommentManage, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &CommentManage{}
	if obj.CreateComment, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_CommentManage_CreateComment))); err != nil {
		return nil, fmt.Errorf("prepare _CommentManage_CreateComment error: %w", err)
	}
	if obj.CreateCommentContent, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_CommentManage_CreateCommentContent))); err != nil {
		return nil, fmt.Errorf("prepare _CommentManage_CreateCommentContent error: %w", err)
	}
	if obj.CreateCommentReply, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_CommentManage_CreateCommentReply))); err != nil {
		return nil, fmt.Errorf("prepare _CommentManage_CreateCommentReply error: %w", err)
	}
	return
}

func BuildContactManager(p PreparexBuilder, ctx ...context.Context) (obj *ContactManager, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &ContactManager{}
	if obj.CreateContact, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_ContactManager_CreateContact))); err != nil {
		return nil, fmt.Errorf("prepare _ContactManager_CreateContact error: %w", err)
	}
	return
}

func BuildMessage(p PreparexBuilder, ctx ...context.Context) (obj *Message, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Message{}
	if obj.CreateMessage, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_Message_CreateMessage))); err != nil {
		return nil, fmt.Errorf("prepare _Message_CreateMessage error: %w", err)
	}
	return
}

func BuildTweet(p PreparexBuilder, ctx ...context.Context) (obj *Tweet, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Tweet{}
	if obj.ListIndexHotsTweets, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Tweet_ListIndexHotsTweets))); err != nil {
		return nil, fmt.Errorf("prepare _Tweet_ListIndexHotsTweets error: %w", err)
	}
	return
}

func BuildTweetManage(p PreparexBuilder, ctx ...context.Context) (obj *TweetManage, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &TweetManage{}
	if obj.AddAttachment, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TweetManage_AddAttachment))); err != nil {
		return nil, fmt.Errorf("prepare _TweetManage_AddAttachment error: %w", err)
	}
	if obj.AddPostCollection, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TweetManage_AddPostCollection))); err != nil {
		return nil, fmt.Errorf("prepare _TweetManage_AddPostCollection error: %w", err)
	}
	if obj.AddPostStar, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TweetManage_AddPostStar))); err != nil {
		return nil, fmt.Errorf("prepare _TweetManage_AddPostStar error: %w", err)
	}
	if obj.AddPost, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_TweetManage_AddPost))); err != nil {
		return nil, fmt.Errorf("prepare _TweetManage_AddPost error: %w", err)
	}
	if obj.AddPostContent, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_TweetManage_AddPostContent))); err != nil {
		return nil, fmt.Errorf("prepare _TweetManage_AddPostContent error: %w", err)
	}
	return
}

func BuildUserManage(p PreparexBuilder, ctx ...context.Context) (obj *UserManage, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &UserManage{}
	if obj.CreateUser, err = p.PrepareNamedContext(c, p.Rebind(p.QueryHook(_UserManage_CreateUser))); err != nil {
		return nil, fmt.Errorf("prepare _UserManage_CreateUser error: %w", err)
	}
	return
}

func BuildWallet(p PreparexBuilder, ctx ...context.Context) (obj *Wallet, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &Wallet{}
	if obj.CreateRecharge, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_Wallet_CreateRecharge))); err != nil {
		return nil, fmt.Errorf("prepare _Wallet_CreateRecharge error: %w", err)
	}
	return
}
