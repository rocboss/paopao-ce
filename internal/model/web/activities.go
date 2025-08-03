package web

type UserCheckInResp struct {
	// CheckInCount 签到次数
	CheckInExp int `json:"check_in_exp"`
	// CheckInDays 签到天数
	CheckInDays int `json:"check_in_days"`
}
