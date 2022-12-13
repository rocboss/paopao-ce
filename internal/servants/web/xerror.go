// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

var (
	errUsernameHasExisted      = xerror.NewError(20001, "用户名已存在")
	errUsernameLengthLimit     = xerror.NewError(20002, "用户名长度3~12")
	errUsernameCharLimit       = xerror.NewError(20003, "用户名只能包含字母、数字")
	errPasswordLengthLimit     = xerror.NewError(20004, "密码长度6~16")
	errUserRegisterFailed      = xerror.NewError(20005, "用户注册失败")
	errUserHasBeenBanned       = xerror.NewError(20006, "该账户已被封停")
	errNoPermission            = xerror.NewError(20007, "无权限执行该请求")
	errUserHasBindOTP          = xerror.NewError(20008, "当前用户已绑定二次验证")
	errUserOTPInvalid          = xerror.NewError(20009, "二次验证码验证失败")
	errUserNoBindOTP           = xerror.NewError(20010, "当前用户未绑定二次验证")
	errErrorOldPassword        = xerror.NewError(20011, "当前用户密码验证失败")
	errErrorCaptchaPassword    = xerror.NewError(20012, "图形验证码验证失败")
	errAccountNoPhoneBind      = xerror.NewError(20013, "拒绝操作: 账户未绑定手机号")
	errTooManyLoginError       = xerror.NewError(20014, "登录失败次数过多，请稍后再试")
	errGetPhoneCaptchaError    = xerror.NewError(20015, "短信验证码获取失败")
	errTooManyPhoneCaptchaSend = xerror.NewError(20016, "短信验证码获取次数已达今日上限")
	errExistedUserPhone        = xerror.NewError(20017, "该手机号已被绑定")
	errErrorPhoneCaptcha       = xerror.NewError(20018, "手机验证码不正确")
	errMaxPhoneCaptchaUseTimes = xerror.NewError(20019, "手机验证码已达最大使用次数")
	errNicknameLengthLimit     = xerror.NewError(20020, "昵称长度2~12")
	errNoExistUsername         = xerror.NewError(20021, "用户不存在")
	errNoAdminPermission       = xerror.NewError(20022, "无管理权限")

	errGetPostsFailed          = xerror.NewError(30001, "获取动态列表失败")
	errCreatePostFailed        = xerror.NewError(30002, "动态发布失败")
	errGetPostFailed           = xerror.NewError(30003, "获取动态详情失败")
	errDeletePostFailed        = xerror.NewError(30004, "动态删除失败")
	errLockPostFailed          = xerror.NewError(30005, "动态锁定失败")
	errGetPostTagsFailed       = xerror.NewError(30006, "获取话题列表失败")
	errInvalidDownloadReq      = xerror.NewError(30007, "附件下载请求不合法")
	errDownloadReqError        = xerror.NewError(30008, "附件下载请求失败")
	errInsuffientDownloadMoney = xerror.NewError(30009, "附件下载失败:账户资金不足")
	errDownloadExecFail        = xerror.NewError(30010, "附件下载失败:扣费失败")
	errStickPostFailed         = xerror.NewError(30011, "动态置顶失败")
	errVisblePostFailed        = xerror.NewError(30012, "更新可见性失败")

	errGetCommentsFailed   = xerror.NewError(40001, "获取评论列表失败")
	errCreateCommentFailed = xerror.NewError(40002, "评论发布失败")
	errGetCommentFailed    = xerror.NewError(40003, "获取评论详情失败")
	errDeleteCommentFailed = xerror.NewError(40004, "评论删除失败")
	errCreateReplyFailed   = xerror.NewError(40005, "评论回复失败")
	errGetReplyFailed      = xerror.NewError(40006, "获取评论详情失败")
	errMaxCommentCount     = xerror.NewError(40007, "评论数已达最大限制")

	errGetMessagesFailed = xerror.NewError(50001, "获取消息列表失败")
	errReadMessageFailed = xerror.NewError(50002, "标记消息已读失败")
	errSendWhisperFailed = xerror.NewError(50003, "私信发送失败")
	errNoWhisperToSelf   = xerror.NewError(50004, "不允许给自己发送私信")
	errTooManyWhisperNum = xerror.NewError(50005, "今日私信次数已达上限")

	errGetCollectionsFailed = xerror.NewError(60001, "获取收藏列表失败")
	errGetStarsFailed       = xerror.NewError(60002, "获取点赞列表失败")

	errRechargeReqFail     = xerror.NewError(70001, "充值请求失败")
	errRechargeNotifyError = xerror.NewError(70002, "充值回调失败")
	errGetRechargeFailed   = xerror.NewError(70003, "充值详情获取失败")

	errNoRequestingFriendToSelf   = xerror.NewError(80001, "不允许添加自己为好友")
	errNotExistFriendId           = xerror.NewError(80002, "好友id不存在")
	errSendRequestingFriendFailed = xerror.NewError(80003, "申请添加朋友请求发送失败")
	errAddFriendFailed            = xerror.NewError(80004, "添加好友失败")
	errRejectFriendFailed         = xerror.NewError(80005, "拒绝好友失败")
	errDeleteFriendFailed         = xerror.NewError(80006, "删除好友失败")
	errGetContactsFailed          = xerror.NewError(80007, "获取联系人列表失败")
	errNoActionToSelf             = xerror.NewError(80008, "不允许对自己操作")
)