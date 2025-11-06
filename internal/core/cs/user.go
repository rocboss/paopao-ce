// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

// UserInfo 用户基本信息
type UserInfo struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
	Status    int    `json:"status"`
	Avatar    string `json:"avatar"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedOn int64  `json:"created_on"`
}

type UserProfile struct {
	ID          int64  `json:"id" db:"id"`
	Nickname    string `json:"nickname"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Status      int    `json:"status"`
	Avatar      string `json:"avatar"`
	Balance     int64  `json:"balance"`
	IsAdmin     bool   `json:"is_admin"`
	CreatedOn   int64  `json:"created_on"`
	TweetsCount int    `json:"tweets_count"`
	Experience  int    `json:"experience"`
	Level       int    `json:"level"`
}

