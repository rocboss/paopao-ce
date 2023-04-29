// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/version"
)

type TweetDetailReq struct {
	TweetId int64 `form:"id"`
}

type TweetDetailResp core.PostFormated

type GetCaptchaResp struct {
	Id      string `json:"id"`
	Content string `json:"b64s"`
}

type SendCaptchaReq struct {
	Phone        string `json:"phone" form:"phone" binding:"required"`
	ImgCaptcha   string `json:"img_captcha" form:"img_captcha" binding:"required"`
	ImgCaptchaID string `json:"img_captcha_id" form:"img_captcha_id" binding:"required"`
}

type VersionResp struct {
	BuildInfo *version.BuildInfo `json:"build_info"`
}

type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterResp struct {
	UserId   int64  `json:"id"`
	Username string `json:"username"`
}
