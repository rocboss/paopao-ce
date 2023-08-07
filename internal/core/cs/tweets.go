// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	// 推文内容分快类型
	TweetBlockTitle TweetBlockType = iota + 1
	TweetBlockText
	TweetBlockImage
	TweetBlockVideo
	TweetBlockAudio
	TweetBlockLink
	TweetBlockAttachment
	TweetBlockChargeAttachment

	// 推文可见性
	TweetVisitPublic TweetVisibleType = iota
	TweetVisitPrivate
	TweetVisitFriend
	TweetVisitInvalid

	// 附件类型
	AttachmentTypeImage AttachmentType = iota + 1
	AttachmentTypeVideo
	AttachmentTypeOther
)

type (
	// TweetBlockType 推文内容分块类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源
	// TODO: 优化一下类型为 uint8， 需要底层数据库同步修改
	TweetBlockType int

	// TweetVisibleType 推文可见性，0公开，1私密，2好友
	TweetVisibleType uint8

	// AttachmentType 附件类型， 1图片， 2视频， 3其他
	// TODO: 优化一下类型为 uint8， 需要底层数据库同步修改
	AttachmentType int

	// TweetList 推文列表
	TweetList []*TweetItem

	// TweetInfoList 推文信息列表
	TweetInfoList []*TweetInfo

	// FavoriteList 收藏列表
	FavoriteList []*FavoriteItem

	// ReactionList 点赞列表
	ReactionList []*ReactionItem

	// TweetBlockList 推文分块列表
	TweetBlockList []*TweetBlock
)

// TweetBlock 推文分块
type TweetBlock struct {
	ID      int64          `json:"id" binding:"-"`
	PostID  int64          `json:"post_id" binding:"-"`
	Content string         `json:"content" binding:"required"`
	Type    TweetBlockType `json:"type" binding:"required"`
	Sort    int64          `json:"sort" binding:"required"`
}

// TweetInfo 推文信息
type TweetInfo struct {
	ID              int64            `json:"id"`
	UserID          int64            `json:"user_id"`
	CommentCount    int64            `json:"comment_count"`
	CollectionCount int64            `json:"collection_count"`
	UpvoteCount     int64            `json:"upvote_count"`
	Visibility      TweetVisibleType `json:"visibility"`
	IsTop           int              `json:"is_top"`
	IsEssence       int              `json:"is_essence"`
	IsLock          int              `json:"is_lock"`
	LatestRepliedOn int64            `json:"latest_replied_on"`
	Tags            string           `json:"tags"`
	AttachmentPrice int64            `json:"attachment_price"`
	IP              string           `json:"ip"`
	IPLoc           string           `json:"ip_loc"`
	CreatedOn       int64            `json:"created_on"`
	ModifiedOn      int64            `json:"modified_on"`
}

// TweetItem 一条推文信息
type TweetItem struct {
	ID              int64            `json:"id"`
	UserID          int64            `json:"user_id"`
	User            *UserInfo        `db:"user" json:"user"`
	Contents        []*TweetBlock    `db:"-" json:"contents"`
	CommentCount    int64            `json:"comment_count"`
	CollectionCount int64            `json:"collection_count"`
	UpvoteCount     int64            `json:"upvote_count"`
	Visibility      TweetVisibleType `json:"visibility"`
	IsTop           int              `json:"is_top"`
	IsEssence       int              `json:"is_essence"`
	IsLock          int              `json:"is_lock"`
	LatestRepliedOn int64            `json:"latest_replied_on"`
	CreatedOn       int64            `json:"created_on"`
	ModifiedOn      int64            `json:"modified_on"`
	Tags            map[string]int8  `json:"tags"`
	AttachmentPrice int64            `json:"attachment_price"`
	IPLoc           string           `json:"ip_loc"`
}

type Attachment struct {
	ID        int64          `json:"id"`
	UserID    int64          `json:"user_id"`
	FileSize  int64          `json:"file_size"`
	ImgWidth  int            `json:"img_width"`
	ImgHeight int            `json:"img_height"`
	Type      AttachmentType `json:"type"`
	Content   string         `json:"content"`
}

// Favorite 收藏
type FavoriteItem struct {
	ID      int64      `json:"id"`
	Tweet   *TweetInfo `json:"-"`
	TweetID int64      `json:"post_id"`
	UserID  int64      `json:"user_id"`
}

// Reaction 反应、表情符号， 点赞、喜欢等
type ReactionItem struct {
	ID      int64      `json:"id"`
	Tweet   *TweetInfo `json:"-"`
	TweetID int64      `json:"post_id"`
	UserID  int64      `json:"user_id"`
}

type NewTweetReq struct {
	Contents        TweetBlockList   `json:"contents" binding:"required"`
	Tags            []string         `json:"tags" binding:"required"`
	Users           []string         `json:"users" binding:"required"`
	AttachmentPrice int64            `json:"attachment_price"`
	Visibility      TweetVisibleType `json:"visibility"`
	ClientIP        string           `json:"-" binding:"-"`
}
