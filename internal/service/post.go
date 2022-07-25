package service

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/model/rest"
	"github.com/rocboss/paopao-ce/pkg/errcode"
	"github.com/rocboss/paopao-ce/pkg/util"
	"github.com/sirupsen/logrus"
)

type TagType string

const TagTypeHot TagType = "hot"
const TagTypeNew TagType = "new"

type PostListReq struct {
	Conditions *model.ConditionsT
	Offset     int
	Limit      int
}
type PostTagsReq struct {
	Type TagType `json:"type" form:"type" binding:"required"`
	Num  int     `json:"num" form:"num" binding:"required"`
}

type PostCreationReq struct {
	Contents        []*PostContentItem `json:"contents" binding:"required"`
	Tags            []string           `json:"tags" binding:"required"`
	Users           []string           `json:"users" binding:"required"`
	AttachmentPrice int64              `json:"attachment_price"`
	Visibility      model.PostVisibleT `json:"visibility"`
}

type PostDelReq struct {
	ID int64 `json:"id" binding:"required"`
}

type PostLockReq struct {
	ID int64 `json:"id" binding:"required"`
}

type PostStickReq struct {
	ID int64 `json:"id" binding:"required"`
}

type PostVisibilityReq struct {
	ID         int64              `json:"id" binding:"required"`
	Visibility model.PostVisibleT `json:"visibility"`
}

type PostStarReq struct {
	ID int64 `json:"id" binding:"required"`
}

type PostCollectionReq struct {
	ID int64 `json:"id" binding:"required"`
}

type PostContentItem struct {
	Content string             `json:"content"  binding:"required"`
	Type    model.PostContentT `json:"type"  binding:"required"`
	Sort    int64              `json:"sort"  binding:"required"`
}

// Check 检查PostContentItem属性
func (p *PostContentItem) Check() error {
	// 检查附件是否是本站资源
	if p.Type == model.CONTENT_TYPE_IMAGE || p.Type == model.CONTENT_TYPE_VIDEO || p.Type == model.
		CONTENT_TYPE_ATTACHMENT {
		if err := ds.CheckAttachment(p.Content); err != nil {
			return err
		}
	}
	// 检查链接是否合法
	if p.Type == model.CONTENT_TYPE_LINK {
		if strings.Index(p.Content, "http://") != 0 && strings.Index(p.Content, "https://") != 0 {
			return fmt.Errorf("链接不合法")
		}
	}

	return nil
}

func tagsFrom(originTags []string) []string {
	tags := make([]string, 0, len(originTags))
	for _, tag := range originTags {
		// TODO: 优化tag有效性检测
		if tag = strings.TrimSpace(tag); len(tag) > 0 {
			tags = append(tags, tag)
		}
	}
	return tags
}

// CreatePost 创建文章
// TODO: 推文+推文内容需要在一个事务中添加，后续优化
func CreatePost(c *gin.Context, userID int64, param PostCreationReq) (post *model.Post, err error) {
	var mediaContents []string

	defer func() {
		if err != nil {
			deleteOssObjects(mediaContents)
		}
	}()

	if mediaContents, err = persistMediaContents(param.Contents); err != nil {
		return
	}

	ip := c.ClientIP()
	tags := tagsFrom(param.Tags)
	post = &model.Post{
		UserID:          userID,
		Tags:            strings.Join(tags, ","),
		IP:              ip,
		IPLoc:           util.GetIPLoc(ip),
		AttachmentPrice: param.AttachmentPrice,
		Visibility:      param.Visibility,
	}
	post, err = ds.CreatePost(post)
	if err != nil {
		return nil, err
	}

	for _, item := range param.Contents {
		if err = item.Check(); err != nil {
			// 属性非法
			logrus.Infof("contents check err: %v", err)
			continue
		}

		if item.Type == model.CONTENT_TYPE_ATTACHMENT && param.AttachmentPrice > 0 {
			item.Type = model.CONTENT_TYPE_CHARGE_ATTACHMENT
		}

		postContent := &model.PostContent{
			PostID:  post.ID,
			UserID:  userID,
			Content: item.Content,
			Type:    item.Type,
			Sort:    item.Sort,
		}
		if _, err = ds.CreatePostContent(postContent); err != nil {
			return nil, err
		}
	}

	// 私密推文不创建标签与用户提醒
	if post.Visibility != model.PostVisitPrivate {
		// 创建标签
		for _, t := range tags {
			tag := &model.Tag{
				UserID: userID,
				Tag:    t,
			}
			ds.CreateTag(tag)
		}

		// 创建用户消息提醒
		for _, u := range param.Users {
			user, err := ds.GetUserByUsername(u)
			if err != nil || user.ID == userID {
				continue
			}

			// 创建消息提醒
			// TODO: 优化消息提醒处理机制
			go ds.CreateMessage(&model.Message{
				SenderUserID:   userID,
				ReceiverUserID: user.ID,
				Type:           model.MsgTypePost,
				Brief:          "在新发布的泡泡动态中@了你",
				PostID:         post.ID,
			})
		}
	}

	// 推送Search
	PushPostToSearch(post)

	return post, nil
}

func DeletePost(user *model.User, id int64) *errcode.Error {
	if user == nil {
		return errcode.NoPermission
	}

	post, err := ds.GetPostByID(id)
	if err != nil {
		return errcode.GetPostFailed
	}
	if post.UserID != user.ID && !user.IsAdmin {
		return errcode.NoPermission
	}

	mediaContents, err := ds.DeletePost(post)
	if err != nil {
		logrus.Errorf("service.DeletePost delete post failed: %s", err)
		return errcode.DeletePostFailed
	}

	// 删除推文的媒体内容
	deleteOssObjects(mediaContents)

	// 删除索引
	DeleteSearchPost(post)

	return nil
}

// deleteOssObjects 删除推文的媒体内容, 宽松处理错误(就是不处理), 后续完善
func deleteOssObjects(mediaContents []string) {
	mediaContentsSize := len(mediaContents)
	if mediaContentsSize > 1 {
		objectKeys := make([]string, 0, mediaContentsSize)
		for _, cUrl := range mediaContents {
			objectKeys = append(objectKeys, oss.ObjectKey(cUrl))
		}
		// TODO: 优化处理尽量使用channel传递objectKeys使用可控数量的Goroutine集中处理object删除动作，后续完善
		go oss.DeleteObjects(objectKeys)
	} else if mediaContentsSize == 1 {
		oss.DeleteObject(oss.ObjectKey(mediaContents[0]))
	}
}

func LockPost(id int64) error {
	post, _ := ds.GetPostByID(id)

	err := ds.LockPost(post)

	if err != nil {
		return err
	}

	return nil
}

func StickPost(id int64) error {
	post, _ := ds.GetPostByID(id)

	err := ds.StickPost(post)

	if err != nil {
		return err
	}

	return nil
}

func VisiblePost(user *model.User, postId int64, visibility model.PostVisibleT) *errcode.Error {
	if visibility >= model.PostVisitInvalid {
		return errcode.InvalidParams
	}

	post, err := ds.GetPostByID(postId)
	if err != nil {
		return errcode.GetPostFailed
	}

	if err := checkPermision(user, post.UserID); err != nil {
		return err
	}

	if err = ds.VisiblePost(post, visibility); err != nil {
		logrus.Warnf("update post failure: %v", err)
		return errcode.VisblePostFailed
	}

	// 推送Search
	post.Visibility = visibility
	PushPostToSearch(post)

	return nil
}

func GetPostStar(postID, userID int64) (*model.PostStar, error) {
	return ds.GetUserPostStar(postID, userID)
}

func CreatePostStar(postID, userID int64) (*model.PostStar, error) {
	// 加载Post
	post, err := ds.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	// 私密post不可操作
	if post.Visibility == model.PostVisitPrivate {
		return nil, errors.New("no permision")
	}

	star, err := ds.CreatePostStar(postID, userID)
	if err != nil {
		return nil, err
	}

	// 更新Post点赞数
	post.UpvoteCount++
	ds.UpdatePost(post)

	// 更新索引
	PushPostToSearch(post)

	return star, nil
}

func DeletePostStar(star *model.PostStar) error {
	err := ds.DeletePostStar(star)
	if err != nil {
		return err
	}
	// 加载Post
	post, err := ds.GetPostByID(star.PostID)
	if err != nil {
		return err
	}

	// 私密post不可操作
	if post.Visibility == model.PostVisitPrivate {
		return errors.New("no permision")
	}

	// 更新Post点赞数
	post.UpvoteCount--
	ds.UpdatePost(post)

	// 更新索引
	PushPostToSearch(post)

	return nil
}

func GetPostCollection(postID, userID int64) (*model.PostCollection, error) {
	return ds.GetUserPostCollection(postID, userID)
}

func CreatePostCollection(postID, userID int64) (*model.PostCollection, error) {
	// 加载Post
	post, err := ds.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	// 私密post不可操作
	if post.Visibility == model.PostVisitPrivate {
		return nil, errors.New("no permision")
	}

	collection, err := ds.CreatePostCollection(postID, userID)
	if err != nil {
		return nil, err
	}

	// 更新Post点赞数
	post.CollectionCount++
	ds.UpdatePost(post)

	// 更新索引
	PushPostToSearch(post)

	return collection, nil
}

func DeletePostCollection(collection *model.PostCollection) error {
	err := ds.DeletePostCollection(collection)
	if err != nil {
		return err
	}
	// 加载Post
	post, err := ds.GetPostByID(collection.PostID)
	if err != nil {
		return err
	}

	// 私密post不可操作
	if post.Visibility == model.PostVisitPrivate {
		return errors.New("no permision")
	}

	// 更新Post点赞数
	post.CollectionCount--
	ds.UpdatePost(post)

	// 更新索引
	PushPostToSearch(post)

	return nil
}

func GetPost(id int64) (*model.PostFormated, error) {
	post, err := ds.GetPostByID(id)

	if err != nil {
		return nil, err
	}

	postContents, err := ds.GetPostContentsByIDs([]int64{post.ID})
	if err != nil {
		return nil, err
	}

	users, err := ds.GetUsersByIDs([]int64{post.UserID})
	if err != nil {
		return nil, err
	}

	// 数据整合
	postFormated := post.Format()
	for _, user := range users {
		postFormated.User = user.Format()
	}
	for _, content := range postContents {
		if content.PostID == post.ID {
			postFormated.Contents = append(postFormated.Contents, content.Format())
		}
	}
	return postFormated, nil
}

func GetPostContentByID(id int64) (*model.PostContent, error) {
	return ds.GetPostContentByID(id)
}

func GetIndexPosts(user *model.User, offset int, limit int) (*rest.IndexTweetsResp, error) {
	return ds.IndexPosts(user, offset, limit)
}

func GetPostList(req *PostListReq) ([]*model.PostFormated, error) {
	posts, err := ds.GetPosts(req.Conditions, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	return ds.MergePosts(posts)
}

func GetPostCount(conditions *model.ConditionsT) (int64, error) {
	return ds.GetPostCount(conditions)
}

func GetPostListFromSearch(user *model.User, q *core.QueryReq, offset, limit int) ([]*model.PostFormated, int64, error) {
	resp, err := ts.Search(user, q, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	posts, err := ds.RevampPosts(resp.Items)
	if err != nil {
		return nil, 0, err
	}
	return posts, resp.Total, nil
}

func GetPostListFromSearchByQuery(user *model.User, query string, offset, limit int) ([]*model.PostFormated, int64, error) {
	q := &core.QueryReq{
		Query: query,
		Type:  "search",
	}
	return GetPostListFromSearch(user, q, offset, limit)
}

func PushPostToSearch(post *model.Post) {
	postFormated := post.Format()
	postFormated.User = &model.UserFormated{
		ID: post.UserID,
	}
	contents, _ := ds.GetPostContentsByIDs([]int64{post.ID})
	for _, content := range contents {
		postFormated.Contents = append(postFormated.Contents, content.Format())
	}

	contentFormated := ""

	for _, content := range postFormated.Contents {
		if content.Type == model.CONTENT_TYPE_TEXT || content.Type == model.CONTENT_TYPE_TITLE {
			contentFormated = contentFormated + content.Content + "\n"
		}
	}

	tagMaps := map[string]int8{}
	for _, tag := range strings.Split(post.Tags, ",") {
		tagMaps[tag] = 1
	}

	data := core.DocItems{{
		"id":                post.ID,
		"user_id":           post.UserID,
		"comment_count":     post.CommentCount,
		"collection_count":  post.CollectionCount,
		"upvote_count":      post.UpvoteCount,
		"visibility":        post.Visibility,
		"is_top":            post.IsTop,
		"is_essence":        post.IsEssence,
		"content":           contentFormated,
		"tags":              tagMaps,
		"ip_loc":            post.IPLoc,
		"latest_replied_on": post.LatestRepliedOn,
		"attachment_price":  post.AttachmentPrice,
		"created_on":        post.CreatedOn,
		"modified_on":       post.ModifiedOn,
	}}

	ts.AddDocuments(data, fmt.Sprintf("%d", post.ID))
}

func DeleteSearchPost(post *model.Post) error {
	return ts.DeleteDocuments([]string{fmt.Sprintf("%d", post.ID)})
}

func PushPostsToSearch(c *gin.Context) {
	if ok, _ := conf.Redis.SetNX(c, "JOB_PUSH_TO_SEARCH", 1, time.Hour).Result(); ok {
		splitNum := 1000
		totalRows, _ := GetPostCount(&model.ConditionsT{
			"visibility IN ?": []model.PostVisibleT{model.PostVisitPublic, model.PostVisitFriend},
		})

		pages := math.Ceil(float64(totalRows) / float64(splitNum))
		nums := int(pages)

		for i := 0; i < nums; i++ {
			posts, _ := GetPostList(&PostListReq{
				Conditions: &model.ConditionsT{},
				Offset:     i * splitNum,
				Limit:      splitNum,
			})

			for _, post := range posts {
				contentFormated := ""

				for _, content := range post.Contents {
					if content.Type == model.CONTENT_TYPE_TEXT || content.Type == model.CONTENT_TYPE_TITLE {
						contentFormated = contentFormated + content.Content + "\n"
					}
				}

				docs := core.DocItems{{
					"id":                post.ID,
					"user_id":           post.User.ID,
					"comment_count":     post.CommentCount,
					"collection_count":  post.CollectionCount,
					"upvote_count":      post.UpvoteCount,
					"visibility":        post.Visibility,
					"is_top":            post.IsTop,
					"is_essence":        post.IsEssence,
					"content":           contentFormated,
					"tags":              post.Tags,
					"ip_loc":            post.IPLoc,
					"latest_replied_on": post.LatestRepliedOn,
					"attachment_price":  post.AttachmentPrice,
					"created_on":        post.CreatedOn,
					"modified_on":       post.ModifiedOn,
				}}
				ts.AddDocuments(docs, fmt.Sprintf("%d", post.ID))
			}
		}

		conf.Redis.Del(c, "JOB_PUSH_TO_SEARCH")
	}
}

func GetPostTags(param *PostTagsReq) ([]*model.TagFormated, error) {
	num := param.Num
	if num > conf.AppSetting.MaxPageSize {
		num = conf.AppSetting.MaxPageSize
	}

	conditions := &model.ConditionsT{}
	if param.Type == TagTypeHot {
		// 热门标签
		conditions = &model.ConditionsT{
			"ORDER": "quote_num DESC",
		}
	}
	if param.Type == TagTypeNew {
		// 热门标签
		conditions = &model.ConditionsT{
			"ORDER": "id DESC",
		}
	}

	tags, err := ds.GetTags(conditions, 0, num)
	if err != nil {
		return nil, err
	}

	// 获取创建者User IDs
	userIds := []int64{}
	for _, tag := range tags {
		userIds = append(userIds, tag.UserID)
	}

	users, _ := ds.GetUsersByIDs(userIds)

	tagsFormated := []*model.TagFormated{}
	for _, tag := range tags {
		tagFormated := tag.Format()
		for _, user := range users {
			if user.ID == tagFormated.UserID {
				tagFormated.User = user.Format()
			}
		}
		tagsFormated = append(tagsFormated, tagFormated)
	}

	return tagsFormated, nil
}

func CheckPostAttachmentIsPaid(postID, userID int64) bool {
	bill, err := ds.GetPostAttatchmentBill(postID, userID)

	return err == nil && bill.Model != nil && bill.ID > 0
}
