package service

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/rocboss/paopao-api/global"
	"github.com/rocboss/paopao-api/internal/dao"
	"github.com/rocboss/paopao-api/internal/model"
	"github.com/rocboss/paopao-api/pkg/util"
	"github.com/rocboss/paopao-api/pkg/zinc"
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

func (svc *Service) CreatePost(userID int64, param PostCreationReq) (*model.Post, error) {
	ip := svc.ctx.ClientIP()

	post := &model.Post{
		UserID:          userID,
		Tags:            strings.Join(param.Tags, ","),
		IP:              ip,
		IPLoc:           util.GetIPLoc(ip),
		AttachmentPrice: param.AttachmentPrice,
	}
	post, err := svc.dao.CreatePost(post)
	if err != nil {
		return nil, err
	}

	// 创建标签
	for _, t := range param.Tags {
		tag := &model.Tag{
			UserID: userID,
			Tag:    t,
		}
		svc.dao.CreateTag(tag)
	}

	for _, item := range param.Contents {

		// 检查附件是否是本站资源
		if item.Type == model.CONTENT_TYPE_IMAGE || item.Type == model.CONTENT_TYPE_VIDEO || item.Type == model.CONTENT_TYPE_ATTACHMENT {
			if strings.Index(item.Content, "https://"+global.AliossSetting.AliossDomain) != 0 {
				continue
			}
		}
		// 检查链接是否合法
		if item.Type == model.CONTENT_TYPE_LINK {
			if strings.Index(item.Content, "http://") != 0 && strings.Index(item.Content, "https://") != 0 {
				continue
			}
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
		svc.dao.CreatePostContent(postContent)
	}

	// 推送Search
	go svc.PushPostToSearch(post)

	// 创建用户消息提醒
	for _, u := range param.Users {
		user, err := svc.dao.GetUserByUsername(u)
		if err != nil || user.ID == userID {
			continue
		}

		// 创建消息提醒
		go svc.dao.CreateMessage(&model.Message{
			SenderUserID:   userID,
			ReceiverUserID: user.ID,
			Type:           model.MESSAGE_POST,
			Breif:          "在新发布的泡泡动态中@了你",
			PostID:         post.ID,
		})
	}

	return post, nil
}

func (svc *Service) DeletePost(id int64) error {
	post, _ := svc.dao.GetPostByID(id)

	// tag删除
	tags := strings.Split(post.Tags, ",")
	for _, t := range tags {

		tag := &model.Tag{
			Tag: t,
		}
		svc.dao.DeleteTag(tag)
	}

	err := svc.dao.DeletePost(post)

	if err != nil {
		return err
	}

	// 删除索引
	go svc.DeleteSearchPost(post)

	return nil
}

func (svc *Service) LockPost(id int64) error {
	post, _ := svc.dao.GetPostByID(id)

	err := svc.dao.LockPost(post)

	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetPostStar(postID, userID int64) (*model.PostStar, error) {
	return svc.dao.GetUserPostStar(postID, userID)
}

func (svc *Service) CreatePostStar(postID, userID int64) (*model.PostStar, error) {
	// 加载Post
	post, err := svc.dao.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	star, err := svc.dao.CreatePostStar(postID, userID)
	if err != nil {
		return nil, err
	}

	// 更新Post点赞数
	post.UpvoteCount++
	svc.dao.UpdatePost(post)

	// 更新索引
	go svc.PushPostToSearch(post)

	return star, nil
}

func (svc *Service) DeletePostStar(star *model.PostStar) error {
	err := svc.dao.DeletePostStar(star)
	if err != nil {
		return err
	}
	// 加载Post
	post, err := svc.dao.GetPostByID(star.PostID)
	if err != nil {
		return err
	}
	// 更新Post点赞数
	post.UpvoteCount--
	svc.dao.UpdatePost(post)

	// 更新索引
	go svc.PushPostToSearch(post)

	return nil
}

func (svc *Service) GetPostCollection(postID, userID int64) (*model.PostCollection, error) {
	return svc.dao.GetUserPostCollection(postID, userID)
}

func (svc *Service) CreatePostCollection(postID, userID int64) (*model.PostCollection, error) {
	// 加载Post
	post, err := svc.dao.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	collection, err := svc.dao.CreatePostCollection(postID, userID)
	if err != nil {
		return nil, err
	}

	// 更新Post点赞数
	post.CollectionCount++
	svc.dao.UpdatePost(post)

	// 更新索引
	go svc.PushPostToSearch(post)

	return collection, nil
}

func (svc *Service) DeletePostCollection(collection *model.PostCollection) error {
	err := svc.dao.DeletePostCollection(collection)
	if err != nil {
		return err
	}
	// 加载Post
	post, err := svc.dao.GetPostByID(collection.PostID)
	if err != nil {
		return err
	}
	// 更新Post点赞数
	post.CollectionCount--
	svc.dao.UpdatePost(post)

	// 更新索引
	go svc.PushPostToSearch(post)

	return nil
}

func (svc *Service) GetPost(id int64) (*model.PostFormated, error) {
	post, err := svc.dao.GetPostByID(id)

	if err != nil {
		return nil, err
	}

	postContents, err := svc.dao.GetPostContentsByIDs([]int64{post.ID})
	if err != nil {
		return nil, err
	}

	users, err := svc.dao.GetUsersByIDs([]int64{post.UserID})
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

func (svc *Service) GetPostContentByID(id int64) (*model.PostContent, error) {
	return svc.dao.GetPostContentByID(id)
}

func (svc *Service) GetPostList(req *PostListReq) ([]*model.PostFormated, error) {
	posts, err := svc.dao.GetPosts(req.Conditions, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	return svc.FormatPosts(posts)
}

func (svc *Service) FormatPosts(posts []*model.Post) ([]*model.PostFormated, error) {
	postIds := []int64{}
	userIds := []int64{}
	for _, post := range posts {
		postIds = append(postIds, post.ID)
		userIds = append(userIds, post.UserID)
	}

	postContents, err := svc.dao.GetPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := svc.dao.GetUsersByIDs(userIds)
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

func (svc *Service) GetPostCount(conditions *model.ConditionsT) (int64, error) {
	return svc.dao.GetPostCount(conditions)
}

func (svc *Service) GetPostListFromSearch(q *dao.QueryT, offset, limit int) ([]*model.PostFormated, int64, error) {
	queryResult, err := svc.dao.QueryAll(q, global.SearchSetting.ZincIndex, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	posts, err := svc.FormatZincPost(queryResult)
	if err != nil {
		return nil, 0, err
	}

	return posts, queryResult.Hits.Total.Value, nil
}

func (svc *Service) GetPostListFromSearchByQuery(query string, offset, limit int) ([]*model.PostFormated, int64, error) {
	queryResult, err := svc.dao.QuerySearch(global.SearchSetting.ZincIndex, query, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	posts, err := svc.FormatZincPost(queryResult)
	if err != nil {
		return nil, 0, err
	}

	return posts, queryResult.Hits.Total.Value, nil
}

func (svc *Service) PushPostToSearch(post *model.Post) {
	indexName := global.SearchSetting.ZincIndex

	postFormated := post.Format()
	postFormated.User = &model.UserFormated{
		ID: post.UserID,
	}
	contents, _ := svc.dao.GetPostContentsByIDs([]int64{post.ID})
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

	svc.dao.BulkPushDoc(data)
}

func (svc *Service) DeleteSearchPost(post *model.Post) error {
	indexName := global.SearchSetting.ZincIndex

	return svc.dao.DelDoc(indexName, fmt.Sprintf("%d", post.ID))
}

func (svc *Service) PushPostsToSearch() {
	if ok, _ := global.Redis.SetNX(svc.ctx, "JOB_PUSH_TO_SEARCH", 1, time.Hour).Result(); ok {
		splitNum := 1000
		totalRows, _ := svc.GetPostCount(&model.ConditionsT{})

		pages := math.Ceil(float64(totalRows) / float64(splitNum))
		nums := int(pages)

		indexName := global.SearchSetting.ZincIndex
		// 创建索引
		svc.dao.CreateSearchIndex(indexName)

		for i := 0; i < nums; i++ {
			data := []map[string]interface{}{}

			posts, _ := svc.GetPostList(&PostListReq{
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
				svc.dao.BulkPushDoc(data)
			}
		}

		global.Redis.Del(svc.ctx, "JOB_PUSH_TO_SEARCH")
	}
}

func (svc *Service) FormatZincPost(queryResult *zinc.QueryResultT) ([]*model.PostFormated, error) {
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
	postContents, err := svc.dao.GetPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := svc.dao.GetUsersByIDs(userIds)
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

func (svc *Service) GetPostTags(param *PostTagsReq) ([]*model.TagFormated, error) {
	num := param.Num
	if num > global.AppSetting.MaxPageSize {
		num = global.AppSetting.MaxPageSize
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

	tags, err := svc.dao.GetTags(conditions, 0, num)
	if err != nil {
		return nil, err
	}

	// 获取创建者User IDs
	userIds := []int64{}
	for _, tag := range tags {
		userIds = append(userIds, tag.UserID)
	}

	users, _ := svc.dao.GetUsersByIDs(userIds)

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

func (svc *Service) CheckPostAttachmentIsPaid(postID, userID int64) bool {
	bill, err := svc.dao.GetPostAttatchmentBill(postID, userID)

	return err == nil && bill.Model != nil && bill.ID > 0
}
