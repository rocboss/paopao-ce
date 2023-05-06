// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// nolint
var (
	_errUsernameHasExisted      = xerror.NewError(20001, "用户名已存在")
	_errUsernameLengthLimit     = xerror.NewError(20002, "用户名长度3~12")
	_errUsernameCharLimit       = xerror.NewError(20003, "用户名只能包含字母、数字")
	_errPasswordLengthLimit     = xerror.NewError(20004, "密码长度6~16")
	_errUserRegisterFailed      = xerror.NewError(20005, "用户注册失败")
	_errUserHasBeenBanned       = xerror.NewError(20006, "该账户已被封停")
	_errNoPermission            = xerror.NewError(20007, "无权限执行该请求")
	_errUserHasBindOTP          = xerror.NewError(20008, "当前用户已绑定二次验证")
	_errUserOTPInvalid          = xerror.NewError(20009, "二次验证码验证失败")
	_errUserNoBindOTP           = xerror.NewError(20010, "当前用户未绑定二次验证")
	_errErrorOldPassword        = xerror.NewError(20011, "当前用户密码验证失败")
	_errErrorCaptchaPassword    = xerror.NewError(20012, "图形验证码验证失败")
	_errAccountNoPhoneBind      = xerror.NewError(20013, "拒绝操作: 账户未绑定手机号")
	_errTooManyLoginError       = xerror.NewError(20014, "登录失败次数过多，请稍后再试")
	_errGetPhoneCaptchaError    = xerror.NewError(20015, "短信验证码获取失败")
	_errTooManyPhoneCaptchaSend = xerror.NewError(20016, "短信验证码获取次数已达今日上限")
	_errExistedUserPhone        = xerror.NewError(20017, "该手机号已被绑定")
	_errErrorPhoneCaptcha       = xerror.NewError(20018, "手机验证码不正确")
	_errMaxPhoneCaptchaUseTimes = xerror.NewError(20019, "手机验证码已达最大使用次数")
	_errNicknameLengthLimit     = xerror.NewError(20020, "昵称长度2~12")
	_errNoExistUsername         = xerror.NewError(20021, "用户不存在")
	_errNoAdminPermission       = xerror.NewError(20022, "无管理权限")
	_errDisallowUserRegister    = xerror.NewError(20023, "系统不允许注册用户")

	_errGetPostsFailed          = xerror.NewError(30001, "获取动态列表失败")
	_errCreatePostFailed        = xerror.NewError(30002, "动态发布失败")
	_errGetPostFailed           = xerror.NewError(30003, "获取动态详情失败")
	_errDeletePostFailed        = xerror.NewError(30004, "动态删除失败")
	_errLockPostFailed          = xerror.NewError(30005, "动态锁定失败")
	_errGetPostTagsFailed       = xerror.NewError(30006, "获取话题列表失败")
	_errInvalidDownloadReq      = xerror.NewError(30007, "附件下载请求不合法")
	_errDownloadReqError        = xerror.NewError(30008, "附件下载请求失败")
	_errInsuffientDownloadMoney = xerror.NewError(30009, "附件下载失败:账户资金不足")
	_errDownloadExecFail        = xerror.NewError(30010, "附件下载失败:扣费失败")
	_errStickPostFailed         = xerror.NewError(30011, "动态置顶失败")
	_errVisblePostFailed        = xerror.NewError(30012, "更新可见性失败")

	_errGetCommentsFailed   = xerror.NewError(40001, "获取评论列表失败")
	_errCreateCommentFailed = xerror.NewError(40002, "评论发布失败")
	_errGetCommentFailed    = xerror.NewError(40003, "获取评论详情失败")
	_errDeleteCommentFailed = xerror.NewError(40004, "评论删除失败")
	_errCreateReplyFailed   = xerror.NewError(40005, "评论回复失败")
	_errGetReplyFailed      = xerror.NewError(40006, "获取评论详情失败")
	_errMaxCommentCount     = xerror.NewError(40007, "评论数已达最大限制")
	_errGetCommentThumbs    = xerror.NewError(40008, "获取评论点赞信息失败")

	_errGetMessagesFailed = xerror.NewError(50001, "获取消息列表失败")
	_errReadMessageFailed = xerror.NewError(50002, "标记消息已读失败")
	_errSendWhisperFailed = xerror.NewError(50003, "私信发送失败")
	_errNoWhisperToSelf   = xerror.NewError(50004, "不允许给自己发送私信")
	_errTooManyWhisperNum = xerror.NewError(50005, "今日私信次数已达上限")

	_errGetCollectionsFailed = xerror.NewError(60001, "获取收藏列表失败")
	_errGetStarsFailed       = xerror.NewError(60002, "获取点赞列表失败")

	_errRechargeReqFail       = xerror.NewError(70001, "充值请求失败")
	_errRechargeNotifyError   = xerror.NewError(70002, "充值回调失败")
	_errGetRechargeFailed     = xerror.NewError(70003, "充值详情获取失败")
	_errUserWalletBillsFailed = xerror.NewError(70004, "用户钱包账单获取失败")

	_errNoRequestingFriendToSelf   = xerror.NewError(80001, "不允许添加自己为好友")
	_errNotExistFriendId           = xerror.NewError(80002, "好友id不存在")
	_errSendRequestingFriendFailed = xerror.NewError(80003, "申请添加朋友请求发送失败")
	_errAddFriendFailed            = xerror.NewError(80004, "添加好友失败")
	_errRejectFriendFailed         = xerror.NewError(80005, "拒绝好友失败")
	_errDeleteFriendFailed         = xerror.NewError(80006, "删除好友失败")
	_errGetContactsFailed          = xerror.NewError(80007, "获取联系人列表失败")
	_errNoActionToSelf             = xerror.NewError(80008, "不允许对自己操作")

	_errFollowTopicFailed      = xerror.NewError(90001, "关注话题失败")
	_errUnfollowTopicFailed    = xerror.NewError(90002, "取消关注话题失败")
	_errStickTopicFailed       = xerror.NewError(90003, "更行话题置顶状态失败")
	_errThumbsUpTweetComment   = xerror.NewError(90101, "评论点赞失败")
	_errThumbsDownTweetComment = xerror.NewError(90102, "评论点踩失败")
	_errThumbsUpTweetReply     = xerror.NewError(90103, "评论回复点赞失败")
	_errThumbsDownTweetReply   = xerror.NewError(90104, "评论回复点踩失败")

	_errFileUploadFailed = xerror.NewError(10200, "文件上传失败")
	_errFileInvalidExt   = xerror.NewError(10201, "文件类型不合法")
	_errFileInvalidSize  = xerror.NewError(10202, "文件大小超限")
)
