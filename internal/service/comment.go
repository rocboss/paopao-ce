package service

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/errcode"
	"github.com/rocboss/paopao-ce/pkg/util"
)

type CommentCreationReq struct {
	PostID   int64              `json:"post_id" binding:"required"`
	Contents []*PostContentItem `json:"contents" binding:"required"`
	Users    []string           `json:"users" binding:"required"`
}
type CommentReplyCreationReq struct {
	CommentID int64  `json:"comment_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
	AtUserID  int64  `json:"at_user_id"`
}

type CommentDelReq struct {
	ID int64 `json:"id" binding:"required"`
}
type ReplyDelReq struct {
	ID int64 `json:"id" binding:"required"`
}

func GetPostComments(postID int64, sort string, offset, limit int) ([]*model.CommentFormated, int64, error) {
	conditions := &model.ConditionsT{
		"post_id": postID,
		"ORDER":   sort,
	}
	comments, err := ds.GetComments(conditions, offset, limit)

	if err != nil {
		return nil, 0, err
	}

	userIDs := []int64{}
	commentIDs := []int64{}
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
		commentIDs = append(commentIDs, comment.ID)
	}

	users, err := ds.GetUsersByIDs(userIDs)
	if err != nil {
		return nil, 0, err
	}

	contents, err := ds.GetCommentContentsByIDs(commentIDs)
	if err != nil {
		return nil, 0, err
	}

	replies, err := ds.GetCommentRepliesByID(commentIDs)
	if err != nil {
		return nil, 0, err
	}

	commentsFormated := []*model.CommentFormated{}
	for _, comment := range comments {
		commentFormated := comment.Format()
		for _, content := range contents {
			if content.CommentID == comment.ID {
				commentFormated.Contents = append(commentFormated.Contents, content)
			}
		}
		for _, reply := range replies {
			if reply.CommentID == commentFormated.ID {
				commentFormated.Replies = append(commentFormated.Replies, reply)
			}
		}
		for _, user := range users {
			if user.ID == comment.UserID {
				commentFormated.User = user.Format()
			}
		}

		commentsFormated = append(commentsFormated, commentFormated)
	}

	// 获取总量
	totalRows, _ := ds.GetCommentCount(conditions)

	return commentsFormated, totalRows, nil
}

func CreatePostComment(ctx *gin.Context, userID int64, param CommentCreationReq) (*model.Comment, error) {
	// 加载Post
	post, err := ds.GetPostByID(param.PostID)

	if err != nil {
		return nil, err
	}

	if post.CommentCount >= global.AppSetting.MaxCommentCount {
		return nil, errcode.MaxCommentCount
	}
	ip := ctx.ClientIP()
	comment := &model.Comment{
		PostID: post.ID,
		UserID: userID,
		IP:     ip,
		IPLoc:  util.GetIPLoc(ip),
	}
	comment, err = ds.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	for _, item := range param.Contents {
		// 检查附件是否是本站资源
		if item.Type == model.CONTENT_TYPE_IMAGE || item.Type == model.CONTENT_TYPE_VIDEO || item.Type == model.CONTENT_TYPE_ATTACHMENT {
			if strings.Index(item.Content, "https://"+global.AliOSSSetting.Domain) != 0 {
				continue
			}
		}

		postContent := &model.CommentContent{
			CommentID: comment.ID,
			UserID:    userID,
			Content:   item.Content,
			Type:      item.Type,
			Sort:      item.Sort,
		}
		ds.CreateCommentContent(postContent)
	}

	// 更新Post回复数
	post.CommentCount++
	post.LatestRepliedOn = time.Now().Unix()
	ds.UpdatePost(post)

	// 更新索引
	go PushPostToSearch(post)

	// 创建用户消息提醒
	postMaster, err := ds.GetUserByID(post.UserID)
	if err == nil && postMaster.ID != userID {
		go ds.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: postMaster.ID,
			Type:           model.MESSAGE_COMMENT,
			Brief:          "在泡泡中评论了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
		})
	}
	for _, u := range param.Users {
		user, err := ds.GetUserByUsername(u)
		if err != nil || user.ID == userID || user.ID == postMaster.ID {
			continue
		}

		// 创建消息提醒
		go ds.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: user.ID,
			Type:           model.MESSAGE_COMMENT,
			Brief:          "在泡泡评论中@了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
		})
	}

	return comment, nil
}

func GetPostComment(id int64) (*model.Comment, error) {
	return ds.GetCommentByID(id)
}

func DeletePostComment(comment *model.Comment) error {
	// 加载post
	post, err := ds.GetPostByID(comment.PostID)
	if err == nil {
		// 更新post回复数
		post.CommentCount--
		ds.UpdatePost(post)
	}

	return ds.DeleteComment(comment)
}

func createPostPreHandler(commentID int64, userID, atUserID int64) (*model.Post, *model.Comment, int64,
	error) {
	// 加载Comment
	comment, err := ds.GetCommentByID(commentID)
	if err != nil {
		return nil, nil, atUserID, err
	}

	// 加载comment的post
	post, err := ds.GetPostByID(comment.PostID)
	if err != nil {
		return nil, nil, atUserID, err
	}

	if post.CommentCount >= global.AppSetting.MaxCommentCount {
		return nil, nil, atUserID, errcode.MaxCommentCount
	}

	if userID == atUserID {
		atUserID = 0
	}

	if atUserID > 0 {
		// 检测目前用户是否存在
		users, _ := ds.GetUsersByIDs([]int64{atUserID})
		if len(users) == 0 {
			atUserID = 0
		}
	}

	return post, comment, atUserID, nil
}

func CreatePostCommentReply(ctx *gin.Context, commentID int64, content string, userID, atUserID int64) (*model.CommentReply, error) {
	var (
		post    *model.Post
		comment *model.Comment
		err     error
	)
	if post, comment, atUserID, err = createPostPreHandler(commentID, userID, atUserID); err != nil {
		return nil, err
	}

	// 创建评论
	ip := ctx.ClientIP()
	reply := &model.CommentReply{
		CommentID: commentID,
		UserID:    userID,
		Content:   content,
		AtUserID:  atUserID,
		IP:        ip,
		IPLoc:     util.GetIPLoc(ip),
	}

	reply, err = ds.CreateCommentReply(reply)
	if err != nil {
		return nil, err
	}

	// 更新Post回复数
	post.CommentCount++
	post.LatestRepliedOn = time.Now().Unix()
	ds.UpdatePost(post)

	// 更新索引
	go PushPostToSearch(post)

	// 创建用户消息提醒
	commentMaster, err := ds.GetUserByID(comment.UserID)
	if err == nil && commentMaster.ID != userID {
		go ds.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: commentMaster.ID,
			Type:           model.MESSAGE_REPLY,
			Brief:          "在泡泡评论下回复了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
			ReplyID:        reply.ID,
		})
	}
	postMaster, err := ds.GetUserByID(post.UserID)
	if err == nil && postMaster.ID != userID && commentMaster.ID != postMaster.ID {
		go ds.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: postMaster.ID,
			Type:           model.MESSAGE_REPLY,
			Brief:          "在泡泡评论下发布了新回复",
			PostID:         post.ID,
			CommentID:      comment.ID,
			ReplyID:        reply.ID,
		})
	}
	if atUserID > 0 {
		user, err := ds.GetUserByID(atUserID)
		if err == nil && user.ID != userID && commentMaster.ID != user.ID && postMaster.ID != user.ID {
			// 创建消息提醒
			go ds.CreateMessage(&model.Message{
				SenderUserID:   userID,
				ReceiverUserID: user.ID,
				Type:           model.MESSAGE_REPLY,
				Brief:          "在泡泡评论的回复中@了你",
				PostID:         post.ID,
				CommentID:      comment.ID,
				ReplyID:        reply.ID,
			})
		}
	}

	return reply, nil
}

func GetPostCommentReply(id int64) (*model.CommentReply, error) {
	return ds.GetCommentReplyByID(id)
}

func DeletePostCommentReply(reply *model.CommentReply) error {
	err := ds.DeleteCommentReply(reply)
	if err != nil {
		return err
	}

	// 加载Comment
	comment, err := ds.GetCommentByID(reply.CommentID)
	if err != nil {
		return err
	}

	// 加载comment的post
	post, err := ds.GetPostByID(comment.PostID)
	if err != nil {
		return err
	}

	// 更新Post回复数
	post.CommentCount--
	post.LatestRepliedOn = time.Now().Unix()
	ds.UpdatePost(post)

	// 更新索引
	go PushPostToSearch(post)

	return nil
}
