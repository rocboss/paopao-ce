import { request } from "@/utils/request";

/** 获取验证码 */
export const getCaptcha = (
  params: NetParams.UserGetCaptcha = {}
): Promise<NetReq.UserGetCaptcha> => {
  return request({
    method: "get",
    url: "/v1/captcha",
    params,
  });
};

/**
 * 发送短信验证码
 * @param {Object} data
 * @returns Promise
 */
export const sendCaptcha = (data: any) => {
  return request({
    method: "post",
    url: "/v1/captcha",
    data,
  });
};

/**
 * 发送用户私信
 * @param {Object} data
 * @returns Promise
 */
export const sendUserWhisper = (
  data: NetParams.UserWhisper
): Promise<NetParams.UserWhisper> => {
  return request({
    method: "post",
    url: "/v1/user/whisper",
    data,
  });
};

/**
 * 添加朋友申请
 * @param {Object} data
 * @returns Promise
 */
export const requestingFriend = (
  data: NetParams.RequestingFriend
): Promise<NetReq.RequestingFriend> => {
  return request({
    method: "post",
    url: "/v1/friend/requesting",
    data,
  });
};

/**
 * 添加好友
 * @param {Object} data
 * @returns Promise
 */
export const addFriend = (
  data: NetParams.AddFriend
): Promise<NetReq.AddFriend> => {
  return request({
    method: "post",
    url: "/v1/friend/add",
    data,
  });
};

// 关注 用户
export const followUser = (
  data: NetParams.FollowUserReq
): Promise<NetReq.FollowUserResp> => {
  return request({
    method: "post",
    url: "/v1/user/follow",
    data,
  });
};

// 取消关注 用户
export const unfollowUser = (
  data: NetParams.UnfollowUserReq
): Promise<NetReq.UnfollowUserResp> => {
  return request({
    method: "post",
    url: "/v1/user/unfollow",
    data,
  });
};

/**
 * 获取关注列表
 * @param {Object} params
 * @returns Promise
 */
export const getUserFollows = (
  params: NetParams.GetUserFollows
): Promise<NetReq.GetContacts> => {
  return request({
    method: "get",
    url: "/v1/user/follows",
    params,
  });
};

/**
 * 获取粉丝列表
 * @param {Object} params
 * @returns Promise
 */
export const getUserFollowings = (
  params: NetParams.GetUserFollowings
): Promise<NetReq.GetContacts> => {
  return request({
    method: "get",
    url: "/v1/user/followings",
    params,
  });
};

/**
 * 拒绝朋友申请
 * @param {Object} data
 * @returns Promise
 */
export const rejectFriend = (
  data: NetParams.RejectFriend
): Promise<NetReq.RejectFriend> => {
  return request({
    method: "post",
    url: "/v1/friend/reject",
    data,
  });
};

/**
 * 删除好友
 * @param {Object} data
 * @returns Promise
 */
export const deleteFriend = (
  data: NetParams.DeleteFriend
): Promise<NetReq.DeleteFriend> => {
  return request({
    method: "post",
    url: "/v1/friend/delete",
    data,
  });
};

/**
 * 获取联系人列表
 * @param {Object} data
 * @returns Promise
 */
export const getContacts = (
  data: NetParams.GetContacts
): Promise<NetReq.GetContacts> => {
  return request({
    method: "post",
    url: "/v1/user/contacts",
    data,
  });
};

/**
 * 绑定用户手机
 * @param {Object} data
 * @returns Promise
 */
export const bindUserPhone = (
  data: NetParams.UserBindUserPhone
): Promise<NetParams.UserBindUserPhone> => {
  return request({
    method: "post",
    url: "/v1/user/phone",
    data,
  });
};

/**
 * 激活码激活
 * @param {Object} data
 * @returns Promise
 */
export const activateUser = (
  data: NetParams.UserActivation
): Promise<NetParams.UserActivation> => {
  return request({
    method: "post",
    url: "/v1/user/activate",
    data,
  });
};

/** 更改密码 */
export const changePassword = (
  data: NetParams.UserChangePassword
): Promise<NetReq.UserChangePassword> => {
  return request({
    method: "post",
    url: "/v1/user/password",
    data,
  });
};

/** 更改昵称 */
export const changeNickname = (
  data: NetParams.UserChangeNickname
): Promise<NetReq.UserChangeNickname> => {
  return request({
    method: "post",
    url: "/v1/user/nickname",
    data,
  });
};

/**
 * 更改头像
 * @param {Object} data
 * @returns Promise
 */
export const changeAvatar = (data: any) => {
  return request({
    method: "post",
    url: "/v1/user/avatar",
    data,
  });
};

/** 获取未读消息数 */
export const getUnreadMsgCount = (
  params: NetParams.UserGetUnreadMsgCount = {}
): Promise<NetReq.UserGetUnreadMsgCount> => {
  return request({
    method: "get",
    url: "/v1/user/msgcount/unread",
    params,
  });
};

/** 获取消息列表 */
export const getMessages = (
  params: NetParams.UserGetMessages
): Promise<NetReq.UserGetMessages> => {
  return request({
    method: "get",
    url: "/v1/user/messages",
    params,
  });
};

/**
 * 阅读消息
 * @param {Object} data
 * @returns Promise
 */
export const readMessage = (data: any) => {
  return request({
    method: "post",
    url: "/v1/user/message/read",
    data,
  });
};

/** 获取收藏列表 */
export const getCollections = (
  params: NetParams.UserGetCollections
): Promise<NetReq.UserGetCollections> => {
  return request({
    method: "get",
    url: "/v1/user/collections",
    params,
  });
};

/** 获取用户基础信息 */
export const getUserProfile = (
  params: NetParams.UserGetUserProfile
): Promise<NetReq.UserGetUserProfile> => {
  return request({
    method: "get",
    url: "/v1/user/profile",
    params,
  });
};

/** 获取用户帖子列表 */
export const getUserPosts = (
  params: NetParams.UserGetUserPosts
): Promise<NetReq.UserGetUserPosts> => {
  return request({
    method: "get",
    url: "/v1/user/posts",
    params,
  });
};

/** 获取账单列表 */
export const getBills = (
  params: NetParams.UserGetBills
): Promise<NetReq.UserGetBills> => {
  return request({
    method: "get",
    url: "/v1/user/wallet/bills",
    params,
  });
};

/**
 * 发起充值请求
 * @param {Object} data
 * @returns Promise
 */
export const reqRecharge = (
  data: NetParams.UserReqRecharge
): Promise<NetReq.UserReqRecharge> => {
  return request({
    method: "post",
    url: "/v1/user/recharge",
    data,
  });
};

/**
 * 获取充值状态
 * @param {Object} params
 * @returns Promise
 */
export const getRecharge = (
  params: NetParams.UserGetRecharge
): Promise<NetReq.UserGetRecharge> => {
  return request({
    method: "get",
    url: "/v1/user/recharge",
    params,
  });
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestUsers = (params: {
  k: string;
}): Promise<NetReq.UserGetSuggestUsers> => {
  return request({
    method: "get",
    url: "/v1/suggest/users",
    params,
  });
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestTags = (params: {
  k: string;
}): Promise<NetReq.UserGetSuggestTags> => {
  return request({
    method: "get",
    url: "/v1/suggest/tags",
    params,
  });
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const precheckAttachment = (
  params: NetParams.UserPrecheckAttachment
): Promise<NetReq.UserPrecheckAttachment> => {
  return request({
    method: "get",
    url: "/v1/attachment/precheck",
    params,
  });
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const getAttachment = (
  params: NetParams.UserGetAttachment
): Promise<NetReq.UserGetAttachment> => {
  return request({
    method: "get",
    url: "/v1/attachment",
    params,
  });
};

/**
 * 管理·用户禁言/解禁
 * @param {Object} data
 * @returns Promise
 */
export const changeUserStatus = (
  data: NetParams.UserStatusReq
): Promise<NetReq.UserChangeStatus> => {
  return request({
    method: "post",
    url: "/v1/admin/user/status",
    data,
  });
};
