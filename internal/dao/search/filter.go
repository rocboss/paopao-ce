// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package search

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/types"
)

type tweetSearchFilter struct {
	ams core.AuthorizationManageService
}

func (s *tweetSearchFilter) filterResp(user *core.User, resp *core.QueryResp) {
	// 管理员不过滤
	if user != nil && user.IsAdmin {
		return
	}

	var item *core.PostFormated
	items := resp.Items
	latestIndex := len(items) - 1
	if user == nil {
		for i := 0; i <= latestIndex; i++ {
			item = items[i]
			if item.Visibility != core.PostVisitPublic {
				items[i] = items[latestIndex]
				items = items[:latestIndex]
				resp.Total--
				latestIndex--
				i--
			}
		}
	} else {
		var cutFriend, cutPrivate bool
		friendFilter := s.ams.BeFriendFilter(user.ID)
		friendFilter[user.ID] = types.Empty{}
		for i := 0; i <= latestIndex; i++ {
			item = items[i]
			cutFriend = (item.Visibility == core.PostVisitFriend && !friendFilter.IsFriend(item.UserID))
			cutPrivate = (item.Visibility == core.PostVisitPrivate && user.ID != item.UserID)
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
