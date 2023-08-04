// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// nolint
var (
	ErrUsernameHasExisted      = xerror.NewError(20001, "用户名已存在")
	ErrUsernameLengthLimit     = xerror.NewError(20002, "用户名长度3~12")
	ErrUsernameCharLimit       = xerror.NewError(20003, "用户名只能包含字母、数字")
	ErrPasswordLengthLimit     = xerror.NewError(20004, "密码长度6~16")
	ErrUserRegisterFailed      = xerror.NewError(20005, "用户注册失败")
	ErrUserHasBeenBanned       = xerror.NewError(20006, "该账户已被封停")
	ErrNoPermission            = xerror.NewError(20007, "无权限执行该请求")
	ErrUserHasBindOTP          = xerror.NewError(20008, "当前用户已绑定二次验证")
	ErrUserOTPInvalid          = xerror.NewError(20009, "二次验证码验证失败")
	ErrUserNoBindOTP           = xerror.NewError(20010, "当前用户未绑定二次验证")
	ErrErrorOldPassword        = xerror.NewError(20011, "当前用户密码验证失败")
	ErrErrorCaptchaPassword    = xerror.NewError(20012, "图形验证码验证失败")
	ErrAccountNoPhoneBind      = xerror.NewError(20013, "拒绝操作: 账户未绑定手机号")
	ErrTooManyLoginError       = xerror.NewError(20014, "登录失败次数过多，请稍后再试")
	ErrGetPhoneCaptchaError    = xerror.NewError(20015, "短信验证码获取失败")
	ErrTooManyPhoneCaptchaSend = xerror.NewError(20016, "短信验证码获取次数已达今日上限")
	ErrExistedUserPhone        = xerror.NewError(20017, "该手机号已被绑定")
	ErrErrorPhoneCaptcha       = xerror.NewError(20018, "手机验证码不正确")
	ErrMaxPhoneCaptchaUseTimes = xerror.NewError(20019, "手机验证码已达最大使用次数")
	ErrNicknameLengthLimit     = xerror.NewError(20020, "昵称长度2~12")
	ErrNoExistUsername         = xerror.NewError(20021, "用户不存在")
	ErrNoAdminPermission       = xerror.NewError(20022, "无管理权限")
	ErrDisallowUserRegister    = xerror.NewError(20023, "系统不允许注册用户")

	ErrGetPostsFailed          = xerror.NewError(30001, "获取动态列表失败")
	ErrCreatePostFailed        = xerror.NewError(30002, "动态发布失败")
	ErrGetPostFailed           = xerror.NewError(30003, "获取动态详情失败")
	ErrDeletePostFailed        = xerror.NewError(30004, "动态删除失败")
	ErrLockPostFailed          = xerror.NewError(30005, "动态锁定失败")
	ErrGetPostTagsFailed       = xerror.NewError(30006, "获取话题列表失败")
	ErrInvalidDownloadReq      = xerror.NewError(30007, "附件下载请求不合法")
	ErrDownloadReqError        = xerror.NewError(30008, "附件下载请求失败")
	ErrInsuffientDownloadMoney = xerror.NewError(30009, "附件下载失败:账户资金不足")
	ErrDownloadExecFail        = xerror.NewError(30010, "附件下载失败:扣费失败")
	ErrStickPostFailed         = xerror.NewError(30011, "动态置顶失败")
	ErrVisblePostFailed        = xerror.NewError(30012, "更新可见性失败")

	ErrGetCommentsFailed   = xerror.NewError(40001, "获取评论列表失败")
	ErrCreateCommentFailed = xerror.NewError(40002, "评论发布失败")
	ErrGetCommentFailed    = xerror.NewError(40003, "获取评论详情失败")
	ErrDeleteCommentFailed = xerror.NewError(40004, "评论删除失败")
	ErrCreateReplyFailed   = xerror.NewError(40005, "评论回复失败")
	ErrGetReplyFailed      = xerror.NewError(40006, "获取评论详情失败")
	ErrMaxCommentCount     = xerror.NewError(40007, "评论数已达最大限制")
	ErrGetCommentThumbs    = xerror.NewError(40008, "获取评论点赞信息失败")

	ErrGetMessagesFailed = xerror.NewError(50001, "获取消息列表失败")
	ErrReadMessageFailed = xerror.NewError(50002, "标记消息已读失败")
	ErrSendWhisperFailed = xerror.NewError(50003, "私信发送失败")
	ErrNoWhisperToSelf   = xerror.NewError(50004, "不允许给自己发送私信")
	ErrTooManyWhisperNum = xerror.NewError(50005, "今日私信次数已达上限")

	ErrGetCollectionsFailed = xerror.NewError(60001, "获取收藏列表失败")
	ErrGetStarsFailed       = xerror.NewError(60002, "获取点赞列表失败")

	ErrRechargeReqFail       = xerror.NewError(70001, "充值请求失败")
	ErrRechargeNotifyError   = xerror.NewError(70002, "充值回调失败")
	ErrGetRechargeFailed     = xerror.NewError(70003, "充值详情获取失败")
	ErrUserWalletBillsFailed = xerror.NewError(70004, "用户钱包账单获取失败")

	ErrNoRequestingFriendToSelf   = xerror.NewError(80001, "不允许添加自己为好友")
	ErrNotExistFriendId           = xerror.NewError(80002, "好友id不存在")
	ErrSendRequestingFriendFailed = xerror.NewError(80003, "申请添加朋友请求发送失败")
	ErrAddFriendFailed            = xerror.NewError(80004, "添加好友失败")
	ErrRejectFriendFailed         = xerror.NewError(80005, "拒绝好友失败")
	ErrDeleteFriendFailed         = xerror.NewError(80006, "删除好友失败")
	ErrGetContactsFailed          = xerror.NewError(80007, "获取联系人列表失败")
	ErrNoActionToSelf             = xerror.NewError(80008, "不允许对自己操作")

	ErrFollowTopicFailed      = xerror.NewError(90001, "关注话题失败")
	ErrUnfollowTopicFailed    = xerror.NewError(90002, "取消关注话题失败")
	ErrStickTopicFailed       = xerror.NewError(90003, "更行话题置顶状态失败")
	ErrThumbsUpTweetComment   = xerror.NewError(90101, "评论点赞失败")
	ErrThumbsDownTweetComment = xerror.NewError(90102, "评论点踩失败")
	ErrThumbsUpTweetReply     = xerror.NewError(90103, "评论回复点赞失败")
	ErrThumbsDownTweetReply   = xerror.NewError(90104, "评论回复点踩失败")

	ErrFileUploadFailed = xerror.NewError(10200, "文件上传失败")
	ErrFileInvalidExt   = xerror.NewError(10201, "文件类型不合法")
	ErrFileInvalidSize  = xerror.NewError(10202, "文件大小超限")

	ErrGetUserKeysFailed          = xerror.NewError(11001, "获取用户Share Key失败")
	ErrDsNil                      = xerror.NewError(11002, "数据源为空")
	ErrUserNameEmpty              = xerror.NewError(11003, "用户名为空")
	ErrDeleteUserKeyFailed        = xerror.NewError(11004, "删除用户Share Key失败")
	ErrGetUserShareKeyCountFailed = xerror.NewError(11005, "获取用户Share Key数量失败")
)
