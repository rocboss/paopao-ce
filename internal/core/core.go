package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// DataService data service interface that process data related logic on database
type DataService interface {
	WalletService
	SearchService
	StorageService

	GetComments(conditions *model.ConditionsT, offset, limit int) ([]*model.Comment, error)
	GetCommentByID(id int64) (*model.Comment, error)
	DeleteComment(comment *model.Comment) error
	GetCommentCount(conditions *model.ConditionsT) (int64, error)
	CreateComment(comment *model.Comment) (*model.Comment, error)
	CreateCommentReply(reply *model.CommentReply) (*model.CommentReply, error)
	GetCommentReplyByID(id int64) (*model.CommentReply, error)
	DeleteCommentReply(reply *model.CommentReply) error
	GetCommentContentsByIDs(ids []int64) ([]*model.CommentContent, error)
	GetCommentRepliesByID(ids []int64) ([]*model.CommentReplyFormated, error)
	CreateCommentContent(content *model.CommentContent) (*model.CommentContent, error)

	CreateMessage(msg *model.Message) (*model.Message, error)
	GetUnreadCount(userID int64) (int64, error)
	GetMessageByID(id int64) (*model.Message, error)
	ReadMessage(message *model.Message) error
	GetMessages(conditions *model.ConditionsT, offset, limit int) ([]*model.MessageFormated, error)
	GetMessageCount(conditions *model.ConditionsT) (int64, error)

	CreatePost(post *model.Post) (*model.Post, error)
	DeletePost(post *model.Post) error
	LockPost(post *model.Post) error
	StickPost(post *model.Post) error
	GetPostByID(id int64) (*model.Post, error)
	GetPosts(conditions *model.ConditionsT, offset, limit int) ([]*model.Post, error)
	GetPostCount(conditions *model.ConditionsT) (int64, error)
	UpdatePost(post *model.Post) error
	GetUserPostStar(postID, userID int64) (*model.PostStar, error)
	GetUserPostStars(userID int64, offset, limit int) ([]*model.PostStar, error)
	GetUserPostStarCount(userID int64) (int64, error)
	CreatePostStar(postID, userID int64) (*model.PostStar, error)
	DeletePostStar(p *model.PostStar) error
	GetUserPostCollection(postID, userID int64) (*model.PostCollection, error)
	GetUserPostCollections(userID int64, offset, limit int) ([]*model.PostCollection, error)
	GetUserPostCollectionCount(userID int64) (int64, error)
	GetUserWalletBills(userID int64, offset, limit int) ([]*model.WalletStatement, error)
	GetUserWalletBillCount(userID int64) (int64, error)
	CreatePostCollection(postID, userID int64) (*model.PostCollection, error)
	DeletePostCollection(p *model.PostCollection) error
	GetPostAttatchmentBill(postID, userID int64) (*model.PostAttachmentBill, error)
	CreatePostContent(content *model.PostContent) (*model.PostContent, error)
	GetPostContentsByIDs(ids []int64) ([]*model.PostContent, error)
	GetPostContentByID(id int64) (*model.PostContent, error)

	CreateTag(tag *model.Tag) (*model.Tag, error)
	DeleteTag(tag *model.Tag) error
	GetTags(conditions *model.ConditionsT, offset, limit int) ([]*model.Tag, error)

	GetUserByID(id int64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByPhone(phone string) (*model.User, error)
	GetUsersByIDs(ids []int64) ([]*model.User, error)
	GetUsersByKeyword(keyword string) ([]*model.User, error)
	GetTagsByKeyword(keyword string) ([]*model.Tag, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	GetLatestPhoneCaptcha(phone string) (*model.Captcha, error)
	UsePhoneCaptcha(captcha *model.Captcha) error
	SendPhoneCaptcha(phone string) error
}
