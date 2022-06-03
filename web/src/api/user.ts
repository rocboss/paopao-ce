import { request } from '@/utils/request';

/** 获取验证码 */
export const getCaptcha = (params: NetParams.UserGetCaptcha = {}): Promise<NetReq.UserGetCaptcha> => {
    return request({
        method: 'get',
        url: '/v1/captcha',
        params
    });
};

/**
 * 发送短信验证码
 * @param {Object} data
 * @returns Promise
 */
export const sendCaptcha = (data: any) => {
    return request({
        method: 'post',
        url: '/v1/captcha',
        data
    });
};

/**
 * 发送用户私信
 * @param {Object} data
 * @returns Promise
 */
export const sendUserWhisper = (data: NetParams.UserWhisper): Promise<NetParams.UserWhisper> => {
    return request({
        method: 'post',
        url: '/v1/user/whisper',
        data
    });
};

/**
 * 绑定用户手机
 * @param {Object} data
 * @returns Promise
 */
export const bindUserPhone = (data: NetParams.UserBindUserPhone): Promise<NetParams.UserBindUserPhone> => {
    return request({
        method: 'post',
        url: '/v1/user/phone',
        data
    });
};

/** 更改密码 */
export const changePassword = (data: NetParams.UserChangePassword): Promise<NetReq.UserChangePassword> => {
    return request({
        method: 'post',
        url: '/v1/user/password',
        data
    });
};

/** 更改昵称 */
export const changeNickname = (data: NetParams.UserChangeNickname): Promise<NetReq.UserChangeNickname> => {
    return request({
        method: 'post',
        url: '/v1/user/nickname',
        data
    });
};

/**
 * 更改头像
 * @param {Object} data
 * @returns Promise
 */
export const changeAvatar = (data: any) => {
    return request({
        method: 'post',
        url: '/v1/user/avatar',
        data
    });
};

/** 获取未读消息数 */
export const getUnreadMsgCount = (params: NetParams.UserGetUnreadMsgCount = {}): Promise<NetReq.UserGetUnreadMsgCount> => {
    return request({
        method: 'get',
        url: '/v1/user/msgcount/unread',
        params
    });
};

/** 获取消息列表 */
export const getMessages = (params: NetParams.UserGetMessages): Promise<NetReq.UserGetMessages> => {
    return request({
        method: 'get',
        url: '/v1/user/messages',
        params
    });
};

/**
 * 阅读消息
 * @param {Object} data
 * @returns Promise
 */
export const readMessage = (data: any) => {
    return request({
        method: 'post',
        url: '/v1/user/message/read',
        data
    });
};

/** 获取收藏列表 */
export const getCollections = (params: NetParams.UserGetCollections): Promise<NetReq.UserGetCollections> => {
    return request({
        method: 'get',
        url: '/v1/user/collections',
        params
    });
};

/** 获取点赞列表 */
export const getStars = (params: NetParams.UserGetStars): Promise<NetReq.UserGetStars> => {
    return request({
        method: 'get',
        url: '/v1/user/stars',
        params
    });
};

/** 获取用户基础信息 */
export const getUserProfile = (params: NetParams.UserGetUserProfile): Promise<NetReq.UserGetUserProfile> => {
    return request({
        method: 'get',
        url: '/v1/user/profile',
        params
    });
};

/** 获取用户帖子列表 */
export const getUserPosts = (params: NetParams.UserGetUserPosts): Promise<NetReq.UserGetUserPosts> => {
    return request({
        method: 'get',
        url: '/v1/user/posts',
        params
    });
};

/** 获取账单列表 */
export const getBills = (params: NetParams.UserGetBills): Promise<NetReq.UserGetBills> => {
    return request({
        method: 'get',
        url: '/v1/user/wallet/bills',
        params
    });
};

/**
 * 发起充值请求
 * @param {Object} data
 * @returns Promise
 */
export const reqRecharge = (data: NetParams.UserReqRecharge): Promise<NetReq.UserReqRecharge> => {
    return request({
        method: 'post',
        url: '/v1/user/recharge',
        data
    });
};

/**
 * 获取充值状态
 * @param {Object} params
 * @returns Promise
 */
export const getRecharge = (params: NetParams.UserGetRecharge): Promise<NetReq.UserGetRecharge> => {
    return request({
        method: 'get',
        url: '/v1/user/recharge',
        params
    });
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestUsers = (params: { k: string }): Promise<NetReq.UserGetSuggestUsers> => {
    return request({
        method: 'get',
        url: '/v1/suggest/users',
        params
    });
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestTags = (params: { k: string }): Promise<NetReq.UserGetSuggestTags> => {
    return request({
        method: 'get',
        url: '/v1/suggest/tags',
        params
    });
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const precheckAttachment = (params: NetParams.UserPrecheckAttachment): Promise<NetReq.UserPrecheckAttachment> => {
    return request({
        method: 'get',
        url: '/v1/attachment/precheck',
        params
    });
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const getAttachment = (params: NetParams.UserGetAttachment): Promise<NetReq.UserGetAttachment> => {
    return request({
        method: 'get',
        url: '/v1/attachment',
        params
    });
};
