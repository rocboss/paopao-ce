package service

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/util"
	"github.com/rocboss/paopao-ce/pkg/zinc"
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
		if err := attachmentChecker.Check(p.Content); err != nil {
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

func CreatePost(c *gin.Context, userID int64, param PostCreationReq) (*model.Post, error) {
	ip := c.ClientIP()

	post := &model.Post{
		UserID:          userID,
		Tags:            strings.Join(param.Tags, ","),
		IP:              ip,
		IPLoc:           util.GetIPLoc(ip),
		AttachmentPrice: param.AttachmentPrice,
	}
	post, err := ds.CreatePost(post)
	if err != nil {
		return nil, err
	}

	// 创建标签
	for _, t := range param.Tags {
		tag := &model.Tag{
			UserID: userID,
			Tag:    t,
		}
		ds.CreateTag(tag)
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
		ds.CreatePostContent(postContent)
	}

	// 推送Search
	go PushPostToSearch(post)

	// 创建用户消息提醒
	for _, u := range param.Users {
		user, err := ds.GetUserByUsername(u)
		if err != nil || user.ID == userID {
			continue
		}

		// 创建消息提醒
		go ds.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: user.ID,
			Type:           model.MESSAGE_POST,
			Brief:          "在新发布的泡泡动态中@了你",
			PostID:         post.ID,
		})
	}

	return post, nil
}

func DeletePost(id int64) error {
	post, _ := ds.GetPostByID(id)

	// tag删除
	tags := strings.Split(post.Tags, ",")
	for _, t := range tags {

		tag := &model.Tag{
			Tag: t,
		}
		ds.DeleteTag(tag)
	}

	err := ds.DeletePost(post)

	if err != nil {
		return err
	}

	// 删除索引
	go DeleteSearchPost(post)

	return nil
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

func GetPostStar(postID, userID int64) (*model.PostStar, error) {
	return ds.GetUserPostStar(postID, userID)
}

func CreatePostStar(postID, userID int64) (*model.PostStar, error) {
	// 加载Post
	post, err := ds.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	star, err := ds.CreatePostStar(postID, userID)
	if err != nil {
		return nil, err
	}

	// 更新Post点赞数
	post.UpvoteCount++
	ds.UpdatePost(post)

	// 更新索引
	go PushPostToSearch(post)

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
	// 更新Post点赞数
	post.UpvoteCount--
	ds.UpdatePost(post)

	// 更新索引
	go PushPostToSearch(post)

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
	collection, err := ds.CreatePostCollection(postID, userID)
	if err != nil {
		return nil, err
	}

	// 更新Post点赞数
	post.CollectionCount++
	ds.UpdatePost(post)

	// 更新索引
	go PushPostToSearch(post)

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
	// 更新Post点赞数
	post.CollectionCount--
	ds.UpdatePost(post)

	// 更新索引
	go PushPostToSearch(post)

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

func GetPostList(req *PostListReq) ([]*model.PostFormated, error) {
	posts, err := ds.GetPosts(req.Conditions, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	return FormatPosts(posts)
}

func FormatPosts(posts []*model.Post) ([]*model.PostFormated, error) {
	postIds := []int64{}
	userIds := []int64{}
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}

	postContents, err := ds.GetPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := ds.GetUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}

	// 数据整合
	postsFormated := []*model.PostFormated{}
	for _, post := range posts {
		postFormated := post.Format()

		for _, user := range users {
			if user.ID == postFormated.UserID {
				postFormated.User = user.Format()
			}
		}
		for _, content := range postContents {
			if content.PostID == post.ID {
				postFormated.Contents = append(postFormated.Contents, content.Format())
			}
		}

		postsFormated = append(postsFormated, postFormated)
	}

	return postsFormated, nil
}

func GetPostCount(conditions *model.ConditionsT) (int64, error) {
	return ds.GetPostCount(conditions)
}

func GetPostListFromSearch(q *core.QueryT, offset, limit int) ([]*model.PostFormated, int64, error) {
	queryResult, err := ds.QueryAll(q, conf.ZincSetting.Index, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	posts, err := FormatZincPost(queryResult)
	if err != nil {
		return nil, 0, err
	}

	return posts, queryResult.Hits.Total.Value, nil
}

func GetPostListFromSearchByQuery(query string, offset, limit int) ([]*model.PostFormated, int64, error) {
	queryResult, err := ds.QuerySearch(conf.ZincSetting.Index, query, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	posts, err := FormatZincPost(queryResult)
	if err != nil {
		return nil, 0, err
	}

	return posts, queryResult.Hits.Total.Value, nil
}

func PushPostToSearch(post *model.Post) {
	indexName := conf.ZincSetting.Index

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

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"index": map[string]interface{}{
			"_index": indexName,
			"_id":    fmt.Sprintf("%d", post.ID),
		},
	}, map[string]interface{}{
		"id":                post.ID,
		"user_id":           post.UserID,
		"comment_count":     post.CommentCount,
		"collection_count":  post.CollectionCount,
		"upvote_count":      post.UpvoteCount,
		"is_top":            post.IsTop,
		"is_essence":        post.IsEssence,
		"content":           contentFormated,
		"tags":              tagMaps,
		"ip_loc":            post.IPLoc,
		"latest_replied_on": post.LatestRepliedOn,
		"attachment_price":  post.AttachmentPrice,
		"created_on":        post.CreatedOn,
		"modified_on":       post.ModifiedOn,
	})

	ds.BulkPushDoc(data)
}

func DeleteSearchPost(post *model.Post) error {
	indexName := conf.ZincSetting.Index

	return ds.DelDoc(indexName, fmt.Sprintf("%d", post.ID))
}

func PushPostsToSearch(c *gin.Context) {
	if ok, _ := conf.Redis.SetNX(c, "JOB_PUSH_TO_SEARCH", 1, time.Hour).Result(); ok {
		splitNum := 1000
		totalRows, _ := GetPostCount(&model.ConditionsT{})

		pages := math.Ceil(float64(totalRows) / float64(splitNum))
		nums := int(pages)

		indexName := conf.ZincSetting.Index
		// 创建索引
		ds.CreateSearchIndex(indexName)

		for i := 0; i < nums; i++ {
			data := []map[string]interface{}{}

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

				data = append(data, map[string]interface{}{
					"index": map[string]interface{}{
						"_index": indexName,
						"_id":    fmt.Sprintf("%d", post.ID),
					},
				}, map[string]interface{}{
					"id":                post.ID,
					"user_id":           post.User.ID,
					"comment_count":     post.CommentCount,
					"collection_count":  post.CollectionCount,
					"upvote_count":      post.UpvoteCount,
					"is_top":            post.IsTop,
					"is_essence":        post.IsEssence,
					"content":           contentFormated,
					"tags":              post.Tags,
					"ip_loc":            post.IPLoc,
					"latest_replied_on": post.LatestRepliedOn,
					"attachment_price":  post.AttachmentPrice,
					"created_on":        post.CreatedOn,
					"modified_on":       post.ModifiedOn,
				})
			}

			if len(data) > 0 {
				ds.BulkPushDoc(data)
			}
		}

		conf.Redis.Del(c, "JOB_PUSH_TO_SEARCH")
	}
}

func FormatZincPost(queryResult *zinc.QueryResultT) ([]*model.PostFormated, error) {
	posts := []*model.PostFormated{}
	for _, hit := range queryResult.Hits.Hits {
		item := &model.PostFormated{}

		raw, _ := json.Marshal(hit.Source)
		err := json.Unmarshal(raw, item)
		if err == nil {
			posts = append(posts, item)
		}
	}

	postIds := []int64{}
	userIds := []int64{}
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}
	postContents, err := ds.GetPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := ds.GetUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}

	// 数据整合
	for _, post := range posts {
		for _, user := range users {
			if user.ID == post.UserID {
				post.User = user.Format()
			}
		}
		if post.Contents == nil {
			post.Contents = []*model.PostContentFormated{}
		}
		for _, content := range postContents {
			if content.PostID == post.ID {
				post.Contents = append(post.Contents, content.Format())
			}
		}
	}

	return posts, nil
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
