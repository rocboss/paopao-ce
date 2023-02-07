// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

// TweetBox 推文列表盒子，包含其他一些关于推文列表的信息
type TweetBox struct {
	Tweets TweetList
	Total  int64
}
