import request from '@/utils/request';

/**
 * 获取验证码
 * @param {Object} params
 * @returns Promise
 */
export const getCaptcha = (params: NetParams.UserGetCaptcha = {}) => {
    return request({
        method: 'get',
        url: '/captcha',
        params
    }) as unknown as Promise<NetReq.UserGetCaptcha>;
};

/**
 * 发送短信验证码
 * @param {Object} data
 * @returns Promise
 */
export const sendCaptcha = (data: any) => {
    return request({
        method: 'post',
        url: '/captcha',
        data
    });
};

/**
 * 发送用户私信
 * @param {Object} data
 * @returns Promise
 */
export const sendUserWhisper = (data: NetParams.UserWhisper) => {
    return request({
        method: 'post',
        url: '/user/whisper',
        data
    }) as unknown as Promise<NetParams.UserWhisper>;
};

/**
 * 绑定用户手机
 * @param {Object} data
 * @returns Promise
 */
export const bindUserPhone = (data: NetParams.UserBindUserPhone) => {
    return request({
        method: 'post',
        url: '/user/phone',
        data
    }) as unknown as Promise<NetParams.UserBindUserPhone>;
};

/**
 * 更改密码
 * @param {Object} data
 * @returns Promise
 */
export const changePassword = (data: any) => {
    return request({
        method: 'post',
        url: '/user/password',
        data
    });
};

/**
 * 更改昵称
 * @param {Object} data
 * @returns Promise
 */
export const changeNickname = (data: any) => {
    return request({
        method: 'post',
        url: '/user/nickname',
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
        url: '/user/avatar',
        data
    });
};

/**
 * 获取未读消息数
 * @param {Object} params
 * @returns Promise
 */
export const getUnreadMsgCount = (params: NetParams.UserGetUnreadMsgCount = {}) => {
    return request({
        method: 'get',
        url: '/user/msgcount/unread',
        params
    }) as unknown as Promise<NetReq.UserGetUnreadMsgCount>;
};

/**
 * 获取消息列表
 * @param {Object} params
 * @returns Promise
 */
export const getMessages = (params: NetParams.UserGetMessages) => {
    return request({
        method: 'get',
        url: '/user/messages',
        params
    }) as unknown as Promise<NetReq.UserGetMessages>;
};

/**
 * 阅读消息
 * @param {Object} data
 * @returns Promise
 */
export const readMessage = (data: any) => {
    return request({
        method: 'post',
        url: '/user/message/read',
        data
    });
};

/**
 * 获取收藏列表
 * @param {Object} params
 * @returns Promise
 */
export const getCollections = (params: NetParams.UserGetCollections) => {
    return request({
        method: 'get',
        url: '/user/collections',
        params
    }) as unknown as Promise<NetReq.UserGetCollections>;
};

/**
 * 获取点赞列表
 * @param {Object} params
 * @returns Promise
 */
export const getStars = (params: NetParams.UserGetStars) => {
    return request({
        method: 'get',
        url: '/user/stars',
        params
    }) as unknown as Promise<NetReq.UserGetStars>;
};

/**
 * 获取用户基础信息
 * @param {Object} params
 * @returns Promise
 */
export const getUserProfile = (params: NetParams.UserGetUserProfile) => {
    return request({
        method: 'get',
        url: '/user/profile',
        params
    }) as unknown as Promise<NetReq.UserGetUserProfile>;
};

/**
 * 获取点赞列表
 * @param {Object} params
 * @returns Promise
 */
export const getUserPosts = (params: NetParams.UserGetUserPosts) => {
    return request({
        method: 'get',
        url: '/user/posts',
        params
    }) as unknown as Promise<NetReq.UserGetUserPosts>;
};

/**
 * 获取账单列表
 * @param {Object} params
 * @returns Promise
 */
export const getBills = (params: NetParams.UserGetBills) => {
    return request({
        method: 'get',
        url: '/user/wallet/bills',
        params
    }) as unknown as Promise<NetReq.UserGetBills>;
};

/**
 * 发起充值请求
 * @param {Object} data
 * @returns Promise
 */
export const reqRecharge = (data: NetParams.UserReqRecharge) => {
    return request({
        method: 'post',
        url: '/user/recharge',
        data
    }) as unknown as Promise<NetReq.UserReqRecharge>;
};

/**
 * 获取充值状态
 * @param {Object} params
 * @returns Promise
 */
export const getRecharge = (params: NetParams.UserGetRecharge) => {
    return request({
        method: 'get',
        url: '/user/recharge',
        params
    }) as unknown as Promise<NetReq.UserGetRecharge>;
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestUsers = (params: { k: string }) => {
    return request({
        method: 'get',
        url: '/suggest/users',
        params
    }) as unknown as Promise<NetReq.UserGetSuggestUsers>;
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestTags = (params: { k: string }) => {
    return request({
        method: 'get',
        url: '/suggest/tags',
        params
    }) as unknown as Promise<NetReq.UserGetSuggestTags>;
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const precheckAttachment = (params: NetParams.UserPrecheckAttachment) => {
    return request({
        method: 'get',
        url: '/attachment/precheck',
        params
    }) as unknown as Promise<NetReq.UserPrecheckAttachment>;
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const getAttachment = (params: NetParams.UserGetAttachment) => {
    return request({
        method: 'get',
        url: '/attachment',
        params
    }) as unknown as Promise<NetReq.UserGetAttachment>;
};
