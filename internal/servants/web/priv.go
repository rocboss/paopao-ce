// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"image"
	"io"
	"strings"
	"time"

	"github.com/alimy/mir/v4"
	"github.com/alimy/tryst/cfg"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Priv      = (*privSrv)(nil)
	_ api.PrivChain = (*privChain)(nil)

	_uploadAttachmentTypeMap = map[string]ms.AttachmentType{
		"public/image":  ms.AttachmentTypeImage,
		"public/avatar": ms.AttachmentTypeImage,
		"public/video":  ms.AttachmentTypeVideo,
		"attachment":    ms.AttachmentTypeOther,
	}
)

type privSrv struct {
	api.UnimplementedPrivServant
	*base.DaoServant

	oss core.ObjectStorageService
}

type privChain struct {
	api.UnimplementedPrivChain
}

func (s *privChain) ChainCreateTweet() (res gin.HandlersChain) {
	if cfg.If("UseAuditHook") {
		res = gin.HandlersChain{chain.AuditHook()}
	}
	return
}

func (s *privSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT(), chain.Priv()}
}

func (s *privSrv) ThumbsDownTweetReply(req *web.TweetReplyThumbsReq) mir.Error {
	if err := s.Ds.ThumbsDownReply(req.Uid, req.TweetId, req.CommentId, req.ReplyId); err != nil {
		logrus.Errorf("thumbs down tweet reply error: %s req:%v", err, req)
		return web.ErrThumbsDownTweetReply
	}
	return nil
}

func (s *privSrv) ThumbsUpTweetReply(req *web.TweetReplyThumbsReq) mir.Error {
	if err := s.Ds.ThumbsUpReply(req.Uid, req.TweetId, req.CommentId, req.ReplyId); err != nil {
		logrus.Errorf("thumbs up tweet reply error: %s req:%v", err, req)
		return web.ErrThumbsUpTweetReply
	}
	return nil
}

func (s *privSrv) ThumbsDownTweetComment(req *web.TweetCommentThumbsReq) mir.Error {
	if err := s.Ds.ThumbsDownComment(req.Uid, req.TweetId, req.CommentId); err != nil {
		logrus.Errorf("thumbs down tweet comment error: %s req:%v", err, req)
		return web.ErrThumbsDownTweetComment
	}
	// 缓存处理
	onCommentActionEvent(req.TweetId, req.CommentId, _commentActionThumbsDown)
	return nil
}

func (s *privSrv) ThumbsUpTweetComment(req *web.TweetCommentThumbsReq) mir.Error {
	if err := s.Ds.ThumbsUpComment(req.Uid, req.TweetId, req.CommentId); err != nil {
		logrus.Errorf("thumbs up tweet comment error: %s req:%v", err, req)
		return web.ErrThumbsUpTweetComment
	}
	// 缓存处理
	onCommentActionEvent(req.TweetId, req.CommentId, _commentActionThumbsUp)
	return nil
}

func (s *privSrv) UnfollowTopic(req *web.UnfollowTopicReq) mir.Error {
	if err := s.Ds.UnfollowTopic(req.Uid, req.TopicId); err != nil {
		logrus.Errorf("user(%d) unfollow topic(%d) failed: %s", req.Uid, req.TopicId, err)
		return web.ErrUnfollowTopicFailed
	}
	return nil
}

func (s *privSrv) FollowTopic(req *web.FollowTopicReq) mir.Error {
	if err := s.Ds.FollowTopic(req.Uid, req.TopicId); err != nil {
		logrus.Errorf("user(%d) follow topic(%d) failed: %s", req.Uid, req.TopicId, err)
		return web.ErrFollowTopicFailed
	}
	return nil
}

func (s *privSrv) StickTopic(req *web.StickTopicReq) (*web.StickTopicResp, mir.Error) {
	status, err := s.Ds.StickTopic(req.Uid, req.TopicId)
	if err != nil {
		logrus.Errorf("user(%d) stick topic(%d) failed: %s", req.Uid, req.TopicId, err)
		return nil, web.ErrStickTopicFailed
	}
	return &web.StickTopicResp{
		StickStatus: status,
	}, nil
}

func (s *privSrv) PinTopic(req *web.PinTopicReq) (*web.PinTopicResp, mir.Error) {
	status, err := s.Ds.PinTopic(req.Uid, req.TopicId)
	if err != nil {
		logrus.Errorf("user(%d) pin topic(%d) failed: %s", req.Uid, req.TopicId, err)
		return nil, web.ErrPinTopicFailed
	}
	return &web.PinTopicResp{
		PinStatus: status,
	}, nil
}

func (s *privSrv) UploadAttachment(req *web.UploadAttachmentReq) (*web.UploadAttachmentResp, mir.Error) {
	defer req.File.Close()

	// 生成随机路径
	randomPath := uuid.Must(uuid.NewV4()).String()
	ossSavePath := req.UploadType + "/" + generatePath(randomPath[:8]) + "/" + randomPath[9:] + req.FileExt
	// NOTE: 注意这里将req.File Wrap到一个io.Reader的实例对象中是为了避免下游接口去主动调Close，req.File本身是实现了
	// io.Closer接口的，有的下游接口会断言传参是否实现了io.Closer接口，如果实现了会主动去调，我们这里因为下文中可能还要继续
	// 使用req.File所以应避免下游Close，否则会出现潜在的bug，比如这里的场景就是传一个超大的图片(>10MB)可能就会触发bug了。
	data := io.NopCloser(req.File)
	objectUrl, err := s.oss.PutObject(ossSavePath, data, req.FileSize, req.ContentType, false)
	if err != nil {
		logrus.Errorf("oss.putObject err: %s", err)
		return nil, web.ErrFileUploadFailed
	}

	// 构造附件Model
	attachment := &ms.Attachment{
		UserID:   req.Uid,
		FileSize: req.FileSize,
		Content:  objectUrl,
		Type:     _uploadAttachmentTypeMap[req.UploadType],
	}
	if attachment.Type == ms.AttachmentTypeImage {
		var src image.Image
		src, err = imaging.Decode(req.File)
		if err == nil {
			attachment.ImgWidth, attachment.ImgHeight = getImageSize(src.Bounds())
		}
	}
	attachment.ID, err = s.Ds.CreateAttachment(attachment)
	if err != nil {
		logrus.Errorf("Ds.CreateAttachment err: %s", err)
		return nil, web.ErrFileUploadFailed
	}

	return &web.UploadAttachmentResp{
		UserID:    req.Uid,
		FileSize:  req.FileSize,
		ImgWidth:  attachment.ImgWidth,
		ImgHeight: attachment.ImgHeight,
		Type:      attachment.Type,
		Content:   attachment.Content,
	}, nil
}

func (s *privSrv) DownloadAttachmentPrecheck(req *web.DownloadAttachmentPrecheckReq) (*web.DownloadAttachmentPrecheckResp, mir.Error) {
	content, err := s.Ds.GetPostContentByID(req.ContentID)
	if err != nil {
		logrus.Errorf("Ds.GetPostContentByID err: %s", err)
		return nil, web.ErrInvalidDownloadReq
	}
	resp := &web.DownloadAttachmentPrecheckResp{Paid: true}
	if content.Type == ms.ContentTypeChargeAttachment {
		tweet, err := s.GetTweetBy(content.PostID)
		if err != nil {
			logrus.Errorf("get tweet err: %v", err)
			return nil, web.ErrInvalidDownloadReq
		}
		// 发布者或管理员免费下载
		if tweet.UserID == req.User.ID || req.User.IsAdmin {
			return resp, nil
		}
		// 检测是否有购买记录
		resp.Paid = s.checkPostAttachmentIsPaid(req.ContentID, req.User.ID)
	}
	return resp, nil
}

func (s *privSrv) DownloadAttachment(req *web.DownloadAttachmentReq) (*web.DownloadAttachmentResp, mir.Error) {
	content, err := s.Ds.GetPostContentByID(req.ContentID)
	if err != nil {
		logrus.Errorf("s.GetPostContentByID err: %v", err)
		return nil, web.ErrInvalidDownloadReq
	}
	// 收费附件
	if content.Type == ms.ContentTypeChargeAttachment {
		post, err := s.GetTweetBy(content.PostID)
		if err != nil {
			logrus.Errorf("s.GetTweetBy err: %v", err)
			return nil, xerror.ServerError
		}
		paidFlag := false
		// 发布者或管理员免费下载 或者 检测是否有购买记录
		if post.UserID == req.User.ID || req.User.IsAdmin || s.checkPostAttachmentIsPaid(post.ID, req.User.ID) {
			paidFlag = true
		}
		// 未购买，则尝试购买
		if !paidFlag {
			err := s.buyPostAttachment(&ms.Post{
				Model: &ms.Model{
					ID: post.ID,
				},
				UserID:          post.UserID,
				AttachmentPrice: post.AttachmentPrice,
			}, req.User)
			if err != nil {
				return nil, err
			}
		}
	}
	// 签发附件下载链接
	objectKey := s.oss.ObjectKey(content.Content)
	signedURL, err := s.oss.SignURL(objectKey, 60)
	if err != nil {
		logrus.Errorf("client.SignURL err: %v", err)
		return nil, web.ErrDownloadReqError
	}
	return &web.DownloadAttachmentResp{
		SignedURL: signedURL,
	}, nil
}

func (s *privSrv) CreateTweet(req *web.CreateTweetReq) (_ *web.CreateTweetResp, xerr mir.Error) {
	var mediaContents []string
	defer func() {
		if xerr != nil {
			deleteOssObjects(s.oss, mediaContents)
		}
	}()

	contents, err := persistMediaContents(s.oss, req.Contents)
	if err != nil {
		return nil, web.ErrCreatePostFailed
	}
	mediaContents = contents
	tags := tagsFrom(req.Tags)
	post := &ms.Post{
		UserID:          req.User.ID,
		Tags:            strings.Join(tags, ","),
		IP:              req.ClientIP,
		IPLoc:           utils.GetIPLoc(req.ClientIP),
		AttachmentPrice: req.AttachmentPrice,
		Visibility:      ms.PostVisibleT(req.Visibility.ToVisibleValue()),
	}
	post, err = s.Ds.CreatePost(post)
	if err != nil {
		logrus.Errorf("Ds.CreatePost err: %s", err)
		return nil, web.ErrCreatePostFailed
	}

	// 创建推文内容
	for _, item := range req.Contents {
		if err := item.Check(s.Ds); err != nil {
			// 属性非法
			logrus.Infof("contents check err: %s", err)
			continue
		}
		if item.Type == ms.ContentTypeAttachment && req.AttachmentPrice > 0 {
			item.Type = ms.ContentTypeChargeAttachment
		}
		postContent := &ms.PostContent{
			PostID:  post.ID,
			UserID:  req.User.ID,
			Content: item.Content,
			Type:    item.Type,
			Sort:    item.Sort,
		}
		if _, err = s.Ds.CreatePostContent(postContent); err != nil {
			logrus.Infof("Ds.CreatePostContent err: %s", err)
			return nil, web.ErrCreateCommentFailed
		}
	}

	// 私密推文不创建标签与用户提醒
	if post.Visibility != core.PostVisitPrivate {
		// 创建标签
		s.Ds.UpsertTags(req.User.ID, tags)

		// 创建用户消息提醒
		for _, u := range req.Users {
			user, err := s.Ds.GetUserByUsername(u)
			if err != nil || user.ID == req.User.ID {
				continue
			}

			// 创建消息提醒
			onCreateMessageEvent(&ms.Message{
				SenderUserID:   req.User.ID,
				ReceiverUserID: user.ID,
				Type:           ms.MsgTypePost,
				Brief:          "在新发布的泡泡动态中@了你",
				PostID:         post.ID,
			})
		}
	}
	// 推送Search
	s.PushPostToSearch(post)
	formatedPosts, err := s.Ds.RevampPosts([]*ms.PostFormated{post.Format()})
	if err != nil {
		logrus.Infof("Ds.RevampPosts err: %s", err)
		return nil, web.ErrCreatePostFailed
	}
	// 缓存处理
	// TODO: 缓存逻辑合并处理
	onTrendsActionEvent(_trendsActionCreateTweet, req.User.ID)
	onTweetActionEvent(_tweetActionCreate, req.User.ID, req.User.Username)
	return (*web.CreateTweetResp)(formatedPosts[0]), nil
}

func (s *privSrv) DeleteTweet(req *web.DeleteTweetReq) mir.Error {
	if req.User == nil {
		return web.ErrNoPermission
	}
	post, err := s.Ds.GetPostByID(req.ID)
	if err != nil {
		logrus.Errorf("Ds.GetPostByID err: %s", err)
		return web.ErrGetPostFailed
	}
	if post.UserID != req.User.ID && !req.User.IsAdmin {
		return web.ErrNoPermission
	}
	mediaContents, err := s.Ds.DeletePost(post)
	if err != nil {
		logrus.Errorf("Ds.DeletePost delete post failed: %s", err)
		return web.ErrDeletePostFailed
	}
	// 删除推文的媒体内容
	deleteOssObjects(s.oss, mediaContents)
	// 删除索引
	s.DeleteSearchPost(post)
	if err != nil {
		logrus.Errorf("s.DeleteSearchPost failed: %s", err)
		return web.ErrDeletePostFailed
	}
	// 缓存处理
	// TODO: 缓存逻辑合并处理
	onTrendsActionEvent(_trendsActionDeleteTweet, req.User.ID)
	onTweetActionEvent(_tweetActionDelete, req.User.ID, req.User.Username)
	return nil
}

func (s *privSrv) DeleteCommentReply(req *web.DeleteCommentReplyReq) mir.Error {
	reply, err := s.Ds.GetCommentReplyByID(req.ID)
	if err != nil {
		logrus.Errorf("Ds.GetCommentReplyByID err: %s", err)
		return web.ErrGetReplyFailed
	}
	if req.User.ID != reply.UserID && !req.User.IsAdmin {
		return web.ErrNoPermission
	}
	// 执行删除
	err = s.deletePostCommentReply(reply)
	if err != nil {
		logrus.Errorf("s.deletePostCommentReply err: %s", err)
		return web.ErrDeleteCommentFailed
	}
	// 缓存处理， 宽松处理错误
	if comment, err := s.Ds.GetCommentByID(reply.CommentID); err == nil {
		onCommentActionEvent(comment.PostID, comment.ID, _commentActionReplyDelete)
	}
	return nil
}

func (s *privSrv) CreateCommentReply(req *web.CreateCommentReplyReq) (_ *web.CreateCommentReplyResp, xerr mir.Error) {
	var (
		post     *ms.Post
		comment  *ms.Comment
		atUserID int64
		err      error
	)

	if post, comment, atUserID, err = s.createPostPreHandler(req.CommentID, req.Uid, req.AtUserID); err != nil {
		return nil, web.ErrCreateReplyFailed
	}

	// 创建评论
	reply := &ms.CommentReply{
		CommentID: req.CommentID,
		UserID:    req.Uid,
		Content:   req.Content,
		AtUserID:  atUserID,
		IP:        req.ClientIP,
		IPLoc:     utils.GetIPLoc(req.ClientIP),
	}

	reply, err = s.Ds.CreateCommentReply(reply)
	if err != nil {
		return nil, web.ErrCreateReplyFailed
	}

	// 更新Post回复数
	post.CommentCount++
	post.LatestRepliedOn = time.Now().Unix()
	s.Ds.UpdatePost(post)

	// 更新索引
	s.PushPostToSearch(post)

	// 创建用户消息提醒
	commentMaster, err := s.Ds.GetUserByID(comment.UserID)
	if err == nil && commentMaster.ID != req.Uid {
		onCreateMessageEvent(&ms.Message{
			SenderUserID:   req.Uid,
			ReceiverUserID: commentMaster.ID,
			Type:           ms.MsgTypeReply,
			Brief:          "在泡泡评论下回复了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
			ReplyID:        reply.ID,
		})
	}
	postMaster, err := s.Ds.GetUserByID(post.UserID)
	if err == nil && postMaster.ID != req.Uid && commentMaster.ID != postMaster.ID {
		onCreateMessageEvent(&ms.Message{
			SenderUserID:   req.Uid,
			ReceiverUserID: postMaster.ID,
			Type:           ms.MsgTypeReply,
			Brief:          "在泡泡评论下发布了新回复",
			PostID:         post.ID,
			CommentID:      comment.ID,
			ReplyID:        reply.ID,
		})
	}
	if atUserID > 0 {
		user, err := s.Ds.GetUserByID(atUserID)
		if err == nil && user.ID != req.Uid && commentMaster.ID != user.ID && postMaster.ID != user.ID {
			// 创建消息提醒
			onCreateMessageEvent(&ms.Message{
				SenderUserID:   req.Uid,
				ReceiverUserID: user.ID,
				Type:           ms.MsgTypeReply,
				Brief:          "在泡泡评论的回复中@了你",
				PostID:         post.ID,
				CommentID:      comment.ID,
				ReplyID:        reply.ID,
			})
		}
	}
	// 缓存处理
	onCommentActionEvent(comment.PostID, comment.ID, _commentActionReplyCreate)
	return (*web.CreateCommentReplyResp)(reply), nil
}

func (s *privSrv) DeleteComment(req *web.DeleteCommentReq) mir.Error {
	comment, err := s.Ds.GetCommentByID(req.ID)
	if err != nil {
		logrus.Errorf("Ds.GetCommentByID err: %v\n", err)
		return web.ErrGetCommentFailed
	}
	if req.User.ID != comment.UserID && !req.User.IsAdmin {
		return web.ErrNoPermission
	}
	// 加载post
	post, err := s.Ds.GetPostByID(comment.PostID)
	if err != nil {
		return web.ErrDeleteCommentFailed
	}
	// 更新post回复数
	post.CommentCount--
	if err := s.Ds.UpdatePost(post); err != nil {
		logrus.Errorf("Ds.UpdatePost err: %s", err)
		return web.ErrDeleteCommentFailed
	}
	// TODO: 优化删除逻辑，事务化删除comment
	if err := s.Ds.DeleteComment(comment); err != nil {
		logrus.Errorf("Ds.DeleteComment err: %s", err)
		return web.ErrDeleteCommentFailed
	}
	onCommentActionEvent(comment.PostID, comment.ID, _commentActionDelete)
	return nil
}

func (s *privSrv) HighlightComment(req *web.HighlightCommentReq) (*web.HighlightCommentResp, mir.Error) {
	status, err := s.Ds.HighlightComment(req.Uid, req.CommentId)
	if err == cs.ErrNoPermission {
		return nil, web.ErrNoPermission
	} else if err != nil {
		return nil, web.ErrHighlightCommentFailed
	}
	// 缓存处理， 宽松处理错误
	if comment, err := s.Ds.GetCommentByID(req.CommentId); err == nil {
		onCommentActionEvent(comment.PostID, comment.ID, _commentActionHighlight)
	}
	return &web.HighlightCommentResp{
		HighlightStatus: status,
	}, nil
}

func (s *privSrv) CreateComment(req *web.CreateCommentReq) (_ *web.CreateCommentResp, xerr mir.Error) {
	var (
		mediaContents []string
		err           error
	)
	defer func() {
		if xerr != nil {
			deleteOssObjects(s.oss, mediaContents)
		}
	}()

	if mediaContents, err = persistMediaContents(s.oss, req.Contents); err != nil {
		return nil, xerror.ServerError
	}

	// 加载Post
	post, err := s.Ds.GetPostByID(req.PostID)
	if err != nil {
		logrus.Errorf("Ds.GetPostByID err:%s", err)
		return nil, xerror.ServerError
	}
	if post.CommentCount >= conf.AppSetting.MaxCommentCount {
		return nil, web.ErrMaxCommentCount
	}
	comment := &ms.Comment{
		PostID: post.ID,
		UserID: req.Uid,
		IP:     req.ClientIP,
		IPLoc:  utils.GetIPLoc(req.ClientIP),
	}
	comment, err = s.Ds.CreateComment(comment)
	if err != nil {
		logrus.Errorf("Ds.CreateComment err:%s", err)
		return nil, web.ErrCreateCommentFailed
	}

	for _, item := range req.Contents {
		// 检查附件是否是本站资源
		if item.Type == ms.ContentTypeImage || item.Type == ms.ContentTypeVideo || item.Type == ms.ContentTypeAttachment {
			if err := s.Ds.CheckAttachment(item.Content); err != nil {
				continue
			}
		}
		postContent := &ms.CommentContent{
			CommentID: comment.ID,
			UserID:    req.Uid,
			Content:   item.Content,
			Type:      item.Type,
			Sort:      item.Sort,
		}
		s.Ds.CreateCommentContent(postContent)
	}

	// 更新Post回复数
	post.CommentCount++
	post.LatestRepliedOn = time.Now().Unix()
	s.Ds.UpdatePost(post)

	// 更新索引
	s.PushPostToSearch(post)

	// 创建用户消息提醒
	postMaster, err := s.Ds.GetUserByID(post.UserID)
	if err == nil && postMaster.ID != req.Uid {
		onCreateMessageEvent(&ms.Message{
			SenderUserID:   req.Uid,
			ReceiverUserID: postMaster.ID,
			Type:           ms.MsgtypeComment,
			Brief:          "在泡泡中评论了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
		})
	}
	for _, u := range req.Users {
		user, err := s.Ds.GetUserByUsername(u)
		if err != nil || user.ID == req.Uid || user.ID == postMaster.ID {
			continue
		}

		// 创建消息提醒
		onCreateMessageEvent(&ms.Message{
			SenderUserID:   req.Uid,
			ReceiverUserID: user.ID,
			Type:           ms.MsgtypeComment,
			Brief:          "在泡泡评论中@了你",
			PostID:         post.ID,
			CommentID:      comment.ID,
		})
	}
	// 缓存处理
	onCommentActionEvent(comment.PostID, comment.ID, _commentActionCreate)
	return (*web.CreateCommentResp)(comment), nil
}

func (s *privSrv) CollectionTweet(req *web.CollectionTweetReq) (*web.CollectionTweetResp, mir.Error) {
	status := false
	collection, err := s.Ds.GetUserPostCollection(req.ID, req.Uid)
	if err != nil {
		// 创建Star
		if _, xerr := s.createPostCollection(req.ID, req.Uid); xerr != nil {
			return nil, xerr
		}
		status = true
	} else {
		// 取消Star
		if xerr := s.deletePostCollection(collection); xerr != nil {
			return nil, xerr
		}
	}
	return &web.CollectionTweetResp{
		Status: status,
	}, nil
}

func (s *privSrv) StarTweet(req *web.StarTweetReq) (*web.StarTweetResp, mir.Error) {
	status := false
	star, err := s.Ds.GetUserPostStar(req.ID, req.Uid)
	if err != nil {
		// 创建Star
		if _, xerr := s.createPostStar(req.ID, req.Uid); xerr != nil {
			return nil, xerr
		}
		status = true
	} else {
		// 取消Star
		if xerr := s.deletePostStar(star); xerr != nil {
			return nil, xerr
		}
	}
	return &web.StarTweetResp{
		Status: status,
	}, nil
}

func (s *privSrv) VisibleTweet(req *web.VisibleTweetReq) (*web.VisibleTweetResp, mir.Error) {
	if req.Visibility >= web.TweetVisitInvalid {
		return nil, xerror.InvalidParams
	}
	post, err := s.Ds.GetPostByID(req.ID)
	if err != nil {
		return nil, web.ErrVisblePostFailed
	}
	if xerr := checkPermision(req.User, post.UserID); xerr != nil {
		return nil, xerr
	}
	if err = s.Ds.VisiblePost(post, req.Visibility.ToVisibleValue()); err != nil {
		logrus.Warnf("s.Ds.VisiblePost: %s", err)
		return nil, web.ErrVisblePostFailed
	}

	// 推送Search
	post.Visibility = ms.PostVisibleT(req.Visibility.ToVisibleValue())
	s.PushPostToSearch(post)

	return &web.VisibleTweetResp{
		Visibility: req.Visibility,
	}, nil
}

func (s *privSrv) StickTweet(req *web.StickTweetReq) (*web.StickTweetResp, mir.Error) {
	post, err := s.Ds.GetPostByID(req.ID)
	if err != nil {
		logrus.Errorf("Ds.GetPostByID err: %v\n", err)
		return nil, web.ErrStickPostFailed
	}
	if !req.User.IsAdmin {
		return nil, web.ErrNoPermission
	}
	newStatus := 1 - post.IsTop
	if err = s.Ds.StickPost(post); err != nil {
		return nil, web.ErrStickPostFailed
	}
	return &web.StickTweetResp{
		StickStatus: newStatus,
	}, nil
}

func (s *privSrv) HighlightTweet(req *web.HighlightTweetReq) (res *web.HighlightTweetResp, err mir.Error) {
	if status, xerr := s.Ds.HighlightPost(req.User.ID, req.ID); xerr == nil {
		res = &web.HighlightTweetResp{
			HighlightStatus: status,
		}
	} else if xerr == cs.ErrNoPermission {
		err = web.ErrNoPermission
		logrus.Warnf("highlight tweet occurs no permision error with userId=%d postId=%d", req.User.ID, req.ID)
	} else {
		err = web.ErrHighlightPostFailed
		logrus.Warnf("highlight tweet err: %s with userId=%d postId=%d", xerr, req.User.ID, req.ID)
	}
	return
}

func (s *privSrv) LockTweet(req *web.LockTweetReq) (*web.LockTweetResp, mir.Error) {
	post, err := s.Ds.GetPostByID(req.ID)
	if err != nil {
		return nil, web.ErrLockPostFailed
	}
	if post.UserID != req.User.ID && !req.User.IsAdmin {
		return nil, web.ErrNoPermission
	}
	newStatus := 1 - post.IsLock
	if err := s.Ds.LockPost(post); err != nil {
		return nil, web.ErrLockPostFailed
	}
	return &web.LockTweetResp{
		LockStatus: newStatus,
	}, nil
}

func (s *privSrv) deletePostCommentReply(reply *ms.CommentReply) error {
	err := s.Ds.DeleteCommentReply(reply)
	if err != nil {
		return err
	}
	// 加载Comment
	comment, err := s.Ds.GetCommentByID(reply.CommentID)
	if err != nil {
		return err
	}
	// 加载comment的post
	post, err := s.Ds.GetPostByID(comment.PostID)
	if err != nil {
		return err
	}
	// 更新Post回复数
	post.CommentCount--
	post.LatestRepliedOn = time.Now().Unix()
	s.Ds.UpdatePost(post)
	// 更新索引
	s.PushPostToSearch(post)
	return nil
}

func (s *privSrv) createPostPreHandler(commentID int64, userID, atUserID int64) (*ms.Post, *ms.Comment, int64,
	error) {
	// 加载Comment
	comment, err := s.Ds.GetCommentByID(commentID)
	if err != nil {
		return nil, nil, atUserID, err
	}

	// 加载comment的post
	post, err := s.Ds.GetPostByID(comment.PostID)
	if err != nil {
		return nil, nil, atUserID, err
	}

	if post.CommentCount >= conf.AppSetting.MaxCommentCount {
		return nil, nil, atUserID, web.ErrMaxCommentCount
	}

	if userID == atUserID {
		atUserID = 0
	}

	if atUserID > 0 {
		// 检测目前用户是否存在
		users, _ := s.Ds.GetUsersByIDs([]int64{atUserID})
		if len(users) == 0 {
			atUserID = 0
		}
	}

	return post, comment, atUserID, nil
}

func (s *privSrv) createPostStar(postID, userID int64) (*ms.PostStar, mir.Error) {
	post, err := s.Ds.GetPostByID(postID)
	if err != nil {
		return nil, xerror.ServerError
	}

	// 私密post不可操作
	// TODO: 使用统一的permission checker来检查权限问题，这里好友可见post就没处理，是bug
	if post.Visibility == core.PostVisitPrivate && post.UserID != userID {
		return nil, web.ErrNoPermission
	}

	star, err := s.Ds.CreatePostStar(postID, userID)
	if err != nil {
		return nil, xerror.ServerError
	}

	// 更新Post点赞数
	post.UpvoteCount++
	s.Ds.UpdatePost(post)

	// 更新索引
	s.PushPostToSearch(post)
	return star, nil
}

func (s *privSrv) deletePostStar(star *ms.PostStar) mir.Error {
	post, err := s.Ds.GetPostByID(star.PostID)
	if err != nil {
		return xerror.ServerError
	}

	// 私密post特殊处理
	// TODO: 使用统一的permission checker来检查权限问题，这里好友可见post就没处理，是bug
	if post.Visibility == core.PostVisitPrivate && post.UserID != star.UserID {
		return web.ErrNoPermission
	}

	if err = s.Ds.DeletePostStar(star); err != nil {
		return xerror.ServerError
	}

	// 更新Post点赞数
	post.UpvoteCount--
	s.Ds.UpdatePost(post)

	// 更新索引
	s.PushPostToSearch(post)
	return nil
}

func (s *privSrv) createPostCollection(postID, userID int64) (*ms.PostCollection, mir.Error) {
	post, err := s.Ds.GetPostByID(postID)
	if err != nil {
		return nil, xerror.ServerError
	}

	// 私密post特殊处理
	// TODO: 使用统一的permission checker来检查权限问题，这里好友可见post就没处理，是bug
	if post.Visibility == core.PostVisitPrivate && post.UserID != userID {
		return nil, web.ErrNoPermission
	}

	collection, err := s.Ds.CreatePostCollection(postID, userID)
	if err != nil {
		return nil, xerror.ServerError
	}

	// 更新Post点赞数
	post.CollectionCount++
	s.Ds.UpdatePost(post)

	// 更新索引
	s.PushPostToSearch(post)
	return collection, nil
}

func (s *privSrv) deletePostCollection(collection *ms.PostCollection) mir.Error {
	post, err := s.Ds.GetPostByID(collection.PostID)
	if err != nil {
		return xerror.ServerError
	}

	// 私密post特殊处理
	// TODO: 使用统一的permission checker来检查权限问题，这里好友可见post就没处理，是bug
	if post.Visibility == core.PostVisitPrivate && post.UserID != collection.UserID {
		return web.ErrNoPermission
	}
	if err = s.Ds.DeletePostCollection(collection); err != nil {
		return xerror.ServerError
	}

	// 更新Post点赞数
	post.CollectionCount--
	s.Ds.UpdatePost(post)

	// 更新索引
	s.PushPostToSearch(post)
	return nil
}

func (s *privSrv) checkPostAttachmentIsPaid(postID, userID int64) bool {
	bill, err := s.Ds.GetPostAttatchmentBill(postID, userID)
	return err == nil && bill.Model != nil && bill.ID > 0
}

func (s *privSrv) buyPostAttachment(post *ms.Post, user *ms.User) mir.Error {
	if user.Balance < post.AttachmentPrice {
		return web.ErrInsuffientDownloadMoney
	}
	// 执行购买
	if err := s.Ds.HandlePostAttachmentBought(post, user); err != nil {
		logrus.Errorf("Ds.HandlePostAttachmentBought err: %s", err)
		return xerror.ServerError
	}
	return nil
}

func newPrivSrv(s *base.DaoServant, oss core.ObjectStorageService) api.Priv {
	return &privSrv{
		DaoServant: s,
		oss:        oss,
	}
}

func newPrivChain() api.PrivChain {
	return &privChain{}
}
