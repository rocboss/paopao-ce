package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// UserCheckInReq 用户签到请求
type UserCheckInReq struct {
	UserId int64 `json:"-"` // 用户ID，从JWT中获取
}

// UserCheckInResp 用户签到响应
type UserCheckInResp struct {
	CheckInExp   int `json:"check_in_exp"`   // 签到获得的经验值
	CheckInDays  int `json:"check_in_days"`  // 用户总签到天数
	UserLevel    int `json:"user_level"`     // 用户当前等级
	UserExp      int `json:"user_exp"`       // 用户当前经验值
}

// GetCheckInRankReq 获取签到排行榜请求
type GetCheckInRankReq struct {
	UserId int64 `json:"-"` // 用户ID，从JWT中获取
}

// GetCheckInRankResp 获取签到排行榜响应
type GetCheckInRankResp struct {
	Ranks []CheckInRankItem `json:"ranks"`
}

// CheckInRankItem 签到排行榜项
type CheckInRankItem struct {
	UserID      int64  `json:"user_id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	CheckInExp  int    `json:"check_in_exp"`  // 今日签到获得的经验值
	CheckInTime int64  `json:"check_in_time"` // 签到时间戳
}

// Bind 绑定请求参数
func (r *UserCheckInReq) Bind(c *gin.Context) error {
	userId, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.UserId = userId
	return nil
}

// Bind 绑定请求参数
func (r *GetCheckInRankReq) Bind(c *gin.Context) error {
	userId, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.UserId = userId
	return nil
}
