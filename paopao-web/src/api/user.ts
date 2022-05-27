import request from '@/utils/request';

/**
 * 获取验证码
 * @param {Object} params
 * @returns Promise
 */
export const getCaptcha = (params: any) => {
    return request({
        method: 'get',
        url: '/captcha',
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
        url: '/captcha',
        data
    });
};

/**
 * 绑定用户手机
 * @param {Object} data
 * @returns Promise
 */
export const bindUserPhone = (data: any) => {
    return request({
        method: 'post',
        url: '/user/phone',
        data
    });
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
export const getUnreadMsgCount = (params: any) => {
    return request({
        method: 'get',
        url: '/user/msgcount/unread',
        params
    });
};

/**
 * 获取消息列表
 * @param {Object} params
 * @returns Promise
 */
export const getMessages = (params: any) => {
    return request({
        method: 'get',
        url: '/user/messages',
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
        url: '/user/message/read',
        data
    });
};

/**
 * 获取收藏列表
 * @param {Object} params
 * @returns Promise
 */
export const getCollections = (params: any) => {
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
export const getStars = (params: any) => {
    return request({
        method: 'get',
        url: '/user/stars',
        params
    });
};

/**
 * 获取用户基础信息
 * @param {Object} params
 * @returns Promise
 */
export const getUserProfile = (params: any) => {
    return request({
        method: 'get',
        url: '/user/profile',
        params
    });
};

/**
 * 获取点赞列表
 * @param {Object} params
 * @returns Promise
 */
export const getUserPosts = (params: any) => {
    return request({
        method: 'get',
        url: '/user/posts',
        params
    });
};

/**
 * 获取账单列表
 * @param {Object} params
 * @returns Promise
 */
export const getBills = (params: any) => {
    return request({
        method: 'get',
        url: '/user/wallet/bills',
        params
    });
};

/**
 * 发起充值请求
 * @param {Object} data
 * @returns Promise
 */
export const reqRecharge = (data: any) => {
    return request({
        method: 'post',
        url: '/user/recharge',
        data
    });
};

/**
 * 获取充值状态
 * @param {Object} params
 * @returns Promise
 */
export const getRecharge = (params: any) => {
    return request({
        method: 'get',
        url: '/user/recharge',
        params
    });
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestUsers = (params: {k: string}) => {
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
export const getSuggestTags = (params: {k: string}) => {
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
export const precheckAttachment = (params: any) => {
    return request({
        method: 'get',
        url: '/attachment/precheck',
        params
    });
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const getAttachment = (params: any) => {
    return request({
        method: 'get',
        url: '/attachment',
        params
    });
};
