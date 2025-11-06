package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// CheckInActivitity 签到活动
type CheckInActivitity struct {
	Schema `mir:"v1,chain"`

	// UserCheckIn 用户签到
	UserCheckIn func(Post, web.UserCheckInReq) web.UserCheckInResp `mir:"checkin"`

	// GetCheckInRank 获取签到排行榜
	GetCheckInRank func(Get, web.GetCheckInRankReq) web.GetCheckInRankResp `mir:"checkin/rank"`
}
