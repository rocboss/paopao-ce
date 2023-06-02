// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ms

type (
	ContactItem struct {
		UserId   int64  `json:"user_id"`
		UserName string `json:"username"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		Phone    string `json:"phone"`
	}

	ContactList struct {
		Contacts []ContactItem `json:"contacts"`
		Total    int64         `json:"total"`
	}
)
