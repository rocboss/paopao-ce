package jinzhu

import (
	"fmt"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"strings"
)

type RankService struct {
	db *gorm.DB
}

func NewRankService(db *gorm.DB) *RankService {
	return &RankService{db: db}
}

type UserPosts struct {
	UserId int64
	PostId []int64
}

type UserScore struct {
	//用户id
	UserId int64
	//用户名
	Username string
	//用户头像
	Avatar string
	//帖子数量&得分，一个帖子是1分
	PostCount int64
	//图像得分，一个图像加2分
	ImageCount int64
	//视频得分，一个视频加2分
	VideoCount int64
	//点赞总数
	LikeCount int64
	//评论总数
	CommentCount int64
	//分享码得分
	ShareCodeCount int64
	//总分
	Score int64
}

// GetHighQualityRanking 优质榜单函数
func (s RankService) GetHighQualityRanking() ([]*core.GetHighQualityRankingResp, error) {
	//获取p_post表的所有数据
	var posts []*ms.Post
	err := s.db.Table("p_post").Find(&posts).Error
	//一个用户对应一个总分，建立一个map
	var userScores = make(map[int64]*UserScore)
	//判断posts是否为空
	if len(posts) == 0 {
		fmt.Print("posts is empty")
		return nil, nil
	}
	for _, post := range posts {
		//根据user_id获取用户信息，其中包括用户名，用户头像
		var user ms.User
		//如果userScores中没有该用户的信息，则新建一个用户信息
		if _, ok := userScores[post.UserID]; !ok {
			userScores[post.UserID] = &UserScore{}
			err = s.db.Table("p_user").Select("id,username,avatar").Where("id = ?", post.UserID).Find(&user).Error
			if err != nil {
				fmt.Print("get user info error")
				return nil, err
			}
			userScores[post.UserID].Username = user.Username
			userScores[post.UserID].Avatar = user.Avatar
		}

		var userPostContent []*ms.PostContent
		userScores[post.UserID].LikeCount += post.UpvoteCount
		userScores[post.UserID].CommentCount += post.CommentCount
		//根据post_id查询p_post_content表，获取type
		err = s.db.Table("p_post_content").Select("post_id,type,content").Where("post_id = ?", post.ID).Find(&userPostContent).Error
		if err != nil {
			return nil, err
		}
		//userPostCount[post.UserId] = userPostCount[post.UserId] + 1
		//计算每个用户发帖是否存在正文、图像、视频，即查询p_post_content表，若帖子id对应的type存在3或4，则存在图像或视频
		for _, postContent := range userPostContent {
			//有正文内容的帖子加1分
			if postContent.Type == 2 {
				userScores[post.UserID].PostCount = userScores[post.UserID].PostCount + 1
				//判断正文字符串中是否包含其用户名，若存在，则加5分
				var userNameStr = userScores[post.UserID].Username + "-"
				if strings.Index(postContent.Content, userNameStr) != -1 {
					userScores[post.UserID].ShareCodeCount = userScores[post.UserID].ShareCodeCount + 5
				}
			}
			if postContent.Type == 3 {
				//总分加2
				userScores[post.UserID].ImageCount = userScores[post.UserID].ImageCount + 2
			} else if postContent.Type == 4 {
				//总分加2
				userScores[post.UserID].VideoCount = userScores[post.UserID].VideoCount + 2
			}
		}
	}
	//计算每个用户的总分
	for _, userScore := range userScores {
		userScore.Score = (userScore.PostCount + userScore.ImageCount + userScore.VideoCount + userScore.ShareCodeCount) * ((userScore.LikeCount*3 + userScore.CommentCount) / 10)
	}

	// 对userScore排序
	var userScoreSlice []*UserScore
	for _, userScore := range userScores {
		userScoreSlice = append(userScoreSlice, userScore)
	}
	sort.Slice(userScoreSlice, func(i, j int) bool {
		return userScoreSlice[i].Score > userScoreSlice[j].Score
	})
	var rank []*core.GetHighQualityRankingResp
	var i = 0
	//将userScoreSlice中的数据放到rank中
	for _, userScore := range userScoreSlice {
		if i >= 5 {
			break
		}
		i++
		fmt.Print(userScore.Username + " " + strconv.FormatInt(userScore.CommentCount, 10) +
			" " + strconv.FormatInt(userScore.LikeCount, 10) +
			" " + strconv.FormatInt(userScore.ImageCount, 10) +
			" " + strconv.FormatInt(userScore.PostCount, 10) +
			" " + strconv.FormatInt(userScore.VideoCount, 10) +
			" " + strconv.FormatInt(userScore.ShareCodeCount, 10) +
			" " + strconv.FormatInt(userScore.Score, 10) +
			"\n")
		rank = append(rank, &core.GetHighQualityRankingResp{
			UserName:           userScore.Username,
			Avatar:             userScore.Avatar,
			PostCount:          userScore.PostCount,
			LikesCount:         userScore.LikeCount,
			ComprehensiveScore: userScore.Score,
		})
	}
	return rank, nil
}
