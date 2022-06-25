package dao

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
)

func (s *tweetSearchFilter) filterResp(user *model.User, resp *core.QueryResp) {
	// 管理员不过滤
	if user != nil && user.IsAdmin {
		return
	}

	var item *model.PostFormated
	items := resp.Items
	latestIndex := len(items) - 1
	if user == nil {
		for i := 0; i <= latestIndex; i++ {
			item = items[i]
			if item.Visibility != model.PostVisitPublic {
				items[i] = items[latestIndex]
				items = items[:latestIndex]
				resp.Total--
				latestIndex--
				i--
			}
		}
	} else {
		var cutFriend, cutPrivate bool
		friendFilter := s.ams.GetFriendFilter(user.ID)
		for i := 0; i <= latestIndex; i++ {
			item = items[i]
			cutFriend = (item.Visibility == model.PostVisitFriend && !friendFilter.IsFriend(item.UserID))
			cutPrivate = (item.Visibility == model.PostVisitPrivate && user.ID != item.UserID)
			if cutFriend || cutPrivate {
				items[i] = items[latestIndex]
				items = items[:latestIndex]
				resp.Total--
				latestIndex--
				i--
			}
		}
	}

	resp.Items = items
}
