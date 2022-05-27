package service

import (
	"strings"
	"time"

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

func (svc *Service) GetPostComments(postID int64, sort string, offset, limit int) ([]*model.CommentFormated, int64, error) {
	conditions := &model.ConditionsT{
		"post_id": postID,
		"ORDER":   sort,
	}
	comments, err := svc.dao.GetComments(conditions, offset, limit)

	if err != nil {
		return nil, 0, err
	}

	userIDs := []int64{}
	commentIDs := []int64{}
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
		commentIDs = append(commentIDs, comment.ID)
	}

	users, err := svc.dao.GetUsersByIDs(userIDs)
	if err != nil {
		return nil, 0, err
	}

	contents, err := svc.dao.GetCommentContentsByIDs(commentIDs)
	if err != nil {
		return nil, 0, err
	}

	replies, err := svc.dao.GetCommentRepliesByID(commentIDs)
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
	totalRows, _ := svc.dao.GetCommentCount(conditions)

	return commentsFormated, totalRows, nil
}

func (svc *Service) CreatePostComment(userID int64, param CommentCreationReq) (*model.Comment, error) {
	// 加载Post
	post, err := svc.dao.GetPostByID(param.PostID)

	if err != nil {
		return nil, err
	}

	if post.CommentCount >= global.AppSetting.MaxCommentCount {
		return nil, errcode.MaxCommentCount
	}
	ip := svc.ctx.ClientIP()
	comment := &model.Comment{
		PostID: post.ID,
		UserID: userID,
		IP:     ip,
		IPLoc:  util.GetIPLoc(ip),
	}
	comment, err = svc.dao.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	for _, item := range param.Contents {
		// 检查附件是否是本站资源
		if item.Type == model.CONTENT_TYPE_IMAGE || item.Type == model.CONTENT_TYPE_VIDEO || item.Type == model.CONTENT_TYPE_ATTACHMENT {
			if strings.Index(item.Content, "https://"+global.AliossSetting.AliossDomain) != 0 {
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
		svc.dao.CreateCommentContent(postContent)
	}

	// 更新Post回复数
	post.CommentCount++
	post.LatestRepliedOn = time.Now().Unix()
	svc.dao.UpdatePost(post)

	// 更新索引
	go svc.PushPostToSearch(post)

	// 创建用户消息提醒
	postMaster, err := svc.dao.GetUserByID(post.UserID)
	if err == nil && postMaster.ID != userID {
		go svc.dao.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: postMaster.ID,
			Type:           model.MESSAGE_COMMENT,
			Breif:          "在泡泡中评论了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
		})
	}
	for _, u := range param.Users {
		user, err := svc.dao.GetUserByUsername(u)
		if err != nil || user.ID == userID || user.ID == postMaster.ID {
			continue
		}

		// 创建消息提醒
		go svc.dao.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: user.ID,
			Type:           model.MESSAGE_COMMENT,
			Breif:          "在泡泡评论中@了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
		})
	}

	return comment, nil
}

func (svc *Service) GetPostComment(id int64) (*model.Comment, error) {
	return svc.dao.GetCommentByID(id)
}

func (svc *Service) DeletePostComment(comment *model.Comment) error {
	// 加载post
	post, err := svc.dao.GetPostByID(comment.PostID)
	if err == nil {
		// 更新post回复数
		post.CommentCount--
		svc.dao.UpdatePost(post)
	}

	return svc.dao.DeleteComment(comment)
}

func (svc *Service) CreatePostCommentReply(commentID int64, content string, userID, atUserID int64) (*model.CommentReply, error) {
	// 加载Comment
	comment, err := svc.dao.GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}

	// 加载comment的post
	post, err := svc.dao.GetPostByID(comment.PostID)
	if err != nil {
		return nil, err
	}

	if post.CommentCount >= global.AppSetting.MaxCommentCount {
		return nil, errcode.MaxCommentCount
	}

	if userID == atUserID {
		atUserID = 0
	}

	if atUserID > 0 {
		// 检测目前用户是否存在
		users, _ := svc.dao.GetUsersByIDs([]int64{atUserID})
		if len(users) == 0 {
			atUserID = 0
		}
	}

	// 创建评论
	ip := svc.ctx.ClientIP()
	reply := &model.CommentReply{
		CommentID: commentID,
		UserID:    userID,
		Content:   content,
		AtUserID:  atUserID,
		IP:        ip,
		IPLoc:     util.GetIPLoc(ip),
	}

	reply, err = svc.dao.CreateCommentReply(reply)
	if err != nil {
		return nil, err
	}

	// 更新Post回复数
	post.CommentCount++
	post.LatestRepliedOn = time.Now().Unix()
	svc.dao.UpdatePost(post)

	// 更新索引
	go svc.PushPostToSearch(post)

	// 创建用户消息提醒
	commentMaster, err := svc.dao.GetUserByID(comment.UserID)
	if err == nil && commentMaster.ID != userID {
		go svc.dao.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: commentMaster.ID,
			Type:           model.MESSAGE_REPLY,
			Breif:          "在泡泡评论下回复了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
			ReplyID:        reply.ID,
		})
	}
	postMaster, err := svc.dao.GetUserByID(post.UserID)
	if err == nil && postMaster.ID != userID && commentMaster.ID != postMaster.ID {
		go svc.dao.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: postMaster.ID,
			Type:           model.MESSAGE_REPLY,
			Breif:          "在泡泡评论下发布了新回复",
			PostID:         post.ID,
			CommentID:      comment.ID,
			ReplyID:        reply.ID,
		})
	}
	if atUserID > 0 {
		user, err := svc.dao.GetUserByID(atUserID)
		if err == nil && user.ID != userID && commentMaster.ID != user.ID && postMaster.ID != user.ID {
			// 创建消息提醒
			go svc.dao.CreateMessage(&model.Message{
				SenderUserID:   userID,
				ReceiverUserID: user.ID,
				Type:           model.MESSAGE_REPLY,
				Breif:          "在泡泡评论的回复中@了你",
				PostID:         post.ID,
				CommentID:      comment.ID,
				ReplyID:        reply.ID,
			})
		}
	}

	return reply, nil
}

func (svc *Service) GetPostCommentReply(id int64) (*model.CommentReply, error) {
	return svc.dao.GetCommentReplyByID(id)
}

func (svc *Service) DeletePostCommentReply(reply *model.CommentReply) error {
	err := svc.dao.DeleteCommentReply(reply)
	if err != nil {
		return err
	}

	// 加载Comment
	comment, err := svc.dao.GetCommentByID(reply.CommentID)
	if err != nil {
		return err
	}

	// 加载comment的post
	post, err := svc.dao.GetPostByID(comment.PostID)
	if err != nil {
		return err
	}

	// 更新Post回复数
	post.CommentCount--
	post.LatestRepliedOn = time.Now().Unix()
	svc.dao.UpdatePost(post)

	// 更新索引
	go svc.PushPostToSearch(post)

	return nil
}
